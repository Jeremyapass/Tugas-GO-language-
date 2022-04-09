package main

import "fmt"

func main()  {
	var name, minn, maxx string
	var p int
	var faktor, jumlah, min, mean, max float64
	minn = ""
	maxx = ""
	fmt.Scanln(&name)
	for nama != "STOP"{
		fmt.Scan(&p)
		for i := 0; i < 3*p; i++ {
			fmt.Scan(&faktor)
			jumlah += faktor
		}
		mean = jumlah / float64(3*p)
		if minn == "" || minn != "" && min > mean {
			min = mean 
			minn = nama 
		}
		if maxx == "" || maxx != "" && max < mean {
			max = mean
			maxx = nama
		}
		fmt.Scan(&nama)
		jumlah = 0
	}
	fmt.Printf("%s %2.f\n", maxx, max)
	fmt.Printf("%s %2.f", minn, min)
}