package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/aapi-rp/geo-velocity/config"
	"github.com/aapi-rp/geo-velocity/logger"
	"io"
	"log"
)

// Long ago I used multiple sources to put these encryption and decryption methods together.
// Stack overflow mostly https://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64
// https://golang.org/pkg/crypto/ is another site i looked at for understanding

func Encrypt(text string) string {
	key, _ := hex.DecodeString(config.EncKey256())
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error("Error generating block: ", err)
	}

	byteData := []byte(config.EncIV())
	var r io.Reader
	r = bytes.NewReader(byteData)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(r, iv); err != nil {
		logger.Error("Error adding iv: ", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

func Decrypt(cryptoText string) (string, bool) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
	key, _ := hex.DecodeString(config.EncKey256())
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return "", false
	}

	if len(ciphertext) < aes.BlockSize {
		logger.Error("Improper cipher text size.")
		return "", false
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), true
}
