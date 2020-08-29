package main

import (
	"fmt"
	 "github.com/urfave/cli"
)
const usage = `mydocker is a simple container runtime implementation.
			   The purpose of this project is to learn how docker works and how to write a docker by ourselves
			   Enjoy it, just for fun.`
func main()  {
	fmt.Println("hello world")

	app := cli.NewApp()
	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}
	app.Name = "my docker"
	app.Usage = usage
	println(app.Usage)
	println(app.Commands)
}