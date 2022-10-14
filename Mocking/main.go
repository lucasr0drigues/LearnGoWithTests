package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func main() {
	Countdown(os.Stdout)
}

func Countdown(out io.Writer){ 
	for i := countdownStart; i > 0; i--{
		fmt.Fprintln(out,i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(out, finalWorld)
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}