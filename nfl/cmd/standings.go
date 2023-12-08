/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type teamInfo struct {
	DisplayName string `json:"displayName"`
	Id          string `json:"id"`
}

type teamStats struct {
	DisplayName  string `json:"displayName"`
	DisplayValue string `json:"displayValue"`
}

type divisionStandings struct {
	Name            string `json:"name"`
	StandingEntries []struct {
		TeamStats []teamStats `json:"stats"`
		TeamInfo  teamInfo    `json:"team"`
	} `json:"entries"`
}

type division struct {
	Name      string            `json:"name"`
	Standings divisionStandings `json:"standings"`
}

type conference []struct {
	Name      string     `json:"name"`
	Divisions []division `json:"groups"`
}

type contentStandings struct {
	Name        string     `json:"name"`
	Conferences conference `json:"groups"`
}

type standingsResponse struct {
	Content struct {
		Title       string           `json:"title"`
		Description string           `json:"description"`
		Sport       string           `json:"sport"`
		Standings   contentStandings `json:"standings"`
	} `json:"content"`
}

// standingsCmd represents the standings command
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "NFL standings",
	Long:  `Show the current NFL standings`,
	Run:   getStandings,
}

func getStandings(cmd *cobra.Command, args []string) {
	fmt.Println("Getting NFL standings")
	fmt.Println()

	apiUrl := "https://cdn.espn.com/core/nfl/standings?xhr=1"

	espnClient := http.Client{
		Timeout: time.Second * 10,
	}

	request, error := http.NewRequest(http.MethodGet, apiUrl, nil)

	if error != nil {
		fmt.Println(error)
	}

	response, getErr := espnClient.Do(request)
	if getErr != nil {
		fmt.Println(getErr)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	body, readErr := io.ReadAll(response.Body)

	if readErr != nil {
		fmt.Println(readErr)
	}

	standingsResponse := standingsResponse{}
	jsonErr := json.Unmarshal(body, &standingsResponse)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	printStandings(standingsResponse)
}

func printStandings(standings standingsResponse) {
	// AFC
	printDivision(standings.Content.Standings.Conferences[0].Divisions[0])
	printDivision(standings.Content.Standings.Conferences[0].Divisions[1])
	printDivision(standings.Content.Standings.Conferences[0].Divisions[2])
	printDivision(standings.Content.Standings.Conferences[0].Divisions[3])

	// NFC
	printDivision(standings.Content.Standings.Conferences[1].Divisions[0])
	printDivision(standings.Content.Standings.Conferences[1].Divisions[1])
	printDivision(standings.Content.Standings.Conferences[1].Divisions[2])
	printDivision(standings.Content.Standings.Conferences[1].Divisions[3])
}

func printDivision(division division) {
	fmt.Println(division.Name)
	fmt.Println("-----------")
	for _, team := range division.Standings.StandingEntries {
		fmt.Println("| " + team.TeamInfo.DisplayName + " " + team.TeamStats[0].DisplayValue + "-" + team.TeamStats[1].DisplayValue + "-" + team.TeamStats[2].DisplayValue)
	}
	fmt.Println()
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
