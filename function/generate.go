package function

import "assignment-one/model"

func GenerateData() []model.Biodata {
	return []model.Biodata{
		{Nama: "NamaSatu", Alamat: "Alamat Satu", Pekerjaan: "Pekerjaan Satu", Alasan: "Alasan Satu"},
		{Nama: "NamaDua", Alamat: "Alamat Dua", Pekerjaan: "Pekerjaan Dua", Alasan: "Alasan Dua"},
	}
}
