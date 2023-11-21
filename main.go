package main

import (
	"assignment-one/function"
	"assignment-one/method"
	"assignment-one/model"
	"fmt"
	"strconv"
)

func main() {

	// Console inout
	fmt.Print("Enter text: ")
	var idx string
	fmt.Scanln(&idx)

	idxInt, err := strconv.Atoi(idx)
	if err != nil {
		fmt.Println("Index Should Be Number")
		return
	}

	// Geerate Data
	data := function.GenerateData

	// Search Data and Print
	searchedData := method.SearchData(data(), idxInt)

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
