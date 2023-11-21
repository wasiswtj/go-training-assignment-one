package function

import "assignment-one/model"

func GenerateData() []model.Biodata {
	return []model.Biodata{
		{Nama: "Nama Satu", Alamat: "Alamat Satu", Pekerjaan: "Pekerjaan Satu", Alasan: "Alasan Satu"},
		{Nama: "Nama Dua", Alamat: "Alamat Dua", Pekerjaan: "Pekerjaan Dua", Alasan: "Alasan Dua"},
	}
}
