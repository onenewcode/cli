# cli
## 例子
```go
package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

```

将我们的命令安装到目录中：$GOPATH/bin


>$ go install

最后运行我们的新命令：


>$ greet
Hello friend!

CLI 还会生成简洁的帮助文本：

```shell
$ greet help
NAME:
    greet - fight the loneliness!

USAGE:
    greet [global options] command [command options] [arguments...]

COMMANDS:
    help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS
    --help, -h  show help (default: false)
```
## 命令行参数
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "Nefertiti"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}
			//
			if cCtx.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

```
```shell
go run .\flags\main.go  -lang  spanish
Hola Nefertiti
```
## 子命令行
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "Nefertiti"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}
			// 获取指定的flag
			if cCtx.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

```

```shell
go run .\subcommands\main.go a dsfdsfd 
added task:  dsfdsfd

```

## 官网
https://cli.urfave.org/v2/examples/bash-completions/

#  cobra
## 概念
Cobra 建立在命令、参数和标志这三个结构之上。要使用 Cobra 编写一个命令行程序，需要明确这三个概念。



命令（COMMAND）：命令表示要执行的操作。

参数（ARG）：是命令的参数，一般用来表示操作的对象。

标志（FLAG）：是命令的修饰，可以调整操作的行为。



一个好的命令行程序在使用时读起来像句子，用户会自然的理解并知道如何使用该程序。



要编写一个好的命令行程序，需要遵循的模式是 APPNAME VERB NOUN --ADJECTIVE 或 APPNAME COMMAND ARG --FLAG。
## 快速开始
使用 Cobra 很容易。首先，用于安装最新版本 的图书馆。此命令将安装生成器可执行文件 以及库及其依赖项：go getcobra
>go get -u github.com/spf13/cobra/cobra

接下来，将 Cobra 包含在您的应用程序中：
>import "github.com/spf13/cobra"

此外我们需要下载一个 Cobra 的命令行工具，它将生成一个新项目，并包含一个基本的命令行工具。
>go install github.com/spf13/cobra-cli@latest

然后我们在项目目录下输入以下命令创建我们的命令行工具
>cobra-cli init

之后便会生成以下目录
```shell
D:.
│   go.mod
│   go.sum
│   LICENSE
│   main.go
│   README.md
└───cmd
        root.go
```
root.go是我们的命令行程序，我们可以在main.go中运行。效果如下
```shell
go run .\main.go
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

```