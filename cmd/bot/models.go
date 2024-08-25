package main

type MatchData struct {
	Items []Match `json:"items"`
}

type Match struct {
	Stats MatchStats `json:"stats"`
}

type MatchStats struct {
	Kills               string `json:"Kills"`
	Deaths              string `json:"Deaths"`
	Assists             string `json:"Assists"`
	HeadshotsPercentage string `json:"Headshots %"`
	Adr                 string `json:"ADR"`
	KdRatio             string `json:"K/D Ratio"`
	KrRatio             string `json:"K/R Ratio"`
}

type Profile struct {
	PlayerID  string `json:"player_id"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Country   string `json:"country"`
	FaceitURL string `json:"faceit_url"`
}
