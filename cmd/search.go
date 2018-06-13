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

	"github.com/parsnips/sniper/search"
	"github.com/spf13/cobra"
)

var (
	symbols []string
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for an underlying by symbol",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := search.Search(symbols, sessionToken)
		if err != nil {
			log.Fatalf("Couldnt retrieve streamer token due to %v", err)
		}
		fmt.Printf("%s", string(resp.Body()))
		if resp.StatusCode() != 200 {
			log.Fatalf("Got response status code %s\n", resp.StatusCode())
		}
	},
}

func init() {
	searchCmd.Flags().StringVar(&sessionToken, "session-token", "session token from login", "A valid session token from logging in.")
	searchCmd.MarkFlagRequired("session-token")

	searchCmd.Flags().StringArrayVar(&symbols, "symbol", symbols, "--symbol TSLA --symbol AAPL etc")
	searchCmd.MarkFlagRequired("symbol")
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
