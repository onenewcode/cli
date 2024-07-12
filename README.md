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
### 注释后的代码
```go
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd 表示调用时不带任何子命令的基本命令
var rootCmd = &cobra.Command{
	// 命令行的名称
	Use: "cli_demo",
	// 段描述
	Short: "A brief description of your application",
	// 长描述，可以换行
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// 函数负责执行根命令 rootCmd。当调用 rootCmd.Execute() 时，Cobra 会解析命令行参数，
// 并根据用户的输入执行相应的子命令或动作。如果 Execute 方法返回错误，os.Exit(1) 会立即终止程序，
// 并返回状态码 1，通常表示程序异常退出。
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init 函数会在程序开始执行之前被调用，可以用来初始化全局变量或执行一些预先设定的操作。
// 在 Cobra 中，init 函数经常用于定义和初始化命令行标志。
func init() {
	// 义了一个布尔类型的标志 toggle，它可以通过 -toggle 或 -t （短标记）在命令行中使用。false 是该标志的默认值，如果在命令行中没有指定，则使用这个默认值。
	// 最后的字符串 "Help message for toggle" 是当用户请求帮助信息时，与这个标志相关的描述。
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
```