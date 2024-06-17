package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// this function encrypt recursively all the file contain in the uncrypted folder
func Encrypt(folderPath string) {
	folder, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range folder {
		newPath := filepath.Join(folderPath, file.Name())

		if file.IsDir() {
			err = os.Mkdir(strings.Replace(newPath, "previousFile", "transformFile", 1), os.ModePerm)
			if err != nil {
				fmt.Println(err)
				return
			}

			Encrypt(newPath)
			continue
		}

		datas, err := os.ReadFile(newPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		cryptedStr, err := CryptString(string(datas))
		if err != nil {
			fmt.Println(err)
			return
		}

		newFile, err := os.Create(strings.Replace(newPath, "previousFile", "transformFile", 1))
		if err != nil {
			fmt.Println(err)
			return
		}

		defer newFile.Close()

		newFile.WriteString(cryptedStr)
	}
}

// this function decrypts a string which has been encrypted by this code
func CryptString(sToCrypt string) (string, error) {
	return msgCrypter.Encrypt(sToCrypt)
}
