package main

import "github.com/kevin-zx/go-util/randomUtil"
func main()  {
	//print(randomUtil.GetRandomInt(1,20))
	for i:=0; i<100;i++  {
		rnum,_ :=randomUtil.GetRandomInt(5,20)
		print(rnum,",")
	}
}

