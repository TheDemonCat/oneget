package main

import (
	"fmt"
	"github.com/khorevaa/logos"
	"github.com/v8platform/oneget/cmd"

	"os"
	"strings"

	"github.com/urfave/cli/v2"
	dloader "github.com/v8platform/oneget/downloader"
)

var (
	version = "v0.0.7"
	commit  = ""
	date    = ""
	builtBy = ""
)

var log = logos.New("github.com/v8platform/oneget").Sugar()

func setFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u"},
			EnvVars:  []string{"ONEGET_USER", "ONEC_USERNAME"},
			Usage:    "Пользователь портала 1С",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "pwd",
			Aliases:  []string{"p"},
			EnvVars:  []string{"ONEGET_PASSWORD", "ONEC_PASSWORD"},
			Usage:    "Пароль пользователя портала 1С",
			Required: true,
		},
		&cli.StringFlag{
			Name: "nicks",
			Usage: `Имена приложений, разделенные запятой (например \"platform83, EnterpriseERP20\"), 
					подсмотреть можно в адресе, ссылки имею вид например https://releases.1c.ru/project/EnterpriseERP20`,
		},
		&cli.StringFlag{
			Name:  "version-filter",
			Usage: "Фильтр версий по номеру (регулярное выражение)",
		},
		&cli.StringFlag{
			Name:  "distrib-filter",
			Usage: "Дополнительный фильтр пакетов (регулярное выражение)",
		},
		&cli.StringFlag{
			Name:        "path",
			DefaultText: "downloads",
			Usage:       "Путь к каталогу выгрузки",
		},
		&cli.BoolFlag{
			Name:    "debug",
			EnvVars: []string{"ONEGET_DEBUG"},
			Usage:   "Режим отладки приложения",
		},
		&cli.StringFlag{
			Name:        "logs",
			DefaultText: "oneget.logs",
			Value:       "oneget.logs",
			Usage:       "Файл лога загрузки",
		},
		&cli.BoolFlag{
			Name:  "extract",
			Value: false,
			Usage: "Распаковывать дистрибутив",
		},
		&cli.StringFlag{
			Name:        "extractPath",
			DefaultText: "pack",
			Value:       "pack",
			Usage:       "Каталог распаковки дистрибутива",
		},
		&cli.BoolFlag{
			Name:    "rename",
			Aliases: []string{"sl"},
			Value:   false,
			Usage: `Переименовывать дистрибутивы при распаковке. 
				Примеры: 
					1c-enterprise-8.3.18.1334-client_8.3.18-1334_amd64.deb -> client-8.3.18.1334.deb
					1c-enterprise83-server_8.3.16-1876_amd64.deb -> server_8.3.16-1876.deb`,
		},
	}
}

func main() {
	app := &cli.App{
		Name:    "oneget",
		Usage:   "Приложение для загрузки релизов сайта релизов 1С",
		Version: buildVersion(),
		Flags:   setFlags(),
		Action: func(c *cli.Context) error {
			downloaderConfig := dloader.Config{
				Login:         c.String("user"),
				Password:      c.String("pwd"),
				BasePath:      c.String("path"),
				StartDate:     StartDate(c.String("startDate")),
				Nicks:         Nicks(strings.ToLower(c.String("nicks"))),
				VersionFilter: c.String("version-filter"),
				DistribFilter: c.String("distrib-filter"),
				Extract:       c.Bool("extract"),
				ExtractPath:   c.String("extractPath"),
				Rename:        c.Bool("rename"),
			}

			debug := c.Bool("debug")

			if debug {
				logos.SetLevel("github.com/v8platform/oneget", logos.DebugLevel)
			}

			downloader := dloader.New(downloaderConfig)

			files, err := downloader.Get()

			if err == nil {
				log.Infof("Downloaded <%d> files", len(files))
			}

			return err
		},
	}

	for _, command := range cmd.Commands {
		app.Commands = append(app.Commands, command.Cmd())
	}

	err := app.Run(os.Args)
	defer log.Sync()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func buildVersion() string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
