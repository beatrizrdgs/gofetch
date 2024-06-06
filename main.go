package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	printHostname()
	printUsername()
	printGoVersion()
}

func printHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}

	fmt.Println("hostname:", hostname)
}

func printUsername() {
	username := os.Getenv("USER")
	fmt.Println("username:", username)
}

func printGoVersion() {
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		return
	}
	fmt.Println(string(out[13:19]))
}
