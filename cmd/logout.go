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
	"os"
	"path"

	"github.com/Fathi-BENSARI/mysecret/utils"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout from mysecret",
	Long: `logout from the mysecret application and remove 
all stored user information`,
	Run: func(cmd *cobra.Command, args []string) {
		logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func logout() {
	// get the home dir and build the fullpath name
	home, _ := os.UserHomeDir()
	fullpath := path.Join(home, utils.SECRET_FILE)
	// check if the there is a logged in user
	if _, err := os.Stat(fullpath); err != nil {
		fmt.Printf("[!] no user logged in found\n")
		return
	}
	os.Remove(fullpath)
	fmt.Printf("[+] logout successfully\n")
}
