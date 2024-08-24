package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FindNickName(nickname string) (string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", NameUrl+nickname, nil)
	if err != nil {
		return "", err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + APIKey},
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Player not found. Status code: %d", res.StatusCode)
	}

	var player map[string]interface{}
	if err := json.Unmarshal(body, &player); err != nil {
		return "", fmt.Errorf("Error unmarshalling player data: %s", err)
	}

	// Check if the "player_id" field exists
	if playerID, ok := player["player_id"].(string); ok {
		return playerID, nil
	}

	return "", fmt.Errorf("Player not found.")
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
