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
