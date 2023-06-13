package internal

import (
	"log"
	"strconv"
)

func StrToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(err)
		return 0
	}

	return number
}
