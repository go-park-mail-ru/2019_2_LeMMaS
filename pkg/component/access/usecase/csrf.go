package usecase

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/access"
	"io"
	"time"
)

const (
	CSRFTokenExpire = time.Minute * 30
	CSRFTokenSecret = "09dJ2e4hcM5Tot984E9WQ5ur8Nty7RT2"
)

type csrfUsecase struct {
}

func NewCSRFUsecase() access.CsrfUsecase {
	return csrfUsecase{}
}

type TokenData struct {
	Payload string
	Expires int64
}

func (u csrfUsecase) CreateTokenBySession(sessionID string) (string, error) {
	return u.createToken(sessionID, CSRFTokenExpire)
}

func (u csrfUsecase) CheckTokenBySession(token string, sessionID string) (bool, error) {
	return u.checkToken(token, sessionID)
}

func (u csrfUsecase) createToken(payload string, expire time.Duration) (string, error) {
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

	tokenExpTime := time.Now().Add(expire).Unix()
	td := &TokenData{Payload: payload, Expires: tokenExpTime}
	encodedData, _ := json.Marshal(td)
	ciphertext := aesgcm.Seal(nil, nonce, encodedData, nil)

	res := append([]byte(nil), nonce...)
	res = append(res, ciphertext...)

	token := base64.StdEncoding.EncodeToString(res)
	return token, nil
}

func (u csrfUsecase) checkToken(token string, payload string) (bool, error) {
	block, err := aes.NewCipher([]byte(CSRFTokenSecret))
	if err != nil {
		return false, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(token)
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
		return false, fmt.Errorf("token expired (valid until %v)", time.Unix(td.Expires, 0).String())
	}
	expected := TokenData{Payload: payload}
	td.Expires = 0
	return td == expected, nil
}
