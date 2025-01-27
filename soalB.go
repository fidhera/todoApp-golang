package main
import "fmt"

// Membuat struct untuk data Buku
type Buku struct {
	Judul string
	Penulis string
	harga float64
}
func main() {
	//Membuat instance struct
	buku := Buku{"KISAH KIAMAT", "RAFFAEL", "200000"}

	//Menampilkan data buku
	fmt.Println("Data Buku")
	fmt.Printf("Judul: %s\nPenulis: %s\nHarga: %f\n")
}
