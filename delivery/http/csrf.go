package http

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type TokenData struct {
	SessionID string
	Expires   int64
}

const (
	CSRFTokenSecret = "31Sht?U<up-~f=>@3y8sah3uwA?T(<E8gE92vh4]rs4M3%2EbX,u9SqCk6jQ)}-J"
	CSRFTokenExpire = time.Minute * 30
)

func createCSRFToken(sessionID string) (string, error) {
	block, err := aes.NewCipher([]byte(CSRFTokenSecret))
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	tokenExpTime := time.Now().Add(CSRFTokenExpire).Unix()
	td := &TokenData{SessionID: sessionID, Expires: tokenExpTime}
	data, _ := json.Marshal(td)
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	res := append([]byte(nil), nonce...)
	res = append(res, ciphertext...)

	token := base64.StdEncoding.EncodeToString(res)
	return token, nil
}

func checkCSRFToken(sessionID string, inputToken string) (bool, error) {
	block, err := aes.NewCipher([]byte(CSRFTokenSecret))
	if err != nil {
		return false, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(inputToken)
	if err != nil {
		return false, err
	}
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return false, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, fmt.Errorf("decrypt fail: %v", err)
	}

	td := TokenData{}
	err = json.Unmarshal(plaintext, &td)
	if err != nil {
		return false, fmt.Errorf("bad json: %v", err)
	}

	if td.Expires < time.Now().Unix() {
		return false, fmt.Errorf("token expired")
	}

	expected := TokenData{SessionID: sessionID}
	td.Expires = 0
	return td == expected, nil
}
