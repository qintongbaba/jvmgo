package main

import (
	"flag"
	"fmt"
	"os"
)

//命令行结构体
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

//解析命令行
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "输出此帮助信息")
	flag.BoolVar(&cmd.helpFlag, "?", false, "输出此帮助信息")
	flag.BoolVar(&cmd.versionFlag, "version", false, "输出产品版本并退出")
	flag.BoolVar(&cmd.versionFlag, "showversion", false, "输出产品版本并退出")
	flag.StringVar(&cmd.cpOption, "cp", "", "<目录和 zip/jar 文件的类搜索路径>")
	flag.StringVar(&cmd.cpOption, "classpath", "", "<目录和 zip/jar 文件的类搜索路径>")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

//提示帮助信息
func printUsage() {
	fmt.Printf("用法: %s [-option] class [args...] \n", os.Args[0])
}
