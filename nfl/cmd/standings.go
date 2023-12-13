/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var playoffs bool

type teamInfo struct {
	DisplayName string `json:"displayName"`
	Seed        string `json:"seed"`
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

type conference struct {
	Name      string     `json:"name"`
	Divisions []division `json:"groups"`
}

type contentStandings struct {
	Name        string       `json:"name"`
	Conferences []conference `json:"groups"`
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
	flag.Parse()
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

	if playoffs {
		printPlayoffStandings(standingsResponse)
	} else {
		printStandings(standingsResponse)
	}
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

func printPlayoffStandings(standings standingsResponse) {
	fmt.Println("AFC Conference")
	printConferencePlayoffs(standings.Content.Standings.Conferences[0])
	fmt.Println()

	fmt.Println("NFC Conference")
	printConferencePlayoffs(standings.Content.Standings.Conferences[1])
}

func printConferencePlayoffs(conference conference) {
	var conferenceTeams [16]string
	for _, division := range conference.Divisions {
		for _, team := range division.Standings.StandingEntries {
			seed, err := strconv.Atoi(team.TeamInfo.Seed)
			if err != nil {
				fmt.Println("Abort!!")
			}
			conferenceTeams[seed-1] = team.TeamInfo.DisplayName
		}
	}

	for seed, team := range conferenceTeams {
		if seed == 7 {
			break
		}
		fmt.Println(strconv.Itoa(seed+1) + ". " + team)
	}
}

func init() {
	rootCmd.AddCommand(standingsCmd)

	standingsCmd.Flags().BoolVarP(&playoffs, "playoffs", "p", false, "Shows playoff standings")
}
