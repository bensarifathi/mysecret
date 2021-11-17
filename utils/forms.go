package utils

import (
	"fmt"
	"strings"
)

type Author struct {
	ID       int    `json:"id,omitempty" yaml:"id,omitempty"`
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	Email    string `json:"email,omitempty" yaml:"email,omitempty"`
}

type Account struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Owner    int    `json:"owner"`
}

type Cipher struct {
	Key   string `json:"key"`
	Owner int    `json:"owner"`
}

func (a Account) String() string {
	header := fmt.Sprintf("%-10s%-40s%-10s\n", "Name", "Email", "Password")
	delimiter := strings.Repeat("=", len(header)) + "\n"
	content := fmt.Sprintf("%-10s%-40s%-10s\n", a.Name, a.Email, a.Password)
	return delimiter + header + content
}
