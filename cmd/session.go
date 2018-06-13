// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"

	"github.com/parsnips/sniper/session"
	"github.com/spf13/cobra"
)

var (
	login    string
	password string
)

// sessionCmd represents the session command
var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Retrieve a session token from tastyworks",
	Long:  `Exchange your username/password for a token. Then call other apis with that token.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := session.Login(login, password)
		if err != nil {
			log.Fatalf("Couldnt retrieve session due to %v", err)
		}

		fmt.Printf("%s", string(resp.Body()))
		if resp.StatusCode() != 201 {
			log.Fatalf("Got response status code %s\n", resp.StatusCode())
		}

	},
}

func init() {
	sessionCmd.Flags().StringVar(&login, "login", "Tastyworks username", "The username for your tastywork account")
	sessionCmd.MarkFlagRequired("login")
	sessionCmd.Flags().StringVar(&password, "password", "Tastyworks password", "The password for your tastyworks account")
	sessionCmd.MarkFlagRequired("password")

	rootCmd.AddCommand(sessionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sessionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sessionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
