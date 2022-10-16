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

type DefaultSleeper struct{}

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper){ 
	for i := countdownStart; i > 0; i--{
		fmt.Fprintln(out,i)
		sleeper.Sleep()
	}

	fmt.Fprintf(out, finalWorld)
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

func (d *DefaultSleeper) Sleep(){
	time.Sleep(1 * time.Second)
}

func (s *SpyCountdownOperations) Sleep(){
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error){
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"