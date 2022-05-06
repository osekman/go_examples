package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	basic()
	basic2("ali", "veli")
	fmt.Printf("sonuc > %v", basic3(5, 3))

	read()
}

func read() {
	fmt.Println("\nsayi gir")
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("read err %v", err)
	} else {
		fmt.Printf("girilen  = %v", value)
	}
}
func basic() {
	fmt.Println("basic ")
}

func basic2(str1, str2 string) {
	fmt.Printf("basic2 %v, %v, ", str1, str2)
}

func basic3(x, y float32) float32 {

	fmt.Printf("basic3 degerler %v, %v, ", x, y)
	return x / y
}
