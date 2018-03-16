package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	err := exec.Command("mv", arg1, arg2).Run()
	if err != nil {
		err = exec.Command("mkdir", "-p", strings.Split(arg2, "/")[0]).Run()
		if err != nil {
			panic(err)
		}
	}
	err = exec.Command("mv", arg1, arg2).Run()
	if err != nil {
		panic(err)
	}
}
