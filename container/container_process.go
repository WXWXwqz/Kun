package container

import (
	"fmt"
	"syscall"
	"os/exec"
	"os"
)

func NewParentProcess(tty bool, command string) *exec.Cmd {
	//f,_ := os.Create("jntm.txt")
	args := []string{"init", command}
	fmt.Println(command)
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		//cmd.Stdout=f
		cmd.Stderr = os.Stderr
	}
	return cmd
}
