package cookies

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/chat_app/pkg/database"
)

// Set uses AES-GCM to to encrypt and authorize cookie.
func Set(w http.ResponseWriter, r *http.Request, user database.User, secret []byte) {
	var buff bytes.Buffer
	err := gob.NewEncoder(&buff).Encode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
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
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	plaintext := fmt.Sprintf("%s:%s", cookie.Name, cookie.Value)

	encryptedtext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)

	cookie.Value = string(encryptedtext)

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}

// Get decrypts cookie and returns the User.
func Get(w http.ResponseWriter, r *http.Request, secret []byte) (database.User, error) {
	var user database.User
	cookie, err := r.Cookie("AwesomeKey")
	if err != nil {
		log.Println(err)
		return user, err
	}

	encryptedtext := cookie.Value

	block, err := aes.NewCipher(secret)
	if err != nil {
		log.Println(err)
		return user, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		http.Error(w, "Invalid Cookie", http.StatusBadRequest)
	}

	_, value, ok := strings.Cut(string(plaintext), ":")
	if !ok {
		log.Println(err)
		http.Error(w, "Invalid Cookie", http.StatusBadRequest)
	}

	reader := strings.NewReader(value)

	if err := gob.NewDecoder(reader).Decode(&user); err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	return user, nil
}

// Remove removes the cookie from client.
func Remove(w http.ResponseWriter, r *http.Request) {
}
