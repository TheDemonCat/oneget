package downloader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"sync"
)

func NewClient(loginUrl string, baseUrl string, login string, password string) (*Client, error) {

	cj, _ := cookiejar.New(nil)

	c := &Client{
		login:    login,
		password: password,
		loginUrl: loginUrl,
		baseUrl:  baseUrl,
		cookie:   cj,
	}

	url, err := c.getAuthTicketURL(baseUrl)
	if err != nil {
		return nil, err
	}

	loginResp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer loginResp.Body.Close()

	return c, nil
}

type Client struct {
	cookie   *cookiejar.Jar
	login    string
	password string
	loginUrl string
	baseUrl  string
}

func (c *Client) getAuthTicketURL(url string) (string, error) {

	type loginParams struct {
		Login       string `json:"login"`
		Password    string `json:"password"`
		ServiceNick string `json:"serviceNick"`
	}

	type ticket struct {
		Ticket string `json:"ticket"`
	}

	ticketUrl := c.loginUrl + "/rest/public/ticket/get"
	postBody, err := json.Marshal(
		loginParams{c.login, c.password, url})
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(postBody)
	defer put(buf)
	req, err := http.NewRequest("POST", ticketUrl, buf)

	if err != nil {
		return "", err
	}

	req.SetBasicAuth(c.login, c.password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.doRequest(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:

		var ticketData ticket
		err := bodyToJSON(resp.Body, &ticketData)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(loginURL+"/ticket/auth?token=%s", ticketData.Ticket), nil

	default:

		type ErrorRespond struct {
			Timestamp string `json:"timestamp"`
			Status    int    `json:"status"`
			Error     string `json:"error"`
			Exception string `json:"exception"`
			Message   string `json:"message"`
			Path      string `json:"path"`
		}

		var errData ErrorRespond

		err := bodyToJSON(resp.Body, &errData)
		if err != nil {
			return "", err
		}

		return "", fmt.Errorf("%s: %s", errData.Error, errData.Message)
	}
}

func (c *Client) Get(getUrl string) (*http.Response, error) {

	// для URL вида /total/
	if strings.HasPrefix(getUrl, "/") {
		getUrl = c.baseUrl + getUrl
	}

	req, err := http.NewRequest("GET", getUrl, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.login, c.password)

	resp, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		log.Debugf("Re-authorized with ticket url: %s", getUrl)
		url, err := c.getAuthTicketURL(getUrl)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return nil, err
		}

		req.SetBasicAuth(c.login, c.password)

		return c.doRequest(req)

	case http.StatusBadRequest, http.StatusNotFound:

		return nil, fmt.Errorf("respose CODE:%d  ERR:%s",
			resp.StatusCode, readBodyMustString(resp.Body))
	case http.StatusOK:
		return resp, nil
	default:
		return resp, fmt.Errorf("unknown respose CODE: <%d>", resp.StatusCode)
	}

}

func (c *Client) client() *http.Client {

	return &http.Client{
		Jar: c.cookie,
	}
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {

	return c.client().Do(req)

}

func bodyToJSON(body io.ReadCloser, into interface{}) error {
	b, err := readBody(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, into)
}

func readBody(body io.ReadCloser) ([]byte, error) {
	buf := get()
	defer put(buf)
	_, err := io.Copy(buf, body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), err
}

func readBodyMustString(body io.ReadCloser) string {
	buf, err := readBody(body)
	if err != nil {
		log.Errorf("must read body err: %s", err.Error())
		return ""
	}

	return string(buf)
}

var pool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func get() *bytes.Buffer {
	return pool.Get().(*bytes.Buffer)
}

func put(buf *bytes.Buffer) {
	buf.Reset()
	pool.Put(buf)
}
