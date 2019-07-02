package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "GO!"
const countdownStart = 3

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

// Countdown "It's the final countdown!"
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

// ConfigurableSleeper configs the sleep
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep method for ConfigurableSleeper
func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

// SpyTime struct
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep implementation for SpyTime
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// CountdownOperationsSpy spying on the countdown
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep recording Sleep operations
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write recording Write operations
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// SpySleeper our mock for testing
type SpySleeper struct {
	Calls int
}

// Sleep method for SpySleeper
func (s *SpySleeper) Sleep() {
	s.Calls++
}
