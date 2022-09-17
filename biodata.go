package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	reader := os.Args[1]

	data1 := Biodata{
		Nama:      "Ivan",
		Alamat:    "Bandung",
		Pekerjaan: "Backend Engineer",
		Alasan:    "Mencari Ilmu",
	}

	data2 := Biodata{
		Nama:      "Vina agustina",
		Alamat:    "Purworejo",
		Pekerjaan: "Backend Engineer",
		Alasan:    "Mencari pengalaman",
	}

	data3 := Biodata{
		Nama:      "Ferdinan kuncoro aji",
		Alamat:    "Purwokerto",
		Pekerjaan: "Backend Engineer",
		Alasan:    "Melatih skill",
	}

	data4 := Biodata{
		Nama:      "Dion Fauzi",
		Alamat:    "Purbalingga",
		Pekerjaan: "Backend Engineer",
		Alasan:    "Mencari pengalaman ",
	}

	datas := map[int]Biodata{1: data1, 2: data2, 3: data3, 4: data4}

	data, _ := strconv.Atoi(reader)

	dataPeserta := datas[data]

	Response(dataPeserta)

}

func Response(data Biodata) {
	var newBiodata Biodata
	result := fmt.Sprintf("Nama : %s\nAlamat : %s\nPekerjaan : %s\nAlasan : %s", data.Nama, data.Alamat, data.Pekerjaan, data.Alasan)
	if data == newBiodata {
		result = "input harus angka 1 - 4"
	}
	fmt.Println(result)
}
