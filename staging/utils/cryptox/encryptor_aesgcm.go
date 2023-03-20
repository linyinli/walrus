package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// AesGcm returns an Encryptor with AES-GCM encryption.
func AesGcm(k []byte) (Encryptor, error) {
	var b, err = aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	return aesGcmEncryptor{b: b}, nil
}

// aesGcmEncryptor leverages AES-GCM encryption to implement Encryptor.
type aesGcmEncryptor struct {
	b cipher.Block
}

func (e aesGcmEncryptor) Encrypt(p []byte, a []byte) ([]byte, error) {
	var g, err = cipher.NewGCM(e.b)
	if err != nil {
		return nil, err
	}
	var n = make([]byte, g.NonceSize())
	_, err = io.ReadFull(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	return g.Seal(n, n, p, a), nil
}

func (e aesGcmEncryptor) Decrypt(c []byte, a []byte) ([]byte, error) {
	var g, err = cipher.NewGCM(e.b)
	if err != nil {
		return nil, err
	}
	var n = c[:g.NonceSize()]
	c = c[g.NonceSize():]
	return g.Open(nil, n, c, a)
}