package main

import (
	"cq/logging"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	defaultLines     = 10
	defaultCellWidth = 10
)

func main() {
	l.New(log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime))

	var err error
	_ = flag.Uint64("n", defaultLines, "number of lines")
	_ = flag.Uint64("w", defaultCellWidth, "cell width")
	_ = flag.Bool("ln", false, "display line numbers")
	l.Enabled = flag.Bool("debug", false, "enable debug mode")

	flag.Parse()

	args := flag.Args()
	source := os.Stdin
	switch len(args) {
	case 0:
		fmt.Println("This is help message")
	case 1, 2:
		if len(args) == 2 {
			source, err = os.Open(args[1])
			if err != nil {
				panic(err)
			}
			defer source.Close()
		}

		err = processSource(source, args[0])
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Errorf("unrecognize arguments %v", args))
	}

}

func processSource(source *os.File, query string) error {
	fmt.Println("doing work")

	flag.VisitAll(func(f *flag.Flag) {
		l.Debug(fmt.Sprintf("flag: %s, value: %s", f.Name, f.Value))
	})
	l.Debug(fmt.Sprintf("query: %s", query))
	l.Debug(fmt.Sprintf("source: %v", source.Name()))
	return nil
}
