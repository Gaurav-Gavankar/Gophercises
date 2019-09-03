package secret_api_cli

/*
import (
	"errors"
	"gophercises/gophercise17/secret_api_cli/encrypt"
)

func Memory(encodingKey string) Vault {
	return Vault{
		encodingKey: encodingKey,
		keyValues:   make(map[string]string),
	}
}

type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

func (v *Vault) Get(key string) (string, error) {
	hex, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("No value found for the Key.")
	}
	ret, err := encrypt.Decrypt(v.encodingKey, hex)
	if err != nil {
		return "", err
	}
	return ret, nil
}

func (v *Vault) Set(key, value string) error {
	encryptedValue, err := encrypt.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}
	v.keyValues[key] = encryptedValue
	return nil
}
*/
