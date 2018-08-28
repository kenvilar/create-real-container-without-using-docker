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
}
