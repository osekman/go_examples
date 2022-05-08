package main

//local paketleri import etmek için
//go mod init package_name, komutu ile go.mod dosyasının oluşmasını sağlıyoruz.
//ps klasörü içinde iki tane p1.go ve p2.go isimli dosyalar var, paket isimleri ps olduğu icin aynı paketin farklı dosyaları gibi kullanıyoz
//ps in içindeki mycelc ise ayrı bir klasörde olduğu için farklı bir paket ismi alıyor. dosya ismi önemli değil, dosyanın başındaki paket adı ile bulunduğu klasör aynı olmalı, javadaki gibi
import (
	"fmt"
	"package_ex/src/ps"
	"package_ex/src/ps/mycelc"
)

func main() {
	fmt.Printf("cc %v \n", mycelc.GetCelc(20))
	fmt.Printf("conv SecondToMinute %v => %v \n", 600, ps.SecondToMinute(600))
	fmt.Printf("conv MinuteToSecond %v => %v \n", 10, ps.MinuteToSecond(10))
}
