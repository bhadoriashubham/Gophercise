package encrypt

import (
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestEncrypt(t *testing.T) {
	home, _ := homedir.Dir()
	path := filepath.Join(home, ".test_secret")
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	cipherstream, err := Encrypt("test-key", f)
	if err != nil {
		t.Errorf("Error for %s: %d.", cipherstream, err)
	}
	f.Close()
}
func TestEncryptNeg(t *testing.T) {
	testfile, err := os.OpenFile("", os.O_RDWR|os.O_CREATE, 0755)
	cipherstream, err := Encrypt("test-key", testfile)
	if err == nil {
		t.Errorf("Error for %s: %d.", cipherstream, err)
	}
	testfile.Close()
}

func TestDecrypt(t *testing.T) {

	home, _ := homedir.Dir()
	fp := filepath.Join(home, ".test_secret")
	f, _ := os.Open(fp)
	defer f.Close()
	_, err := Decrypt("test-key", f)
	if err != nil {
		t.Errorf("Expected NO error but got following error : %v ", err)
	}
}
func TestDecryptNeg(t *testing.T) {

	home, _ := homedir.Dir()
	fp := filepath.Join(home, ".test_secre")
	f, _ := os.Open(fp)
	defer f.Close()
	_, err := Decrypt("test-key", f)
	if err == nil {
		t.Errorf("Expected NO error but got following error : %v ", err)
	}
}
