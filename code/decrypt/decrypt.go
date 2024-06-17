package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// this function decrypt recursively all the file contain in the crypted folder
func Decrypt(folderPath string) {
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

			Decrypt(newPath)
			continue
		}

		datas, err := os.ReadFile(newPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		cryptedStr, _, err := DecryptString(string(datas))
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

		newFile.Write(cryptedStr)
	}
}

// this function decrypts a string which has been encrypted by this code
func DecryptString(sCrypt string) ([]byte, string, error) {
	return msgCrypter.Decrypt(sCrypt)
}
