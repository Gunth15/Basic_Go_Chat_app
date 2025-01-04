package cookies

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/chat_app/pkg/database"
)

// Set uses AES-GCM and base64 to encrypt and authorize the cookie.
//
// Secret must be 128-125 Bits
func Set(w http.ResponseWriter, r *http.Request, user database.User, secret []byte) error {
	var buff bytes.Buffer
	err := gob.NewEncoder(&buff).Encode(&user)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     "AwesomeKey",
		Value:    buff.String(),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	block, err := aes.NewCipher(secret)
	if err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}

	plaintext := fmt.Sprintf("%s:%s", cookie.Name, cookie.Value)

	encryptedtext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	b64string := base64.URLEncoding.EncodeToString(encryptedtext)
	cookie.Value = string(b64string)

	http.SetCookie(w, cookie)
	return nil
}

// Get decrypts cookie and returns the User.
func Get(r *http.Request, secret []byte) (database.User, error) {
	var user database.User
	cookie, err := r.Cookie("AwesomeKey")
	if err != nil {
		return user, err
	}

	encryptedtext, err := base64.URLEncoding.Strict().DecodeString(cookie.Value)
	if err != nil {
		return user, err
	}

	block, err := aes.NewCipher(secret)
	if err != nil {
		return user, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return user, err
	}

	nonceSize := aesGCM.NonceSize()

	if len(encryptedtext) < nonceSize {
		err = errors.New("Cookie length does not match a expected nonce size")
		return user, err
	}

	nonce := encryptedtext[:nonceSize]
	ciphertext := encryptedtext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return user, err
	}

	_, value, ok := strings.Cut(string(plaintext), ":")
	if !ok {
		err = errors.New("Cookie is invalid")
		return user, err
	}

	reader := strings.NewReader(value)

	if err = gob.NewDecoder(reader).Decode(&user); err != nil {
		err = errors.New("Could not read cookie")
		return user, err
	}

	return user, nil
}

// Remove removes the cookie from client.
func Remove(w http.ResponseWriter, r *http.Request) {
}
