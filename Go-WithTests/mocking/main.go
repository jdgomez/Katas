package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	_, _ = fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
