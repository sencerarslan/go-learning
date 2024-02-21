package util

import (
	"fmt"
)

func ForeachLoop() {
	example5() // Bir slice içindeki her bir eleman için döngü oluşturuyor ve elemanları ekrana yazdırıyor.
	example6() // Bir dizedeki (string) her bir karakter için döngü oluşturuyor ve karakterleri ekrana yazdırıyor.
	example7() // Bir map içindeki her bir anahtar-değer çifti için döngü oluşturuyor ve anahtarları ve değerleri ekrana yazdırıyor.
}

func example5() {
	var numbers = []int{1, 2, 3, 4}
	for _, value := range numbers {
		fmt.Println(value)
	}
}
func example6() {
	var name = "Sencer"
	for _, c := range name {
		fmt.Printf("%c\n", c)
	}
}

func example7() {
	var names = map[string]int{
		"Sencer":  10,
		"Kubilay": 20,
		"Soner":   30,
	}
	for key, item := range names {
		fmt.Println(key, item)
	}
}
