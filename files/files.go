package files

import (
	"fmt"
	"os"
)

var fileName string = "./files/txt/data.txt"

func CreateFile(data string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file: ", err.Error())
		return
	}
	fmt.Fprintln(file, data)
	file.Close()

}

func AddDataToFile(data string) {
	if !AppendDataToFile(fileName, data) {
		fmt.Println("Error adding data to file")
	}
}

func AppendDataToFile(fileName string, data string) bool {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err.Error())
		return false
	}
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing data to file: ", err.Error())
		return false
	}
	file.Close()
	return true
}

func ReadFile() {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err.Error())
		return
	}
	fmt.Println(string(file))
}
