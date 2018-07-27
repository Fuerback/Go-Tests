package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//@todo read dir name from stdin
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dir := scanner.Text()
	jsonify(dir)
}

func jsonify(dir string) {
	if _, err := os.Stat(dir); err == nil {
		fmt.Println("ok")
		fileName := createFile(dir)
		writeFile(fileName)
	}
	//return fmt.Errorf("Not implemented")
}

func createFile(dir string) (fileName string) {
	path := dir + "files.json"

	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, _ = os.Create(path)
		defer file.Close()
		return path
	}

	return
}

func writeFile(fileName string) {
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
	if isError(err) {
		return
	}

	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}
}

func getFilesDirectory(dir string) (files[]string) {
	return data := []string{"Amazonas", "Pará", "Mato Grosso", "Minas Gerais", "Bahia", "Mato Grosso do Sul", "Goiás",
		"Maranhão", "Rio Grande do Sul", "Tocantins"}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
