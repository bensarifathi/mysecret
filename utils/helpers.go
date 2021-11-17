package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"gopkg.in/yaml.v2"
)

const SECRET_FILE = ".mysecret.yml"

func GetOrCreateFile(filename string) (*os.File, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	fullpath := path.Join(home, filename)
	return os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
}

func JsonToYaml(rawdata []byte) ([]byte, error) {
	user := Author{}
	err := json.Unmarshal(rawdata, &user)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(&user)
}

func GetUserFromFile() (*Author, error) {
	home, _ := os.UserHomeDir()
	userinfo := path.Join(home, SECRET_FILE)
	if _, err := os.Stat(userinfo); err != nil {
		return nil, errors.New("[!] no user logged in found")
	}
	data, err := ioutil.ReadFile(userinfo)
	if err != nil {
		msgErr := fmt.Sprintf("[!] internal error (%s)\n", err)
		return nil, errors.New(msgErr)
	}
	user := &Author{}
	if err := yaml.Unmarshal(data, user); err != nil {
		msgErr := fmt.Sprintf("[!] internal error (%s)\n", err)
		return nil, errors.New(msgErr)
	}
	return user, nil
}

func NewHttpClient(timeout int) *http.Client {
	return &http.Client{Timeout: time.Duration(timeout) * time.Second}
}
