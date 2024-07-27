package usecase

import (
	"strconv"
)

func NewValidateCep(cep string) bool {

	if len(cep) != 8 {
		return false
	}

	_, err := strconv.Atoi(cep)

	return err == nil

	// if err != nil {
	// 	return false
	// }

	// return true
}
