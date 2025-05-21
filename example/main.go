package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	fontFlag := flag.String("font", "", "Font to use for ASCII art")
	timeFlag := flag.Bool("time", false, "Display time in ASCII art")

	flag.Parse()

	if *timeFlag {
		currentTime := time.Now().Format("15:04:05")
		myFigure := figure.NewFigure(currentTime, *fontFlag, false)
		myFigure.Print()
		return
	}

	var input string

	// Check if there's piped input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Read from stdin
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading stdin:", err)
			os.Exit(1)
		}
		input = strings.TrimSpace(string(bytes))
	} else {
		// Read from command-line arguments
		args := flag.Args()
		if len(args) < 1 {
			fmt.Println("Usage: [-font <font>] ./ascii <string>")
			os.Exit(1)
		}
		input = strings.Join(args, " ")
	}

	myFigure := figure.NewFigure(input, *fontFlag, false)
	myFigure.Print()
}
