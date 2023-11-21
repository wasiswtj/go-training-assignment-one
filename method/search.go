package method

import (
	"assignment-one/model"
	"strings"
)

func SearchData(data []model.Biodata, nama string) (idx int, biodata interface{}) {
	for i, v := range data {
		if strings.ToLower(nama) == strings.ToLower(v.Nama) {
			return i, v
		}
	}

	return 0, nil
}
