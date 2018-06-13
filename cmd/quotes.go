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
	"github.com/parsnips/sniper/streamer"
	"github.com/spf13/cobra"
)

var (
	streamerToken string
)

// quotesCmd represents the quotes command
var quotesCmd = &cobra.Command{
	Use:   "quotes",
	Short: "Open a websocket to receive quotes for a set of underlyings",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		streamer.TwQuotes(streamerToken, symbols)
	},
}

func init() {
	quotesCmd.Flags().StringVar(&streamerToken, "streamer-token", "The streamer token from get streamer tokens command", "--streamer-token <an-token>")
	quotesCmd.MarkFlagRequired("streamer-token")
	quotesCmd.Flags().StringArrayVar(&symbols, "symbol", symbols, "--symbol TSLA --symbol AAPL etc")
	quotesCmd.MarkFlagRequired("symbol")
	rootCmd.AddCommand(quotesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quotesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quotesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
