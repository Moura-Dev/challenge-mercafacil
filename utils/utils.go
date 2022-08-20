package utils

import "fmt"

func MaskCellPhoneMacapa(cellphone string) string {
	if len(cellphone) == 13 {
		cellphone = fmt.Sprintf(`+%s (%s) %s-%s"`, cellphone[0:2], cellphone[2:4], cellphone[4:9], cellphone[9:])
	}

	return cellphone
}

func MaskCellPhoneVarejo(cellphone string) string {
	if len(cellphone) == 13 {
		cellphone = cellphone[0:4] + cellphone[5:13]
	}

	return cellphone
}

// 554130306905
