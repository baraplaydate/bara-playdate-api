package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/mergermarket/go-pkcs7"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

func StringToInt(param string) int {

	intValue, err := strconv.Atoi(param)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Integer value:", intValue)
	}
	return intValue
}

func DateToStdNow() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func HmacEncode(msg, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return hex.EncodeToString(mac.Sum(nil))
}

func HmacDecode(msg, key []byte, hash string) (bool, error) {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return hmac.Equal(sig, mac.Sum(nil)), nil
}

func GeneratePassword() (string, error) {
	password := "p@ssw0rd"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func EncryptAes256(dataPayload string, password string, iv string) (string, error) {
	block, errChipper := aes.NewCipher([]byte(password))

	if errChipper != nil {
		return "", fmt.Errorf("decode input text is error. error: " + errChipper.Error())
	}

	if dataPayload == "" {
		return "", fmt.Errorf("data payload error")
	}

	ecb := cipher.NewCBCEncrypter(block, []byte(iv))

	content := []byte(dataPayload)
	content, errPad := pkcs7.Pad(content, block.BlockSize())
	if errPad != nil {
		return "", fmt.Errorf("decode input text is error. error: " + errPad.Error())
	}

	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return base64.StdEncoding.EncodeToString(crypted), nil

}

func DecryptAes256(dataPayload string, password string, iv string) ([]byte, error) {

	ciphertext, errDecode := base64.StdEncoding.DecodeString(dataPayload)
	if errDecode != nil {
		return nil, fmt.Errorf("decode input text is error. error: " + errDecode.Error())
	}

	block, errChipper := aes.NewCipher([]byte(password))

	if errChipper != nil {
		return nil, fmt.Errorf("key chipper is error : " + errChipper.Error())
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	blockMode.CryptBlocks(ciphertext, ciphertext)

	cipherUnPad, _ := pkcs7.Unpad(ciphertext, aes.BlockSize)

	return cipherUnPad, nil
}

// / Encrypt-Decrypt with current condition
func deriveKey(password, salt []byte, iterations, keySize int) []byte {
	return pbkdf2.Key(password, salt, iterations, keySize, sha256.New)
}

func EncryptAes256Sha256(dataPayload string, password string) (string, error) {
	// Create salt and hash from password
	uuid := uuid.New()
	salt := uuid[:]

	// Create IV from random 16 byte
	iv := make([]byte, 16)
	rand.Read(iv)
	byteIv := []byte(iv)

	// Create password from hash, salt, iteration, and how many bytes
	key := deriveKey([]byte(password), salt, 5, 32)
	block, errChipper := aes.NewCipher(key)

	if errChipper != nil {
		fmt.Println("key chipper is error :")
	}

	if dataPayload == "" {
		fmt.Println("input plain text is empty")
	}

	cbc := cipher.NewCBCEncrypter(block, byteIv)

	content := []byte(dataPayload)
	content, errPad := pkcs7.Pad(content, block.BlockSize())
	if errPad != nil {
		fmt.Println("padding input text padding is error")
	}

	crypted := make([]byte, len(content))
	cbc.CryptBlocks(crypted, content)

	crypted = append(append(salt, byteIv...), crypted...)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func DecryptAes256Sha256(dataPayload []byte, password string) (string, error) {
	// Get Hash from
	ciphertext, errDecode := base64.StdEncoding.DecodeString(string(dataPayload))
	if errDecode != nil {
		return "", fmt.Errorf("decode input text is error. error: " + errDecode.Error())
	}
	// Create salt and iv from encrypted text
	salt := ciphertext[:16]
	iv := ciphertext[16:32]
	ciphertext = ciphertext[32:]

	// Create key from hash, salt, iteration, and how many bytes
	key := deriveKey([]byte(password), salt, 5, 32)
	block, errChipper := aes.NewCipher(key)

	if errChipper != nil {
		return "", fmt.Errorf("key chipper is error : " + errChipper.Error())
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(ciphertext, ciphertext)

	cipherUnPad, err := pkcs7.Unpad(ciphertext, aes.BlockSize)
	if err != nil {
		return "", err
	}

	return string(cipherUnPad), nil
}

func DetectJSONOrString(s string) (interface{}, bool) {
	var js map[string]interface{}

	if json.Unmarshal([]byte(s), &js) == nil {
		return js, true // It's valid JSON
	}

	return s, false // It's not valid JSON, treat as a plain string
}
