package gofetch

import (
	"fmt"
	"os"
)

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

const (
	UNKNOWN string = "unknown"
)

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
		Shell:     getShell(),
		CPU:       getCPU(),
		GPU:       getGPU(),
		RAM:       getRAMUsage(),
		Disk:      getDiskUsage(),
	}
}

func (s *System) getUserAtHost() string {
	return s.Username + "@" + s.Hostname
}

func (s *System) getDashes(str string) string {
	var dashes string
	for i := 0; i < len(str); i++ {
		dashes += "-"
	}
	return dashes
}

func (s *System) makeAttrs() []string {
	userAtHost := s.getUserAtHost()
	return []string{
		userAtHost,
		s.getDashes(userAtHost),
		fmt.Sprintf("Host: %s", s.Host),
		fmt.Sprintf("Go version: %s", s.GoVersion),
		fmt.Sprintf("Distro: %s", s.Distro),
		fmt.Sprintf("Kernel: %s", s.Kernel),
		fmt.Sprintf("Shell: %s", s.Shell),
		fmt.Sprintf("CPU: %s", s.CPU),
		fmt.Sprintf("GPU: %s", s.GPU),
		fmt.Sprintf("Memory: %s MiB / %s MiB", s.Disk, s.RAM),
		getColorBar(colors),
		getColorBar(brightColors),
	}
}

func (s *System) Print() {
	fmt.Println(s.replaceASCII(gopherASCII, s.makeAttrs()))
}
