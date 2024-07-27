package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func writeFile(filename, text string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Could not create file ", err)
	}
	defer file.Close()
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	
	if err != nil {
		return "", fmt.Errorf("Could not read file", err)
	}
	
	return string(data), nil
}

func main() {
	filename := "test.txt"
	//text := "This is a test"
	
	// err := writeFile(filename, text)
	data, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(data)
}