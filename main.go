package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/lestrrat/go-strftime"
)

const version = "1.0.0"

var revision = "HEAD"

func main() {
	var (
		fd          *os.File
		err         error
		format      = flag.String("f", "%Y-%m-%d %H:%M:%S", "format")
		showVersion = flag.Bool("version", false, "show version")
	)

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `
Usage:
  ts [-f format] [-V]
  ts --help

Options:
  -f format
    default is "%Y-%m-%d %H:%M:%S"

  -version
    output version information and exit

  --help
    display this help and exit
`)
	}

	flag.Parse()

	if *showVersion {
		fmt.Printf(
			"ts %s (rev: %s/%s)\n",
			version,
			revision,
			runtime.Version(),
		)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		fd = os.Stdin
	} else {
		fd, err = os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer fd.Close()
	}

	ts, err := strftime.New(*format)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		fmt.Printf("%s %s\n", ts.FormatString(time.Now()), scanner.Text())
	}

}
