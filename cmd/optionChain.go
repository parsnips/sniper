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

var symbol string

// optionChainCmd represents the optionChain command
var optionChainCmd = &cobra.Command{
	Use:   "optionChain",
	Short: "Get the option chain for an underlying",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := search.OptionChain(symbol, sessionToken)
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
	optionChainCmd.Flags().StringVar(&sessionToken, "session-token", "session token from login", "A valid session token from logging in.")
	optionChainCmd.MarkFlagRequired("session-token")

	optionChainCmd.Flags().StringVar(&symbol, "symbol", "--symbol TSLA", "symbol for underlying")
	optionChainCmd.MarkFlagRequired("symbol")
	rootCmd.AddCommand(optionChainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optionChainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optionChainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
