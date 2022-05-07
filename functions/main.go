package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	basic()
	basic2("ali", "veli")
	fmt.Printf("sonuc > %v", basic3(5, 3))

	read()

	sn, err := basic4(1, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("sonuc > %v", sn)
	}
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

func basic4(x, y float32) (float32, error) {

	fmt.Printf("basic3 degerler %v, %v, ", x, y)
	if x/y < 1 {
		return 0, errors.New("x y den küçük")
	}
	return x / y, nil
}
