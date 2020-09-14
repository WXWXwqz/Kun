package main

import (
	"Kun/container"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"Kun/cgroups/subsystems"

	//"github.com/WXWXwqz/Kun/container"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			mydocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name: "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name: "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name: "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArray []string
		fmt.Println("*************************************")
		for _, arg := range context.Args() {
			fmt.Println(arg)
			cmdArray = append(cmdArray, arg)
		}
		fmt.Println("*************************************")
		tty := context.Bool("ti")
		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet: context.String("cpuset"),
			CpuShare:context.String("cpushare"),
		}
		fmt.Println("--------------------------------------")
		fmt.Println(context.String("m"))
		fmt.Println(context.String("cpuset"))
		fmt.Println(context.String("cpushare"))
		fmt.Println("--------------------------------------")
		Run(tty, cmdArray, resConf)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}

type buff_in int

var helloCommand = cli.Command{
	Name:  "hello",
	Usage: "sey hello for you",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "rap",
			Usage: "enable cai xun kun rap",
		},
	},
	Action: func(context *cli.Context) error {
		f,_ := os.Create("jntm.txt")
		//args := []string{"hello -rap qqq www"}
		log.Infof("I am CaixuKun")
		log.Infof("I can Chang Tiao Rap Lanqiu")

		for i,arg := range context.Args(){
			fmt.Println(i)
			fmt.Println(arg)

			}
		fmt.Println("I exe myself")
		rap := context.Bool("rap")
		if rap {
			for i:=0;i<100;i++{
				fmt.Print("jinitaimei")
			}
			cmd := exec.Command("/proc/self/exe","hello")

			cmd.Stdin = os.Stdin
			//cmd.Stdout = os.Stdout
			cmd.Stdout = f
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				fmt.Println("error")
				log.Error(err)
			}
			cmd.Wait()
		}

		cmd0 := context.Args().Get(0)
		log.Infof("command %s", cmd0)
		cmd1 := context.Args().Get(1)
		log.Infof("command %s", cmd1)
		fmt.Print("CHang 跳 rangop")
		for i:=0;i<100;i++{
			fmt.Print("chang tiao rap ")
		}
		return nil
	},
}
var tstselfCommand = cli.Command{
	Name:  "self",
	Usage: "tst self",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "rap",
			Usage: "enable cai xun kun rap",
		},
	},
	Action: func(context *cli.Context) error {
		f,_ := os.Create("self.txt")
		cmd := exec.Command("/proc/self/exe","hello")
		rap := context.Bool("rap")
		if rap{
			cmd.Stdout=f
		}
		if err := cmd.Start(); err != nil {
			fmt.Println("error")
			log.Error(err)
		}
		fmt.Println("end")
		return nil
	},
}