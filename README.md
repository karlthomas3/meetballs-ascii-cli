# Welcome to the 'Go ASCII converter' Meetballs™ session!

This is just a simple CLI application built with go to convert text into ASCII art.<br>
I tried to keep it extremely short and simple (possibly **too** short) so it should be fun and easy even if you've never played with go before.<br>
If we get through it way too quickly we can play with adding new features on the fly 😁

## Setup

### Make sure you have go installed

-   `go version`
-   if you dont have go then:

-   `brew install go` or download from go.dev

### Start a new go project

-   `go mod init meetballs.com/go-cli`

### Instal dependencies

-   `go get github.com/common-nighthawk/go-figure`

## Lets get started

Create a file in the main directory called `main.go`

Inside main.go lets declare the package, imports, and the main function.

-   flag will be used to implement cli flags for different features
-   fmt gives us things like print
-   os does a lot but we'll be using it to exit the program
-   strings gives us all the usual string functions
-   go-figure is an ascii art library because we aren't gonna code that from scratch in two hours

```go
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func main(){}

```

We're using the flags package to parse commandline arguments and if no arguments are given then we'll provide usage instructions to the user.

```go
func main() {

flag.Parse()

args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ./ascii <string>")
		os.Exit(1)
	}
}
```

Now each space gets recognized as a new object so we need to get the strings not claimed by flags and combine them into one string

```go
input := strings.Join(args, " ")
```

Then from the go-figure library we're gonna use the NewFigure function to generate the ascii art.

It takes three arguments:

1. The text to be transformed
2. The font to use (it has a default if we dont set one)
3. A boolean for strict mode. True means it'll panic for unknown characters and false means unknowns will be replaced with "?"

Our text is coming from input, leave the font as "" for now

```go
myFigure := figure.NewFigure(input, "", false)
```

The NewFigure will come with a Print() method which will output the ascii art to the console so all we have to do is:

```go
myFigure.Print()
```

Now your code should look like this.

```go
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func main() {

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ./ascii <string>")
		os.Exit(1)
	}
	input := strings.Join(args, " ")
	myFigure := figure.NewFigure(input, "", false)
	myFigure.Print()
}
```

In the terminal use `go build` and go will automatically compile a binary for your machine and os.

Try running it with `./ascii Meetballs in your mouth!` and see for yourself.
If you're on windows you'll have to use `ascii.exe` instead of `./ascii`

## Now for fun with flags

We'll start by creating a flag to choose a font<br>
Add this line to the top of your `main()`

The flag has separate functions for different types of flags (i.e. String, Int, Bool) and takes three arguments.

1. the flag string to be typed by the user
2. the default value
3. a description to be displayed if the user uses `-help` or `-h` (We get this by default without having to implement it ourselves)

```go
fontFlag := flag.String("font", "", "Font to use for ASCII art")
```

Then change the second argument to our `figure.NewFigure()` call to `*font`

```go
myFigure := figure.NewFigure(input, *fontFlag, false)
```

Now type `go build` and try `./ascii -font stop salty meetballs!`<br>
If you go to the go-figure repo (where we're imorting from) you can find all the available fonts. <a href="https://github.com/common-nighthawk/go-figure/tree/master/fonts">go-figure/fonts/</a>

## Time!

First add `"time"` to your imports.
Then create a timeFlag just under your fontFlag.

```go
timeFlag := flag.Bool("time", false, "Display time in ASCII art")
```

Now underneath `flag.Parse()` we want to check for the timeFlag. If the time flag is used we get (and format) the current time, create the ASCII art with `figure.NewFigure()` print it, and return before running the rest of the code.

```go
if *timeFlag {
		currentTime := time.Now().Format("15:04:05")
		myFigure := figure.NewFigure(currentTime, *fontFlag, false)
		myFigure.Print()
		return
	}
```

If your code looks like this

```go
package main

import (
	"flag"
	"fmt"
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

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ./ascii [-font <font>] <string>")
		os.Exit(1)
	}
	input := strings.Join(args, " ")
	myFigure := figure.NewFigure(input, *fontFlag, false)
	myFigure.Print()
}
```

then `go build` and `./ascii -time`

## One last bit of fun!

What if you have a lot of text you want to ASCII-ify?
How about we allow you to pipe input into the tool?

Add "io" to the imports. Then replace everything under the `if *timeFlag` block with this. (We'll go over it in a second)

```go

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
			fmt.Println("Usage: ./ascii [-font <font>] <string>")
			os.Exit(1)
		}
		input = strings.Join(args, " ")
	}

	myFigure := figure.NewFigure(input, *fontFlag, false)
	myFigure.Print()
```

First we declare our input variable

```go
var input string
```

Then we check to see if anything has been piped in via standard input (Stdin)

```go
stat, _ := os.Stdin.Stat()
```

If there's piped input then we read it all as bytes with the "io" package (exiting early if there's an error), convert the bytes to a string and assign it to the `input` variable we just declared a moment ago.

If there's no piped input then we use the same logic as before to assign the input, convert to ASCII, and print it to the console.

Now your code should look something like this

```go
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
			fmt.Println("Usage: ./ascii [-font <font>] <string>")
			os.Exit(1)
		}
		input = strings.Join(args, " ")
	}

	myFigure := figure.NewFigure(input, *fontFlag, false)
	myFigure.Print()
}
```

and if you rebuild with `go build` you can pipe the output of basically anything as ASCII by typing `| ./ascii` (or `| ascii.exe` for windows) after the command.<br>
For instance `echo hello | ./ascii -font wavy`<br>
Or `cat someFile.txt | ./ascii`

### Bonus

If you follow your command with `>> filename.txt` you can save the output to a text file 😁

If you really want to go down the rabbit hole, <a href="cobra.dev">Cobra CLI</a> is a framework for building really powerful CLI applications in Go.
