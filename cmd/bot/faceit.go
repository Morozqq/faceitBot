package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FindNickName(nickname string) (string, Profile, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", NameUrl+nickname, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return "", Profile{}, err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + APIKey},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing name request")
		return "", Profile{}, err
	}

	defer res.Body.Close()

	body, ioError := io.ReadAll(res.Body)
	if ioError != nil {
		fmt.Println("Error reading response body")
		return "", Profile{}, ioError
	}

	if res.StatusCode != 200 {
		return "", Profile{}, fmt.Errorf("Player not found. Status code: %d", res.StatusCode)
	}

	var profile Profile
	jsonError := json.Unmarshal(body, &profile)
	if jsonError != nil {
		return "", Profile{}, fmt.Errorf("Error unmarshalling player data: %s", jsonError)
	}
	profile.FaceitURL = strings.Replace(profile.FaceitURL, "{lang}", "en", 1)

	// Check if the "player_id" field exists
	if profile.PlayerID != "" {
		return profile.PlayerID, profile, nil
	}

	return "", Profile{}, fmt.Errorf("Player not found.")
}

func FetchPlayerStats(playerID string) (MatchData, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", BaseURL+playerID+StatsEndPoint, nil)
	if err != nil {
		return MatchData{}, err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + APIKey},
	}

	res, err := client.Do(req)
	if err != nil {
		return MatchData{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return MatchData{}, fmt.Errorf("Error executing request: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return MatchData{}, err
	}

	var matchData MatchData
	if err := json.Unmarshal(body, &matchData); err != nil {
		return MatchData{}, err
	}

	return matchData, nil
}
