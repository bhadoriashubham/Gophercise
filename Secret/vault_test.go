package secret

import (
	"path/filepath"
	"reflect"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestFile(t *testing.T) {
	type args struct {
		encodingKey string
		filepath    string
	}

	home, _ := homedir.Dir()
	path := filepath.Join(home, ".test_secrets")
	tests := []struct {
		name string
		args args
		want *Vault
	}{
		{name: "t1", args: args{encodingKey: "key", filepath: path},
			want: &Vault{encodingKey: "key", filepath: path},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := File(tt.args.encodingKey, tt.args.filepath); reflect.DeepEqual(got, tt.want) {
				t.Errorf("File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVault_loadKeyvalues(t *testing.T) {

	home, _ := homedir.Dir()
	path := filepath.Join(home, ".secrets")
	v := File("shubham", path)
	err := v.loadKeyvalues()
	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestVault_loadKeyvaluesNeg(t *testing.T) {

	home, _ := homedir.Dir()
	path := filepath.Join(home, ".test_secrets")
	v := File("xyz", path)
	err := v.loadKeyvalues()
	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestVault_Set(t *testing.T) {
	home, _ := homedir.Dir()
	path := filepath.Join(home, ".secrets")
	v := File("shubham", path)
	key := "twitter_api_key"
	err := v.Set(key, "some-value")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestVault_Get(t *testing.T) {

	home, _ := homedir.Dir()
	path := filepath.Join(home, ".secrets")
	v := File("shubham", path)

	key := "twitter_api_key"
	value, err := v.Get(key)
	if value != "some-value" {
		t.Errorf(err.Error())
	}
}

func TestVault_GetNeg(t *testing.T) {

	home, _ := homedir.Dir()
	path := filepath.Join(home, ".secrets")
	v := File("shubham", path)

	key := "twitter_api_ke"
	_, err := v.Get(key)
	if err == nil {
		t.Errorf(err.Error())
	}
}
