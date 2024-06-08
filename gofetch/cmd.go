package gofetch

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func getCmdOutput(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	out, err := cmd.Output()
	if err != nil {
		return UNKNOWN
	}
	return string(out)
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
