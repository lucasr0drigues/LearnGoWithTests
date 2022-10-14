package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const countdownStart = 3

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