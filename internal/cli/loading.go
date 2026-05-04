package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// ANSI: erase entire line + cursor to column 0 (Windows Terminal / modern conhost).
const ansiClearLine = "\x1b[2K\r"
const ansiReset = "\x1b[0m"

// wallbitDotRGB — brand blues from web/app/globals.css (wallbit-500 / 400 / 300).
var wallbitDotRGB = [3][3]uint8{
	{13, 153, 255},  // #0d99ff
	{71, 176, 255},  // #47b0ff
	{127, 200, 255}, // #7fc8ff
}

// Braille frames — shape changes every tick (visible motion); used for mono and color rows.
var brailleFrames = []rune{
	'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏',
}

func writerIsCharDevice(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func useANSIColor() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	if strings.EqualFold(os.Getenv("TERM"), "dumb") {
		return false
	}
	return true
}

func ansiTrueColorFG(r, g, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

// coloredWallbitBraille: one braille cell; shape and Wallbit hue both advance per frame.
func coloredWallbitBraille(frame int) string {
	n := len(brailleFrames)
	ch := brailleFrames[frame%n]
	rgb := wallbitDotRGB[frame%len(wallbitDotRGB)]
	return fmt.Sprintf("%s%s%c%s", ansiClearLine, ansiTrueColorFG(rgb[0], rgb[1], rgb[2]), ch, ansiReset)
}

func monoBrailleLine(frame int) string {
	ch := brailleFrames[frame%len(brailleFrames)]
	return fmt.Sprintf("%s%c", ansiClearLine, ch)
}

func flushWriter(w io.Writer) {
	if f, ok := w.(*os.File); ok {
		_ = f.Sync()
	}
}

// runWithLoading draws animated dots on stderr while fn runs, then clears the line.
func runWithLoading(w io.Writer, fn func() error) error {
	useTTY := writerIsCharDevice(w)
	if !useTTY {
		return fn()
	}

	color := useANSIColor()
	stop := make(chan struct{})
	done := make(chan struct{})

	go func() {
		defer close(done)
		i := 0
		paint := func() {
			if color {
				_, _ = fmt.Fprint(w, coloredWallbitBraille(i))
			} else {
				_, _ = fmt.Fprint(w, monoBrailleLine(i))
			}
			flushWriter(w)
		}
		paint()
		ticker := time.NewTicker(90 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-stop:
				_, _ = fmt.Fprint(w, ansiClearLine)
				flushWriter(w)
				return
			case <-ticker.C:
				i++
				paint()
			}
		}
	}()

	err := fn()
	close(stop)
	<-done
	return err
}
