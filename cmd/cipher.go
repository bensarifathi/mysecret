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
	"log"
	"net/http"
	"os"

	"github.com/Fathi-BENSARI/mysecret/utils"
	"github.com/spf13/cobra"
)

// cipherCmd represents the cipher command
var cipherCmd = &cobra.Command{
	Use:   "cipher",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		addKey()
	},
}

func init() {
	rootCmd.AddCommand(cipherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cipherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cipherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addKey() {
	// grab user input
	var key string
	var backend = os.Getenv("BACKEND")
	var url string = backend + "/passphrase"
	// gather information
	fmt.Printf("Insert secret key: ")
	fmt.Scanln(&key)
	// grab the user info or return an error
	user, err := utils.GetUserFromFile()
	if err != nil {
		log.Fatal(err)
	}
	// create a client
	c := utils.NewHttpClient(10)
	form := &utils.Cipher{
		Key:   key,
		Owner: user.ID,
	}
	data, _ := json.Marshal(form)
	resp, err := c.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("error while saving key (%s)\n", body)
	}
	fmt.Println("[+] cipher key successfully created")
}
