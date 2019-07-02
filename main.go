package main

import (
    "os/exec"
    "fmt"
    "os"
)

func main() {
    cmdStatus := exec.Command("systemctl", "check", "nginx")
    cmdRestart := exec.Command("systemctl", "restart", "nginx")
    out, err := cmdStatus.CombinedOutput()
    if err != nil {
        fmt.Println("Cannot find process")
	cmdRestart.Run()
        os.Exit(1)
    }
    fmt.Printf("Nginx is: %s", string(out))
}
