package main

import "fmt"

type fam []string
type deb []string

func main() {
	var tfam fam
	var tdeb deb
	if tfam == tdeb {
		fmt.Println("Iguales")
	}
}
