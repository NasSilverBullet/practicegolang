package greeting_test

import (
	"bytes"
	"errors"
	"io"
	"testing"
	"time"

	greeting "github.com/NasSilverBullet/practicegolang/testtest"
)

func mockClock(t *testing.T, v string) greeting.Clock {
	t.Helper()
	now, err := time.Parse("2006/01/02 15:04:05", v)
	if err != nil {
		t.Error("unexepected error:", err)
		return nil
	}
	return greeting.ClockFunc(func() time.Time {
		return now
	})
}

type errorWriter struct {
	Err error
}

func (w *errorWriter) Write(p []byte) (n int, err error) {
	return 0, w.Err
}

func TestGreeting_Do(t *testing.T) {

	cases := map[string]struct {
		writer    io.Writer
		clock     greeting.Clock
		msg       string
		expectErr bool
	}{
		"04時": {new(bytes.Buffer), mockClock(t, "2018/08/31 04:00:00"), "おはよう\n", false},
		"09時": {new(bytes.Buffer), mockClock(t, "2018/08/31 09:00:00"), "おはよう\n", false},
		"10時": {new(bytes.Buffer), mockClock(t, "2018/08/31 10:00:00"), "こんにちは\n", false},
		"16時": {new(bytes.Buffer), mockClock(t, "2018/08/31 16:00:00"), "こんにちは\n", false},
		"17時": {new(bytes.Buffer), mockClock(t, "2018/08/31 17:00:00"), "こんばんは\n", false},
		"03時": {new(bytes.Buffer), mockClock(t, "2018/08/31 03:00:00"), "こんばんは\n", false},
		"エラー": {&errorWriter{Err: errors.New("error")}, nil, "", true},
	}

	for n, tt := range cases {
		t.Run(n, func(t *testing.T) {
			g := greeting.Greeting{
				Clock: tt.clock,
			}
			switch err := g.Do(tt.writer); true {
			case err == nil && tt.expectErr:
				t.Error("expected error did not occur")
			case err != nil && !tt.expectErr:
				t.Error("unexpected error:", err)
			}
			if buf, ok := tt.writer.(*bytes.Buffer); ok {
				msg := buf.String()
				if msg != tt.msg {
					t.Errorf("greeting msg wont %s but got %s", tt.msg, msg)
				}
			}
		})
	}
	g := greeting.Greeting{
		Clock: mockClock(t, "2019/07/22 08:30:00"),
	}
	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "おはよう\n", buf.String(); expected != actual {
		t.Errorf("greeting message wont %s but got %s", expected, actual)
	}
}
