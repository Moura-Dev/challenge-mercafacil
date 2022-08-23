package utils

import "fmt"

func MaskPhone(cellphone string, customer string) string {
	switch customer {
	case "macapa":

		if len(cellphone) == 13 {
			cellphone = fmt.Sprintf(`+%s (%s) %s-%s`, cellphone[0:2], cellphone[2:4], cellphone[4:9], cellphone[9:])
			return cellphone
		}
	case "varejao":

		if len(cellphone) == 13 {
			cellphone = cellphone[0:4] + cellphone[5:13]
			return cellphone
		}

		return cellphone

	}
	return cellphone
}
