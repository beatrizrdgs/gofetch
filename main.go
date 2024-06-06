package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/shirou/gopsutil/cpu"
)

type System struct {
	Hostname  string
	Username  string
	GoVersion string
	Distro    string
	Kernel    string
	Shell     string
}

type Hardware struct {
	CPU    string
	GPU    string
	Memory string
}

func main() {
	// system := System{}
	hardware := Hardware{}

	// system.printHostname()
	// system.printUsername()
	// system.printGoVersion()

	hardware.printCPU()
	// hardware.printGPU()
	// hardware.printMemory()
}

func (s *System) printHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}

	fmt.Println("hostname:", hostname)
}

func (s *System) printUsername() {
	username := os.Getenv("USER")
	fmt.Println("username:", username)
}

func (s *System) printGoVersion() {
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		return
	}
	fmt.Println(string(out[13:19]))
}

func (h *Hardware) printCPU() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return
	}

	var cpus []string
	cpuSet := make(map[string]bool)
	for _, cpu := range cpuInfo {
		if _, exists := cpuSet[cpu.ModelName]; !exists {
			cpus = append(cpus, cpu.ModelName)
			cpuSet[cpu.ModelName] = true
		}
	}

	for _, cpu := range cpus {
		fmt.Println("CPU:", cpu)
	}
}
