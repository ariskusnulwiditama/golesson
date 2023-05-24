package star

import "fmt"

func SegitigaSama() {
	var tinggi int
	fmt.Print("masukkan tinggi segitiga: ")
	fmt.Scanln(&tinggi)

	for i:=1; i<=tinggi; i++ {
		for j:=1; j<=tinggi-1; j++ {
			fmt.Print(" ")
		}
		for k:=1; k<=2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}