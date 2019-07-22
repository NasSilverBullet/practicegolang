package greeting

import (
	"fmt"
	"io"
	"time"
)

type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

type Greeting struct {
	Clock Clock
}

func (g *Greeting) now() time.Time {
	if g.Clock != nil {
		return g.Clock.Now()
	}
	return time.Now()
}

func (g *Greeting) Do(w io.Writer) error {
	h := g.now().Hour()
	var msg string
	switch {
	case h >= 4 && h <= 9:
		msg = "おはよう\n"
	case h >= 10 && h <= 16: /* 10時00分 〜 16時59分 */
		msg = "こんにちは\n"
	default:
		msg = "こんばんは\n"
	}
	_, err := fmt.Fprint(w, msg)
	if err != nil {
		return err
	}
	return nil
}
