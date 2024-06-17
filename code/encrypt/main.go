package main

import (
	crypter "github.com/KonishiLee/wechat-crypter"
)

var (
	msgCrypter crypter.MessageCrypter
)

// initialise the key use for the cryptage
func init() {
	encodingAESKey := "MX53u9J3AbE0uLxjrD5GdjEGWZsU8bSCLTcadlYY6tZ"
	msgCrypter, _ = crypter.NewMessageCrypter("", encodingAESKey, "")
}

func main() {
	Encrypt("previousFile")
}
