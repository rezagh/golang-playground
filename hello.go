package main

import "fmt"

//var j int = 5
//var I int = 5 //capital global
// foo := 55

func main() {
	var (
		i int = 555
	)
	fmt.Println(i)
	var j = float32(i)
	fmt.Print(j)

	foo := 44
	fmt.Print(foo)
}
