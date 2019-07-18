package main

import (
	"fmt"
	"log"
	"os"
)



const BUFSIZE = 4096

func main() {
	if err :=exec("abc.txt"); err!=nil {
		log.Fatal()
	}
}

func exec(fp string) error {
	fmt.Println("hoge")
	f, err := os.Open(fp)
	if err!=nil {
		return err
	}

	fmt.Println("hoge")
	s, err := f.Stat()
	if err!=nil {
		return err
	}

	fmt.Println("hoge")
	start := int(s.Size() - BUFSIZE)
	//sep := []byte("\n")
	buf := make([]byte, BUFSIZE)

	defer f.Close()
	_, err = f.ReadAt(buf, int64(start))
	if err != nil {
		return err
	}
	fmt.Println("hoge")
	fmt.Println(string(buf))
	return nil
}
