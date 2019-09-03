package cipher

import (
	"crypto/cipher"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

var encodingKey = "test-secret"

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, "secrets-test")
}

func TestEncryptWriter(t *testing.T) {
	f, _ := os.OpenFile(secretsPath(), os.O_RDWR|os.O_CREATE, 0755)
	_, err := EncryptWriter(encodingKey, f)
	if err != nil {
		t.Error("Got an error: ", err)
	}

}

func TestEncryptWriterNegative(t *testing.T) {
	f, _ := os.OpenFile(secretsPath(), os.O_RDWR|os.O_CREATE, 0755)

	temp := NewCipherBlockFunc
	NewCipherBlockFunc = func(key string) (cipher.Block, error) {
		return nil, errors.New("ERROR TEST")
	}
	_, err := EncryptWriter(encodingKey, f)
	if err == nil {
		t.Error("Expecting an error but got ", err)
	}
	NewCipherBlockFunc = temp

	VarTemp := IoReadFulFunc
	defer func() {
		IoReadFulFunc = VarTemp
	}()
	IoReadFulFunc = func(r io.Reader, buf []byte) (n int, err error) {
		return 0, errors.New("ERROR TEST.")
	}

	defer f.Close()
	_, err1 := EncryptWriter(encodingKey, f)
	if err1 == nil {
		t.Error("Expecting an error but got ", err1)
	}

}

func TestDecryptReader(t *testing.T) {
	f, _ := os.OpenFile(secretsPath(), os.O_RDWR|os.O_CREATE, 0755)
	_, err := DecryptReader(encodingKey, f)
	if err != nil {
		t.Error("Got an error: ", err)
	}
}

func TestDecryptReaderNegative(t *testing.T) {

	temp := NewCipherBlockFunc
	NewCipherBlockFunc = func(key string) (cipher.Block, error) {
		return nil, errors.New("ERROR TEST")
	}
	f, _ := os.OpenFile(secretsPath(), os.O_RDWR|os.O_CREATE, 0755)
	_, err := DecryptReader(encodingKey, f)
	if err == nil {
		t.Error("Expecting an error but got", err)
	}
	NewCipherBlockFunc = temp

	f.Close()
	_, err1 := DecryptReader(encodingKey, f)
	if err1 == nil {
		t.Error("Expecting an error but got", err1)
	}
}

func TestEncryptStream(t *testing.T) {
	temp := NewCipherBlockFunc
	NewCipherBlockFunc = func(key string) (cipher.Block, error) {
		return nil, errors.New("ERROR TEST")
	}
	encryptStream("test-key", []byte{10})
	NewCipherBlockFunc = temp
}

func TestDecryptStream(t *testing.T) {
	temp := NewCipherBlockFunc
	NewCipherBlockFunc = func(key string) (cipher.Block, error) {
		return nil, errors.New("ERROR TEST")
	}
	decryptStream("test-key", []byte{10})
	NewCipherBlockFunc = temp
}
