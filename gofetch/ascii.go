package gofetch

import "strings"

var gopherASCII = `
          ,_---~~~~~----._         
   _,,_,*^____       _____^*,_,,_         %s
  / __/ /'     \    /     '\ \__ \        %s 
 [  @f | @))    |  | @))   | f@  ]        %s
  \/    \~____ / __ \_____/    \/         %s
  |            _l__l_           I         %s
  }           [______]          I         %s
  ]             |_|_|           |         %s
  ]                             |         %s
  |                             |         %s
  |                             |         %s
					  %s
					  %s
   `

func (s *System) replaceASCII(ascii string, ss []string) string {
	for _, s := range ss {
		ascii = strings.Replace(ascii, "%s", s, 1)
	}
	return ascii
}
