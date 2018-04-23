package main

import "fmt"

func main() {
	generator := &IDGenerator{}
	generator.Init()

	for i := 0; i < 8; i++ {
		fmt.Println(generator.GetNewId())
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}