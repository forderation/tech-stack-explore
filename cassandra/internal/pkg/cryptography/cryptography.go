package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type SecretSize int

const (
	SecretAES128 SecretSize = 8
	SecretAES192 SecretSize = 12
	SecretAES256 SecretSize = 32
)

type Cryptography struct {
	secret string
}

func New(secret string) Cryptography {
	return Cryptography{
		secret: secret,
	}
}

func (c Cryptography) Secret(size SecretSize) ([]byte, error) {
	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

func (c Cryptography) encrypt(data, secret []byte, isString bool) (string, []byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", nil, fmt.Errorf("new cipher: %w", err)
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil, fmt.Errorf("new gcm: %w", err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", nil, fmt.Errorf("io read: %w", err)
	}

	bytes := aead.Seal(nonce, nonce, data, nil)
	if isString {
		return base64.URLEncoding.EncodeToString(bytes), nil, nil
	}
	return "", bytes, nil
}

func (c Cryptography) EncryptAsString(data, secret []byte) (string, error) {
	if secret == nil {
		secret = []byte(c.secret)
	}
	val, _, err := c.encrypt(data, secret, true)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c Cryptography) EncryptAsByte(data, secret []byte) ([]byte, error) {
	if secret == nil {
		secret = []byte(c.secret)
	}
	_, val, err := c.encrypt(data, secret, false)
	if err != nil {
		return nil, err
	}
	return val, err
}

func (c Cryptography) DecryptByte(data, secret []byte) ([]byte, error) {
	if secret == nil {
		secret = []byte(c.secret)
	}

	return c.decrypt(data, secret)
}

func (c Cryptography) DecryptString(data string, secret []byte) ([]byte, error) {
	if secret == nil {
		secret = []byte(c.secret)
	}

	byte, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("decode string: %w", err)
	}

	return c.decrypt(byte, secret)
}

func (c Cryptography) decrypt(data, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, fmt.Errorf("new cipher : %w", err)
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("new gcm: %w", err)
	}

	size := aead.NonceSize()
	if len(data) < size {
		return nil, fmt.Errorf("nonce size: invalid length")
	}

	nonce, text := data[:size], data[size:]
	res, err := aead.Open(nil, nonce, text, nil)
	if err != nil {
		return nil, fmt.Errorf("aead open: %w", err)
	}
	return res, nil
}
