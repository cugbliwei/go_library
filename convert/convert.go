package convert

import (
	"log"
	"strconv"

	iconv "github.com/djimenez/iconv-go"
)

func ConvertCharacterEncoding(msg, oldEncoding, newEncoding string) string {
	newMsg, err := iconv.ConvertString(msg, oldEncoding, newEncoding)
	if err != nil {
		log.Println("convert string Encoding error: ", err)
		return ""
	}
	return newMsg
}

func StringToInt(number string) int {
	num, _ := strconv.Atoi(number)
	return num
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func Float64ToString(number float64) string {
	num := strconv.FormatFloat(number, 'f', -1, 64) // -1可改为7，即保留7位小数
	return num
}

func StringToFloat64(number string) float64 {
	num, _ := strconv.ParseFloat(number, 64)
	return num
}
