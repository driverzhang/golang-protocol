package main

import (
	"fmt"
	"github.com/driverzhang/golang-protocol/tool/protoc"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "protocol"
	app.Usage = "protocol-golang协议转换工具集"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "Protocol",
			Aliases: []string{"pb"},
			Usage:   "golang struct for protoc",
			Action:  protoc.Go2protoc,
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "protocol version",
			Action: func(c *cli.Context) error {
				fmt.Println(getVersion())
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
