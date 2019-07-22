package greeting

import (
	"bytes"
	"testing"
	"time"
)

func TestGreeting_Do(t *testing.T) {
	g := Greeting{
		Clock: ClockFunc(func() time.Time {
			return time.Date(2019, 7, 22, 5, 59, 59, 59, time.Local)
		}),
	}
	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "おはよう\n", buf.String(); expected != actual {
		t.Errorf("greeting message wont %s but got %s", expected, actual)
	}
}
