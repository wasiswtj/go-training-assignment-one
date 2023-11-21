package main

import (
	"assignment-one/function"
	"assignment-one/method"
	"assignment-one/model"
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Console inout
	fmt.Print("Search by Name: [ Enter Name ]\n")
	var name string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		name = scanner.Text()
	}

	// Generate Data
	data := function.GenerateData()

	// Initiating populated data struct of search function
	populatedData := method.PopulatedData{Datas: data}

	// Search Data and Print
	idx, searchedData := populatedData.SearchData(name)

	if searchedData == nil {
		fmt.Println("Data Not Found")
		return
	}

	mappedData := searchedData.(model.Biodata)
	fmt.Println("Index", idx+1)
	fmt.Println("Nama", mappedData.Nama)
	fmt.Println("Alamat", mappedData.Alamat)
	fmt.Println("Pekerjaan", mappedData.Pekerjaan)
	fmt.Println("Alasan", mappedData.Alasan)
}
