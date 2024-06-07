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

const (
	UNKNOWN = "unknown"
)

func main() {
	sys := NewSystem()
	sys.printAttributes()
}

type System struct {
	Hostname  string
	Username  string
	Host      string
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
	hostname, err := os.Hostname()
	if err != nil {
		hostname = UNKNOWN
	}

	return &System{
		Hostname:  hostname,
		Username:  os.Getenv("USER"),
		Host:      getHost(),
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
	fmt.Println("Host:", s.Host)
	fmt.Println("Go version:", s.GoVersion)
	fmt.Println("Distro:", s.Distro)
	fmt.Println("Kernel:", s.Kernel)
	fmt.Println("Shell:", s.Shell)
	printMultiValues("CPU", s.CPU)
	printMultiValues("GPU", s.GPU)
	fmt.Println("Memory:", s.Disk, "MiB /", s.RAM, "MiB")
}

func printMultiValues(item, val string) {
	if strings.Contains(val, ", ") {
		ss := strings.Split(val, ", ")
		for _, s := range ss {
			fmt.Println(item+":", s)
		}
	} else {
		fmt.Println(item+":", val)
	}
}

func getHost() string {
	out := getCmdOutput("cat", "/sys/devices/virtual/dmi/id/product_name")
	return out[:len(out)-1]
}

func getGoVersion() string {
	out := getCmdOutput("go", "version")
	return out[13:19]
}

func getDistro() string {
	out := getCmdOutput("cat", "/etc/os-release")
	start := strings.Index(out, "NAME=") + len("NAME=") + 1
	end := strings.Index(out, "\"\n")
	return out[start:end]
}

func getKernel() string {
	out := getCmdOutput("uname", "-r")
	return out[:len(out)-1]
}

func getCPU() string {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return UNKNOWN
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
	out := getCmdOutput("lspci")

	var gpus []string
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		if strings.Contains(line, "VGA compatible controller") || strings.Contains(line, "3D controller") {
			start := strings.Index(line, "controller") + len("controller: ")
			end := strings.Index(line, " (rev")
			gpus = append(gpus, line[start:end])
		}
	}

	return strings.Join(gpus, ", ")
}

func getRAMUsage() string {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return UNKNOWN
	}
	ram := vmStat.Total / 1024 / 1024
	return fmt.Sprint(ram)
}

func getDiskUsage() string {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return UNKNOWN
	}
	disk := diskStat.Total / 1024 / 1024
	return fmt.Sprint(disk)
}

func getCmdOutput(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	out, err := cmd.Output()
	if err != nil {
		return UNKNOWN
	}
	return string(out)
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
