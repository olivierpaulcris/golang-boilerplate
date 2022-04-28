package main

import (
	"strconv"
	"fmt"
)

func main(){

	var edad int

	edad = 12

	nombre := "paul"

	edad_str := strconv.Itoa(edad)

	fmt.Println(edad , nombre, edad_str)
}
