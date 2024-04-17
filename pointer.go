package main

import "fmt"

func main() {


	var i int = 10
	var p *int
	i = 100
	p = &i

	fmt.Println(i, &p, *p)
}
