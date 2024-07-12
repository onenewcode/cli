package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		// 命令行应用的名称
		Name: "greet",
		// 程序的简短描述
		Usage: "fight the loneliness!",
		// 当应用程序被调用时会执行这个函数
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
