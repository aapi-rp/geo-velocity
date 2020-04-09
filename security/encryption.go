package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/aapi-rp/geo-velocity/base"
	"github.com/aapi-rp/geo-velocity/logger"
	"io"
	"net/http"
)

func Encrypt(text string) string {
	key := []byte(base.EncKey())
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	byteData := []byte("0E&@w85hetEO7ry6")
	var r io.Reader
	r = bytes.NewReader(byteData)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(r, iv); err != nil {
		//panic(err)
		logger.Error(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func Decrypt(cryptoText string) (string, bool) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
	key := []byte("0E&@w85hetEO7ry6")
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error(err)
		return "", false
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		logger.Debug("ciphertext too short")
		return "", false
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), true
}

//ReadCookie read the cookie
func ReadCookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		logger.Error("error in reading cookie : ", err.Error())
		return ""
	}
	return c.Value
}
