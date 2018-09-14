package cobra

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSetCmdPos(t *testing.T) {
	file, _ := os.OpenFile("testing1.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"twitter_api_key1", "some-value"}
	setCmd.Run(setCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Value set successfully")
	assert.Equal(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
}

func TestGetCmdPos(t *testing.T) {
	file, _ := os.OpenFile("testing1.txt", os.O_CREATE|os.O_RDWR, 0666)
	//defer os.Remove(file.Name())
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"twitter_api_key1"}
	getCmd.Run(getCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "some-value")
	assert.Equal(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
}
func TestGetCmdNeg(t *testing.T) {
	a := []string{"key value pair with whitespaces"}
	getCmd.Run(getCmd, a)

}
