package main

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

// fmt.Println(dortISLEM(10))
// topla, cikar, carp, _ := dortISLEM(10)
// fmt.Println(topla, cikar, carp)
// metin := "Merhaba Dünya"

// func(a string) {
// 	fmt.Println(a)
// }(metin)

// for i := 0; i < 10; i++ {
// 	fmt.Println(i)
// }

// deger := 0
// for deger < 10 {
// 	fmt.Println("deger", deger)
// 	deger++
// }

// var a int = 8
// ekle(&a)
// fmt.Println(a)
// }

// func dortISLEM(sayi int) (x, y, a, d int) {
// 	x = sayi + 5
// 	y = sayi / 5
// 	a = sayi * 5
// 	d = sayi - 5
// 	return x, y, a, d
// }

// func ekle(v *int) {
// 	*v += 5
// }

func main() {

}

// ---------------------- struct  -------------------

// type user struct {
// 	name    string
// 	surname string
// 	age     int
// }

// user1 := user{"hasan furkan", "koprulu", 21}
// user1.name = "ihsan"
// user1.surname = "cetin"
// user1.age = 55
// fmt.Println(user1)

// user1 := user{}
// user1.name = "hasan"
// user1.surname = "furkan"
// user1.name, user1.surname = "hasan furkan", "koprulu"
// user1.age = 21

// user1 := user{age: 21, surname: "koprulu"}
// fmt.Println(user1)

// -------------------- anonim struct ------------
// user := struct {
// 	name, surname string
// }{"hasan", "furkan"}

// fmt.Println(user)

// ------------------ function struct ----------------

// type user struct {
// 	name, surname string
// 	age           int
// }

// func (i user) meeting() {
// 	fmt.Printf("Merhaba, Ben %s %s %d yaşındayım.", i.name, i.surname, i.age)
// }

// user1 := user{"hasan", "furkan", 21}
// 	user1.meeting()
