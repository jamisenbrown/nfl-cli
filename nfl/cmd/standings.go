/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// standingsCmd represents the standings command
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "NFL standings",
	Long:  `Show the current NFL standings`,
	Run:   getStandings,
}

func getStandings(cmd *cobra.Command, args []string) {
	fmt.Println("Getting NFL standings")

	apiUrl := "https://cdn.espn.com/core/nfl/standings?xhr=1"
	request, error := http.NewRequest("GET", apiUrl, nil)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(request)
}

func init() {
	rootCmd.AddCommand(standingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// standingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// standingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
