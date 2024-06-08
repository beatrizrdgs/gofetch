package gofetch

import "fmt"

const (
	BLACK   = "\033[0;30m"
	RED     = "\033[0;31m"
	GREEN   = "\033[0;32m"
	YELLOW  = "\033[0;33m"
	BLUE    = "\033[0;34m"
	MAGENTA = "\033[0;35m"
	CYAN    = "\033[0;36m"
	WHITE   = "\033[0;37m"

	GRAY           = "\033[1;30m"
	BRIGHT_RED     = "\033[1;31m"
	BRIGHT_GREEN   = "\033[1;32m"
	BRIGHT_YELLOW  = "\033[1;33m"
	BRIGHT_BLUE    = "\033[1;34m"
	BRIGHT_MAGENTA = "\033[1;35m"
	BRIGHT_CYAN    = "\033[1;36m"
	BRIGHT_WHITE   = "\033[1;37m"

	RESET = "\033[0m"
)

var (
	colors = []string{
		BLACK,
		RED,
		GREEN,
		YELLOW,
		BLUE,
		MAGENTA,
		CYAN,
		WHITE,
	}

	brightColors = []string{
		GRAY,
		BRIGHT_RED,
		BRIGHT_GREEN,
		BRIGHT_YELLOW,
		BRIGHT_BLUE,
		BRIGHT_MAGENTA,
		BRIGHT_CYAN,
		BRIGHT_WHITE,
	}
)

func getColorBar(cc []string) string {
	var bar string
	for _, c := range cc {
		bar += fmt.Sprintf("%s█████", c)
	}
	bar += RESET
	return bar
}
