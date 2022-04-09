package main

import "fmt"

func main() {
	var j, hasil int
	museum(&hasil, &j)
	fmt.Printf("Museum buka selama %d hari", j)
	fmt.Printf("\nRata-rata pengunjung %.2f orang", float64(hasil)/float64(j))
}

func museum(jumlah *int, i *int) {
	var turun int
	var orang, temp int
	*i = 1
	for turun != 3 {
		fmt.Printf("Hari %d: ", *i)
		fmt.Scanln(&orang)
		for orang < 0 || orang > 100 {
			fmt.Printf("Hari %d: ", *i)
			fmt.Scanln(&orang)
		}
		*jumlah = *jumlah + orang
		cek(orang, temp, &turun)
		temp = orang
		*i++
	}
}

func cek(orang, temp int, turun *int) {
	if orang < temp {
		*turun++
	} else {
		*turun = 0
	}
}
