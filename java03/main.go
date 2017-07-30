package main

import (
	"fmt"
	"jvmgo/java03/classfile"
	"jvmgo/java03/classpath"
	"log"
	"strings"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	log.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	data, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}

	cf, _ := classfile.Parse(data)
	fmt.Printf("classFile:%#v\n", cf)
}
