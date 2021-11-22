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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/Fathi-BENSARI/mysecret/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "get user information",
	Long: `get user login information and return them to stdout in case finding
user credential`,
	Run: func(cmd *cobra.Command, args []string) {
		getUserInfo()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getUserInfo() {
	home, _ := os.UserHomeDir()
	fullpath := path.Join(home, utils.SECRET_FILE)
	if _, err := os.Stat(fullpath); err != nil {
		fmt.Println("[!) No user login found")
		return
	}
	rawData, err := ioutil.ReadFile(fullpath)
	if err != nil {
		log.Fatal("[!] Internal error", err)
	}
	u := utils.Author{}
	if err := yaml.Unmarshal(rawData, &u); err != nil {
		log.Fatal("[!] Internal error", err)
	}
	fmt.Printf("[+] Current User Info:\nUsername: %-10sEmail: %s\n", u.Name, u.Email)
}
