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
	Nickname            string `json:"nickname"`
}
