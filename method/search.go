package method

import (
	"assignment-one/model"
	"strings"
)

type PopulatedData struct {
	Datas []model.Biodata
}

func (p PopulatedData) SearchData(nama string) (idx int, biodata interface{}) {
	for i, v := range p.Datas {
		if strings.ToLower(nama) == strings.ToLower(v.Nama) {
			return i, v
		}
	}

	return 0, nil
}
