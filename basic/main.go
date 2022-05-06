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
	fmt.Println()

	a := 15
	if a%2 == 0 {
		fmt.Printf("sayi cift, kalan %v", a%2)
	} else {
		fmt.Printf("sayi tek, kalan %v", a%2)
	}

	fmt.Println()

	b := 10

	switch b {
	case 1:
		fmt.Printf("düşük %v", b)
	case 9:
		fmt.Printf("iyi %v", b)
	default:
		fmt.Printf("yorumsuz %v", b)
	}

}

func yaz() {
	fmt.Println(strGlobalScope)
	//fmt.Print(strFuncScope) // function scopetaki değişken kullanılamaz.
}
