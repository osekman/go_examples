package main

import "fmt"

var strGlobalScope = "global scopta değişken"

func main() {
	var strFuncScope = "function scope"
	fmt.Println(strFuncScope)
	yaz()

	str, sayi, kosul := "ali", 12, false // çoklu değişken tanımlama

	const x = 3
	const y = 5.8

	fmt.Printf("%#v %T, ", str, str)
	fmt.Printf("%#v %T, ", sayi, sayi)
	fmt.Printf("%#v %T, ", kosul, kosul)

	fmt.Println()
	fmt.Println("--- typless const ---")
	fmt.Printf("%#v %T, %#v %T, ", x, x, y, y)
	fmt.Println()
	fmt.Printf("%#v %T, ", x+y, x+y)
}

func yaz() {
	fmt.Println(strGlobalScope)
	//fmt.Print(strFuncScope) // function scopetaki değişken kullanılamaz.
}
