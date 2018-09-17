package secret

import (
	"encoding/json"
	"errors"
	"github.com/bhadoriashubham/gophercises/secret/encrypt"
	"io"
	"os"
	"sync"
)

//File Returns encoding key
func File(encodingKey, filepath string) *Vault {

	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
		KeyValues:   make(map[string]string)}
}

//Vault vault struct
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	KeyValues   map[string]string
}

func (v *Vault) loadKeyvalues() error {

	f, err := os.Open(v.filepath)
	if err != nil {
		v.KeyValues = make(map[string]string)
		return nil
	}

	defer f.Close()
	r, err := encrypt.Decrypt(v.encodingKey, f)
	return v.readKeyValues(r)

}

func (v *Vault) readKeyValues(r io.Reader) error {

	dec := json.NewDecoder(r)
	err := dec.Decode(&v.KeyValues)
	return err

}

//Get gets the value associated with the key
func (v *Vault) Get(Key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.loadKeyvalues()
	var value string
	var ok bool
	if err == nil {

		value, ok = v.KeyValues[Key]

		if !ok {
			return "", errors.New("No value found")
		}

	}

	return value, nil
}

//Set Sets the key and value
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.loadKeyvalues()

	if err == nil {
		v.KeyValues[key] = value
		err = v.saveKeyvalues()

	}

	return err

}
func (v *Vault) saveKeyvalues() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	var w io.Writer
	if err == nil {
		defer f.Close()
		w, err = encrypt.Encrypt(v.encodingKey, f)
	}
	return v.writeKeyValues(w)
}

func (v *Vault) writeKeyValues(w io.Writer) error {

	enc := json.NewEncoder(w)
	return enc.Encode(v.KeyValues)
}
