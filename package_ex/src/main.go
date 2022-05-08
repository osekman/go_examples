package main

//local paketleri import etmek için
//go mod init package_name, komutu ile go.mod dosyasının oluşmasını sağlıyoruz.
import (
	"fmt"
	"package_ex/src/mycelc"
)

func main() {
	fmt.Printf("cc %v", mycelc.GetCelc(20))
}
