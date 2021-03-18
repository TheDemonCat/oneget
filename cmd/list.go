package cmd

import (
	"github.com/k0kubun/pp"
	"github.com/khorevaa/logos"
	"github.com/urfave/cli/v2"
	dloader "github.com/v8platform/oneget/downloader"
)

type listCmd struct {
	showUnavailablePrograms bool
}

func (c *listCmd) run(ctx *cli.Context) error {

	user := ctx.String("user")
	password := ctx.String("pwd")

	if ctx.Bool("debug") {
		logos.SetLevel("github.com/v8platform/oneget", logos.DebugLevel)
	}

	dl := dloader.NewDownloader(
		user,
		password,
	)

	projects, err := dl.GetListProject(c.showUnavailablePrograms)
	if err != nil {
		return err
	}

	pp.Println(projects)

	pp.Println("Всего проектов:", len(projects))

	return nil
}

func (c *listCmd) Cmd() *cli.Command {

	cmd := &cli.Command{
		Name:      "list",
		Usage:     "Выводит список проектов и их релив сайта 1С",
		ArgsUsage: "PROJECTS...",
		Flags:     c.Flags(),
		Action:    c.run,
		Before: func(ctx *cli.Context) error {

			//if !ctx.Args().Present() {
			//	err := cli.ShowSubcommandHelp(ctx)
			//	if err != nil {
			//		return err
			//	}
			//	return fmt.Errorf("WRONG USAGE: Requires a RELEASE argument")
			//}
			//
			//c.releases = ctx.Args().Slice()

			return nil
		},
	}

	return cmd
}

func (c *listCmd) Flags() []cli.Flag {
	return []cli.Flag{

		&cli.BoolFlag{
			Name:        "show-unavailable",
			Destination: &c.showUnavailablePrograms,
			Aliases:     []string{"U"},
			Value:       false,
			Usage:       "Вывести не доступные проекты 1С",
		},
	}
}
