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

func (s *System) coloredUserAtHost() string {
	return fmt.Sprintf(CYAN + s.Username + RESET + "@" + CYAN + s.Hostname + RESET)
}

func (s *System) getDashes(str string) string {
	var dashes string
	for i := 0; i < len(str); i++ {
		dashes += "-"
	}
	return dashes
}

func (s *System) makeAttrs() []string {
	return []string{
		s.coloredUserAtHost(),
		s.getDashes(s.getUserAtHost()),
		fmt.Sprintf("%sHost%s: %s", CYAN, RESET, s.Host),
		fmt.Sprintf("%sGo version%s: %s", CYAN, RESET, s.GoVersion),
		fmt.Sprintf("%sDistro%s: %s", CYAN, RESET, s.Distro),
		fmt.Sprintf("%sKernel%s: %s", CYAN, RESET, s.Kernel),
		fmt.Sprintf("%sShell%s: %s", CYAN, RESET, s.Shell),
		fmt.Sprintf("%sCPU%s: %s", CYAN, RESET, s.CPU),
		fmt.Sprintf("%sGPU%s: %s", CYAN, RESET, s.GPU),
		fmt.Sprintf("%sMemory%s: %s MiB / %s MiB", CYAN, RESET, s.Disk, s.RAM),
		getColorBar(colors),
		getColorBar(brightColors),
	}
}

func (s *System) Print() {
	fmt.Println(s.replaceASCII(gopherASCII, s.makeAttrs()))
}
