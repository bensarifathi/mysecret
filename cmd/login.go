/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Fathi-BENSARI/mysecret/utils"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to your mysecret account",
	Long:  `login to your mysecret account`,
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func login() {
	var email string
	var backend = "https://my-secret-store.herokuapp.com"
	var url string = backend + "/user/login"
	// gather information
	fmt.Printf("Email: ")
	fmt.Scanln(&email)
	fmt.Printf("Passowrd: ")
	password, _ := gopass.GetPasswdMasked()
	// post request to the server
	client := utils.NewHttpClient(10)
	form := utils.Author{
		Email:    email,
		Password: string(password),
	}
	data, _ := json.Marshal(form)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("login error (%s)\n", body)
		return
	}
	data, err = utils.JsonToYaml(body)
	if err != nil {
		fmt.Printf("[!] internal error %s\n", err)
		return
	}
	fd, err := utils.GetOrCreateFile(utils.SECRET_FILE)
	if err != nil {
		fmt.Printf("[!] internal error please try again\n")
		return
	}
	defer fd.Close()
	_, err = fd.Write(data)
	if err != nil {
		fmt.Printf("[!] error while saving file %s (%s)", utils.SECRET_FILE, err)
		return
	}
	fmt.Printf("[+] user login successfully\n")
}
