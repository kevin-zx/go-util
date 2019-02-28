package main

import (
	"fmt"
	"github.com/kevin-zx/go-util/randomUtil"
)

func main() {
	//print(randomUtil.GetRandomInt(1,20))
	for i := 0; i < 100; i++ {
		fmt.Println(randomUtil.GetRandomInt(-20, 20))
	}
}
