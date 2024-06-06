package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	sys := NewSystem()
	sys.printAttributes()
}

type System struct {
	Hostname  string
	Username  string
	GoVersion string
	Distro    string
	Kernel    string
	Shell     string
	CPU       string
	GPU       string
	RAM       string
	Disk      string
}

func NewSystem() *System {
	return &System{
		Hostname:  getHostname(),
		Username:  getUsername(),
		GoVersion: getGoVersion(),
		Distro:    getDistro(),
		Kernel:    getKernel(),
		Shell:     os.Getenv("SHELL"),
		CPU:       getCPU(),
		GPU:       getGPU(),
		RAM:       getRAMUsage(),
		Disk:      getDiskUsage(),
	}
}

func (s *System) printAttributes() {
	fmt.Println(s.Username + "@" + s.Hostname)
	fmt.Println("------------------------")
	fmt.Println("Hostname:", s.Hostname)
	fmt.Println("Username:", s.Username)
	fmt.Println("Go version:", s.GoVersion)
	fmt.Println("Distro:", s.Distro)
	fmt.Println("Kernel:", s.Kernel)
	fmt.Println("Shell:", s.Shell)
	fmt.Println("CPU:", s.CPU)
	fmt.Println("GPU:", s.GPU)
	fmt.Println("Memory:", s.Disk, "MiB /", s.RAM, "MiB")
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}

	return hostname
}

func getUsername() string {
	username := os.Getenv("USER")
	if username == "" {
		return ""
	}

	return username
}

func getGoVersion() string {
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(out[13:19])
}

func getDistro() string {
	cmd := exec.Command("cat", "/etc/os-release")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	start := strings.Index(string(out), "NAME=") + len("NAME=") + 1
	end := strings.Index(string(out), "\"\n")
	return string(out)[start:end]
}

func getKernel() string {
	cmd := exec.Command("uname", "-r")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(out[:len(out)-1])
}

func getCPU() string {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return ""
	}

	var cpus []string
	cpuSet := make(map[string]bool)
	for _, cpu := range cpuInfo {
		if _, exists := cpuSet[cpu.ModelName]; !exists {
			cpus = append(cpus, cpu.ModelName)
			cpuSet[cpu.ModelName] = true
		}
	}

	return strings.Join(cpus, ", ")
}

func getGPU() string {
	cmd := exec.Command("lspci")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "VGA compatible controller") || strings.Contains(line, "3D controller") {
			start := strings.Index(line, "controller") + len("controller: ")
			end := strings.Index(line, " (rev")
			return line[start:end]
		}
	}

	return ""
}

func getRAMUsage() string {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return ""
	}
	ram := vmStat.Total / 1024 / 1024
	return fmt.Sprint(ram)
}

func getDiskUsage() string {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return ""
	}
	disk := diskStat.Total / 1024 / 1024
	return fmt.Sprint(disk)
}

var gopherASCII = `
                    %%%%%%%%%%%
    %%%%    %%%%%%%%%%%%%%%%%%%%%%%%%%%   %%%%
   %%%( %%%%@@@@@@@@@@%%%%%%%@@@@@@@@@%%%%   %%
     %%%%%@@@@@@@@@@@@@@%%%@@@@@@@@@@@@@%%%%%%
       %%@@@@@@@@      @%%@@@@@@@@      @%%%
       %%@@@@@@@@.    @@%%%@@@@@@@     @@%%%
       %%%&@@@@@@@@@@@@.....@@@@@@@@@@@%%%%%
       %%%%%%%@@@@@@(((.....(((@@@@@%%%%%%%%
       %%%%%%%%%%%%(((((((((((((%%%%%%%%%%%%
       %%%%%%%%%%%%%%%%@@@@@%%%%%%%%%%%%%%%%
       %%%%%%%%%%%%%%%%(((((%%%%%%%%%%%%%%%%
       %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
       %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
       %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
       %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
       %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
       %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
       %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
`
