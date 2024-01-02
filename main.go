package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/shu-go/elapsed"
	"github.com/shu-go/gli/v2"
)

type globalCmd struct {
	Day         bool `cli:"day"`
	Hour        bool `cli:"hour"`
	Minute      bool `cli:"minute,min,m"`
	Second      bool `cli:"second,sec,s"`
	MilliSecond bool `cli:"milli-second,ms"`
	MicroSecond bool `cli:"micro-second,micro"`
	NanoSecond  bool `cli:"nano-second,nano,ns"`
}

func (c globalCmd) Run(args []string) error {
	if len(args) == 0 {
		return nil
	}

	ts := elapsed.Start()
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	delta := ts.Elapsed()
	s := ""

	switch true {
	case c.Day:
		s = fmt.Sprintf("%fdays", float64(delta/time.Hour)/24)
	case c.Hour:
		s = fmt.Sprintf("%fhours", float64(delta/time.Hour))
	case c.Minute:
		s = fmt.Sprintf("%fminutes", float64(delta/time.Minute))
	case c.Second:
		s = fmt.Sprintf("%fs", float64(delta.Nanoseconds())/float64(time.Second))
	case c.MilliSecond:
		s = fmt.Sprintf("%fms", float64(delta.Nanoseconds())/float64(time.Millisecond))
	case c.MicroSecond:
		s = fmt.Sprintf("%fµs", float64(delta.Nanoseconds())/float64(time.Microsecond))
	case c.NanoSecond:
		s = fmt.Sprintf("%dns", delta.Nanoseconds())
	default:
		s = fmt.Sprintf("%v", delta)
	}

	fmt.Fprintf(os.Stderr, "elapsed: %s\n", s)

	return nil
}

// Version is app version
var Version string

func main() {
	app := gli.NewWith(&globalCmd{})
	app.Name = "timeit"
	app.Desc = "outputs elapsed time to stderr"
	app.Version = Version
	app.Usage = `timeit CMD [ARGS...]`
	app.Copyright = "(C) 2024 Shuhei Kubota"
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
