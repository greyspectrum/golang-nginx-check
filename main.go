package main

import (
    "os/exec"
    "fmt"
    "os"
)

func main() {
    cmd := exec.Command("systemctl", "check", "nginx")
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Cannot find process")
        cmd := exec.Command("systemctl", "restart", "nginx")
	os.Exit(1)
    }
    fmt.Printf("Nginx is: %s", string(out))
}
