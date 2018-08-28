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
	
	must(cmd.Run())
}

func child() {
	fmt.Printf("running: %v\n", os.Args[2:])
	
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
