package main

import (
	"assignment-one/function"
	"assignment-one/method"
	"assignment-one/model"
	"fmt"
)

func main() {

	// Console inout
	fmt.Print("Search by Name: ")
	var name string
	fmt.Scanln(&name)

	// Geerate Data
	data := function.GenerateData

	// Search Data and Print
	idx, searchedData := method.SearchData(data(), name)

	if searchedData == nil {
		fmt.Println("Data Not Found")
		return
	}

	mappedData := searchedData.(model.Biodata)
	fmt.Println("Index", idx)
	fmt.Println("Nama", mappedData.Nama)
	fmt.Println("Alamat", mappedData.Alamat)
	fmt.Println("Pekerjaan", mappedData.Pekerjaan)
	fmt.Println("Alasan", mappedData.Alasan)
}
