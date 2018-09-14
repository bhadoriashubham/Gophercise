package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	return cipher.NewCFBEncrypter(block, iv), err
}

// EncryptWriter will return a writer that will write encrypted data to
// the original writer.
func Encrypt(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	_, err := io.ReadFull(rand.Reader, iv)
	stream, err := encryptStream(key, iv)
	n, err := w.Write(iv)
	err = checkIV(n, iv, err)
	return &cipher.StreamWriter{S: stream, W: w}, err
}

func checkIV(n int, iv []byte, err error) error {
	if len(iv) != n || err != nil {
		return errors.New("Unable to write IV into writer")
	}
	return nil
}

func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	return cipher.NewCFBDecrypter(block, iv), err
}

// DecryptReader will return a reader that will decrypt data from the
// provided reader and give the user a way to read that data as it if was
// not encrypted.
func Decrypt(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("encrypt: unable to read the full iv")
	}
	stream, err := decryptStream(key, iv)
	return &cipher.StreamReader{S: stream, R: r}, err
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
