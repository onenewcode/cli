/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var name string

// rootCmd 表示调用时不带任何子命令的基本命令
var rootCmd = &cobra.Command{
	// 命令行的名称
	Use: "cli_demo",
	// 段描述
	Short: "A brief description of your application",
	// 长描述，可以换行
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		if name != "" {
			fmt.Printf("Hello, %s!\n", name)
		} else {
			fmt.Println("Hello, World!")
		}
	},
}
var versionCmd = &cobra.Command{
	// 命令行的名称
	Use: "version",
	// 段描述
	Short: "A brief description of your application",
	// 长描述，可以换行
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		if name != "" {
			fmt.Printf("Hello, %s!\n", name)
		} else {
			fmt.Println("Hello, World!")
		}
	},
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
	//
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name to greet")
	// 添加子命令行
	rootCmd.AddCommand(versionCmd)
}
