package main

import (
	"fmt"

	"github.com/NasSilverBullet/practicegolang/testinterface/user"
)

func main() {
	u := user.New("taro", "tanaka", "male", "19941012")
	fmt.Println(u)
}
