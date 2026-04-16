package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

const (
	logoColor = "\033[38;2;13;153;255m"
	logoReset = "\033[0m"
)

const logo = `                                                
                       ██                       
                    ████████                    
                 ██████████████                 
              █████████████████████             
          ███████████████████████████           
       ██████████████████████████████   █       
    █████████████████      ██████████   █████   
 ████████████████        ████████████   ████████
██████████████        ███████████████   ████████
███████████       ██████████████████    ████████
████████       █████████████████        ████████
█████       ██████████████████          ████████
█              ███████████      █████   ████████
     ██████       █████      ████████   ████████
 ██████████████             ████████    ████████
██████████████████          ████████    ████████
█████████████████████       ████████    ████████
████████████████████████    ████████    ████████
████████   █████████████    ████████    ████████
████████      ██████████    ████████    ████████
██████████       ███████    ████████   █████████
█████████████       ████    ████████████████████
████████████████       █    ████████████████████
  █████████████████         ██████████████████  
     ██████████████████      ██████████████     
         █████████████████      ████████        
            █████████████████                   
               █████████████████                
                   ███████████                  
                      █████                     `

const logoShiftCols = 1

func indentEachLine(s string, cols int) string {
	if cols <= 0 {
		return s
	}
	pad := strings.Repeat(" ", cols)
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = pad + lines[i]
	}
	return strings.Join(lines, "\n")
}

func fprintLogo(w io.Writer) {
	body := indentEachLine(strings.TrimPrefix(logo, "\n"), logoShiftCols)
	if !colorEnabled(w) {
		fmt.Fprint(w, body, "\n\n")
		return
	}
	fmt.Fprint(w, logoColor, body, logoReset, "\n\n")
}

func colorEnabled(w io.Writer) bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	fd := int(f.Fd())
	return term.IsTerminal(fd)
}
