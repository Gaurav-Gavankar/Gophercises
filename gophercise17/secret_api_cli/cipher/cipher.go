package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

var IoReadFulFunc = io.ReadFull
var NewCipherBlockFunc = newCipherBlock

func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := NewCipherBlockFunc(key)
	if err != nil {
		return nil, err
	}

	return cipher.NewCFBEncrypter(block, iv), nil
}

//Encrypt takes in a Key and Plaintext and returns encrypted value of it in Hex format.
/*func Encrypt(key, plaintext string) (string, error) {

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode, err := encryptStream(key, iv)
	if err != nil {
		return "", err
	}

	mode.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), nil
}
*/

//Encrypt writer will return a writer that will write encrypted data to original writer.
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := IoReadFulFunc(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream, err := encryptStream(key, iv)
	if err != nil {
		return nil, err
	}

	_, err = w.Write(iv)
	if err != nil {
		return nil, errors.New("Encrypt: enable to write full iv to writer.")
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := NewCipherBlockFunc(key)
	if err != nil {
		return nil, err
	}

	return cipher.NewCFBDecrypter(block, iv), nil
}

//Decrypt takes a key and ciphertext in hex format and decrypt it.
/*func Decrypt(key, cipherHex string) (string, error) {

	ciphertext, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", nil
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: cipher too short.")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	fmt.Println("ct", ciphertext, "iv", iv)
	stream, err := decryptStream(key, iv)
	if err != nil {
		return "", err
	}

	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}*/

//DecryptReader will return a reader that will decrypt data from the provided reader
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("Encrypt: Enable to read full iv.")
	}
	stream, err := decryptStream(key, iv)
	if err != nil {
		return nil, err
	}

	return &cipher.StreamReader{S: stream, R: r}, nil

}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprintf(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)

}
