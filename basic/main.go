package main

import "fmt"

var strGlobalScope = "global scopta değişken"

func main() {
	var strFuncScope = "function scope"
	fmt.Println(strFuncScope)
	yaz()

	str, sayi, kosul := "ali", 12, false // çoklu değişken tanımlama

	fmt.Printf("%#v %T, ", str, str)
	fmt.Printf("%#v %T, ", sayi, sayi)
	fmt.Printf("%#v %T, ", kosul, kosul)
}

func yaz() {
	fmt.Println(strGlobalScope)
	//fmt.Print(strFuncScope) // function scopetaki değişken kullanılamaz.
}
