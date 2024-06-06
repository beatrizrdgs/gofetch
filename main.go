package main

import (
	"fmt"
	"os"
)

func main() {
	// uname := unix.Utsname{}
	// err := unix.Uname(&uname)
	// if err != nil {
	// 	fmt.Errorf("err %w", err)
	// 	return
	// }

	// fmt.Println("sysname:", string(uname.Sysname[:]))
	// fmt.Println("nodename:", string(uname.Nodename[:]))
	// fmt.Println("release:", string(uname.Release[:]))
	// fmt.Println("version:", string(uname.Version[:]))
	// fmt.Println("machine:", string(uname.Machine[:]))

	printHostname()
	printUsername()
}

func printHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Errorf("err %w", err)
		return
	}

	fmt.Println("hostname:", hostname)
}

func printUsername() {
	username := os.Getenv("USER")
	fmt.Println("username:", username)
}
