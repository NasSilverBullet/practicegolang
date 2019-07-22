package main

import (
	"log"
	"os"

	greeting "github.com/NasSilverBullet/practicegolang/testtest"
)

func main() {
	var g greeting.Greeting
	err := g.Do(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
