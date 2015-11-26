package file

import (
	"log"
	"os"
)

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
