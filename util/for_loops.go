package util

import (
	"fmt"
)

func ForLoop() {
	example1() // 1'den 10'a kadar olan sayıların toplamını hesaplayan fonksiyonu çağırır.
	example2() // Bir dizideki öğeleri indeks kullanarak yazdıran fonksiyonu çağırır.
	example3() // Bir döngü içinde belirli bir koşul sağlandığında döngüyü kırarak çıkan fonksiyonu çağırır.
	example4() // Belirli koşullar sağlandığında döngüyü geçen fonksiyonu çağırır.
}

func example1() {
	total := 0
	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Println(total)
}

func example2() {

	index := 0
	var names = [3]string{"A", "B", "C"}

	for index < 3 {
		fmt.Println(names[index])
		index++
	}

}

func example3() {

	for i := 1; i <= 10; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

}
func example4() {

	for i := 1; i <= 10; i++ {
		if i == 3 || i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
