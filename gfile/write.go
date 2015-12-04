package file

import (
	"log"
	"os"
)

//把字符串写到文件中
func WriteToFile(filename, content string) error {
	ret, err := os.Create(filename)
	defer ret.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	ret.WriteString(content)
	return nil
}
