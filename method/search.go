package method

import (
	"assignment-one/model"
)

func SearchData(data []model.Biodata, idx int) interface{} {
	for i, v := range data {
		if idx == i {
			return v
		}
	}

	return nil
}
