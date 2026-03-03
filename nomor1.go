package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	soal1()
}

func soal1() {
	var myNum int32 = 50
	var myNum2 float32 = 51
	var myNumStr string = "50"
	fmt.Println("Nilai myNum:", myNum)
	fmt.Println("Nilai myNum2:", myNum2)
	fmt.Println("Nilai myNumStr:", myNumStr)
}

func soal2() {
	var x int32 = 5
	var y int32 = 10
	var z int32 = x + y
	fmt.Println("Hasil penjumlahan x dan y:", z)
}
func soal3() {
	var nama string

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	fmt.Println("hello", nama)
}
func soal4() {
	var people []string = []string{"Walter", "Jesse", "Skyler", "Saul"}
	fmt.Print("panjang people:", len(people))
	people = append(people, "Hank", "Marie")
	fmt.Print("\nMenambahkan Hank dan Marie ke dalam people")
	fmt.Print("\nPanjang People:", len(people))
	fmt.Println("\npeople:", people)
}

func soal5() {
	people := make([]map[string]string, 0, 3)

	people = append(people, map[string]string{
		"name":   "Hank",
		"gender": "M",
	})

	people = append(people, map[string]string{
		"name":   "Heisenberg",
		"gender": "M",
	})

	people = append(people, map[string]string{
		"name":   "Skyler",
		"gender": "F",
	})

	fmt.Println("Data awal:")
	for _, person := range people {
		fmt.Println(person)
	}

	for i, person := range people {
		if person["gender"] == "M" {
			people[i]["name"] = "Mr " + person["name"]
		} else if person["gender"] == "F" {
			people[i]["name"] = "Mrs " + person["name"]
		}
	}

	fmt.Println("\nData setelah modifikasi:")
	for _, person := range people {
		fmt.Println(person)
	}
}

func soal6() {
	var nama string

	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)

	rand.Seed(time.Now().UnixNano())

	angka := rand.Intn(100) + 1

	fmt.Println("Angka yang didapat:", angka)

	if angka > 80 {
		fmt.Printf("Selamat %s, anda sangat beruntung\n", nama)
	} else if angka > 60 {
		fmt.Printf("Selamat %s, anda beruntung\n", nama)
	} else if angka > 40 {
		fmt.Printf("Mohon maaf %s, anda kurang beruntung\n", nama)
	} else {
		fmt.Printf("Mohon maaf %s, anda sial\n", nama)
	}
}

func soal7() {
	var umur int

	fmt.Print("Masukkan Umur: ")
	fmt.Scanln(&umur)

	if umur < 0 {
		fmt.Printf("Silahkan masukan umur yang valid\n")
	} else if umur < 18 {
		fmt.Printf("Dilarang masuk, maksimal umur 18 tahun\n")
	} else if umur > 100 {
		fmt.Printf("Silahkan masukan umur yang valid\n")
	} else {
		fmt.Printf("Selamat datang, umur anda %d tahun\n", umur)
	}
}

func soal8() {
	people := []map[string]interface{}{
		{
			"name": "Hank",
			"Age":  50,
			"Job":  "Polisi",
		},
		{
			"name": "Heisenberg",
			"Age":  52,
			"Job":  "Ilmuwan",
		},
		{
			"name": "Skyler",
			"Age":  48,
			"Job":  "Akuntan",
		},
	}

	for _, person := range people {
		fmt.Printf(
			"Hi Perkenalkan, Nama saya %s, umur saya %d, dan saya bekerja sebagai %s\n",
			person["name"],
			person["Age"],
			person["Job"],
		)
	}
}

func soal9() {
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	sum1 := 0.0
	for _, val := range slice1 {
		sum1 += val
	}
	avg1 := sum1 / float64(len(slice1))

	median1 := calculateMedian(slice1)

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Jumlah:", sum1)
	fmt.Println("Rata-rata:", avg1)
	fmt.Println("Median:", median1)
	fmt.Println("----------------------")

	sum2 := 0.0
	for _, val := range slice2 {
		sum2 += val
	}
	avg2 := sum2 / float64(len(slice2))

	median2 := calculateMedian(slice2)

	fmt.Println("Slice 2:", slice2)
	fmt.Println("Jumlah:", sum2)
	fmt.Println("Rata-rata:", avg2)
	fmt.Println("Median:", median2)
}

func calculateMedian(slice []float64) float64 {
	s := make([]float64, len(slice))
	copy(s, slice)
	n := len(s)
	if n%2 == 1 {
		return s[n/2]
	} else {
		return (s[n/2-1] + s[n/2]) / 2
	}
}

func soal10() {
	var kata string

	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&kata)

	kata = strings.ToLower(strings.ReplaceAll(kata, " ", ""))

	if isPalindrome(kata) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func soal11() {
	var kata string

	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&kata)

	kata = strings.ToLower(kata)

	countX := 0
	countO := 0

	for _, ch := range kata {
		if ch == 'x' {
			countX++
		} else if ch == 'o' {
			countO++
		}
	}

	fmt.Println("Jumlah x:", countX)
	fmt.Println("Jumlah o:", countO)

	if countX == countO {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func soal12() {

	nums := []int{5, 2, 9, 1, 7, 3}

	fmt.Println("Slice sebelum diurutkan:", nums)

	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	fmt.Println("Slice setelah diurutkan (besar ke kecil):", nums)
}

// Fungsi utama: menerima argumen tak terbatas dan menggabungkannya menjadi satu kalimat alay
func AlayGen(words ...string) string {
	alayWords := make([]string, len(words))
	for i, word := range words {
		alayWords[i] = toAlay(word) // ubah tiap kata menjadi alay
	}
	return strings.Join(alayWords, " ")
}

// Fungsi modular: ubah kata menjadi versi alay
func toAlay(word string) string {
	replacer := strings.NewReplacer(
		"a", "4",
		"e", "3",
		"i", "!",
		"l", "1",
		"n", "N",
		"s", "5",
		"x", "*",
	)
	return replacer.Replace(word)
}

func soal13() {
	result := AlayGen("hello", "world", "amazing")
	fmt.Println(result)
}

// Fungsi untuk mendapatkan bilangan Fibonacci ke-n
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	// Inisialisasi dua bilangan pertama
	prev, curr := 0, 1

	// Hitung Fibonacci iteratif
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

func soal14() {
	n := 10
	fmt.Printf("Fibonacci ke-%d adalah %d\n", n, Fibonacci(n))
}
