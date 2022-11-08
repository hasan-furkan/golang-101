package main

import "fmt"

// func main() {
// 	fmt.Printf("hello world")
// }

// func main() {
// 	değişken := "bir"
// 	{
// 		değişken := "iki"
// 		fmt.Println(değişken)
// 	}
// 	fmt.Println(değişken)
// }

// func main() {
// değişken := "bir"
// {
// 	değişken := "iki"
// 	fmt.Println(değişken)
// 	fmt.Println(&değişken)
// }
// fmt.Println(değişken)
// fmt.Println(&değişken)

// 	deneme := string(8378)
// 	fmt.Println(deneme)
// }

// func topla(a int, b int) int {
// 	return a + b //a ve b’nin toplamını döndürür.
// }

// func main() {
// 	fmt.Println(topla(2, 5)) //2+5 sonucunu ekrana bastır
// }

func dortISLEM(sayi int) (x, y, a, d int) {
	x = sayi + 5
	y = sayi / 5
	a = sayi * 5
	d = sayi - 5
	return x, y, a, d
}

func main() {
	// fmt.Println(dortISLEM(10))
	// topla, cikar, carp, _ := dortISLEM(10)
	// fmt.Println(topla, cikar, carp)
	// metin := "Merhaba Dünya"

	// func(a string) {
	// 	fmt.Println(a)
	// }(metin)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	deger := 0
	for deger < 10 {
		fmt.Println("deger", deger)
		deger++
	}
}
