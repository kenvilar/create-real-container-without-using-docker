package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
		case "child":
			child()
		default:
			parent()
	}
}

func parent() {
	cmd := exec.Command(
		"/proc/self/exe",
		append(
			[]string{"child"},
			os.Args[1:]...,
		)...,
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS
	}
	
	must(cmd.Run())
}

func child() {
	fmt.Printf("running: %v\n", os.Args[2:])
	
	exit, err := chroot("/home/kenvilar/Documents/docker_projects/create-real-container-without-using-docker/rootfs")
	defer exit() //run before exit
	must(err) //warn if there's an error
	must(os.Chdir("/"))
	
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	must(cmd.Run())
}
