package main

import (
	"fmt"
	"os"

	greeting "github.com/NasSilverBullet/practicegolang/testtest"
)

func main() {
	var g greeting.Greeting
	err := g.Do(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
}
