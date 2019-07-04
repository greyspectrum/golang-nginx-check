package main

import (
	"fmt"
	"os/exec"
)

func main() {
	for {
		cmdStatus := exec.Command("systemctl", "check", "nginx")
		cmdRestart := exec.Command("systemctl", "restart", "nginx")
		out, err := cmdStatus.CombinedOutput()
		if err != nil {
			fmt.Println("Cannot find process")
			cmdRestart.Run()
		} else {
			fmt.Printf("Nginx is: %s", string(out))
		}
	}
}
