package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := exec(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("End of main")
}

func exec() error {
	f, err := os.Open("")
	if err != nil {
		return err
	}
	defer f.Close()
	ch := make(chan []byte)
	done := make(chan bool)

	go printLine(f, ch, done)

	// ファイルRead用forループ
	if err := getLine(f, ch, done); err != nil {
		return err
	}
	return nil
}

func printLine(f *os.File, ch chan []byte, done chan bool) {
loop:
	for {
		select {
		case b, ok := <-ch:
			if !ok { // selectでchanのクローズを検知する方法
				fmt.Println("ch is closed")
				break loop
			}
			fmt.Print(string(b))
		}
	}
	fmt.Println("End of goroutine")
	done <- true
}

func getLine(f *os.File, ch chan []byte, done chan bool) error {
	for {
		buf := make([]byte, 32)
		if n, err := f.Read(buf); err != nil {
			if err != io.EOF {
				return err
			}
			break
		} else {
			ch <- buf[0:n]
		}
	}
	close(ch) // ファイル終端でchをcloseする

	// 非同期に実行されているゴルーチンの終了を待つ。
	<-done
	return nil
}
