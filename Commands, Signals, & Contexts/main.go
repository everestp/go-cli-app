package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-la")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(os.Stderr, "failed to run command", err)
	}

	fmt.Println("Commmnad :", string(stdout))
}
