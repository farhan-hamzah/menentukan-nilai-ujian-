package main
import "fmt"
const NMax int = 100
type tabNilai [NMax]int
func main(){
	var n int
	var nilai tabNilai
	fmt.Scan(&n)
	isiAray(n, &nilai)
}
func isiAray(n int, A *tabNilai){
	var i int
	var cek, avgS1, avgS2 float64
	if n%2 != 0{
		fmt.Print("Jumlah array invalid !!!")
		return
	}else{
		for i = 0; i < n; i++{
			fmt.Scan(&A[i])
		}
	}
	avgS1 = avgSesi1(n, *A)
	avgS2 = avgSesi2(n, *A)
	fmt.Println("Rata rata sesi 1 : ", avgS1)
	fmt.Println("Rata rata sesi 2 : ", avgS2)
	cek = avgS1-avgS2
	if cek < 0{
		cek = -cek
	}
	cekKecurangan(cek)
}
func cekKecurangan(cek float64){
	if cek >= 20{
		fmt.Println("Kemungkinan terjadi kecurangan")
	}else{
		fmt.Println("Tidak terjadi kecurangan")
	}
}
func avgSesi1(n int, A tabNilai)float64{
	var i, c int
	var rerata float64
	c = n/2
	for i = 0; i < c; i++{
		rerata += float64(A[i])
	}
	return rerata/float64(c)
}
func avgSesi2(n int, A tabNilai)float64{
	var i, c int
	var rerata float64
	c = n/2
	for i = 0; i < n; i++{
		if i >= c{
			rerata += float64(A[i])
		}
	}
	return rerata/float64(c)
}
