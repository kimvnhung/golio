package lol

import "github.com/kimvnhung/golio/internal"

// Client pools all methods for endpoints of the League of Legends API.
type Client struct {
	Challenge       *ChallengesClient
	ChampionMastery *ChampionMasteryClient
	Champion        *ChampionClient
	League          *LeagueClient
	Status          *StatusClient
	Match           *MatchClient
	Spectator       *SpectatorClient
	Summoner        *SummonerClient
	ThirdPartyCode  *ThirdPartyCodeClient
	Tournament      *TournamentClient
}

// NewClient returns a new instance of a League of Legends client.
func NewClient(base *internal.Client) *Client {
	return &Client{
		Challenge:       &ChallengesClient{c: base},
		ChampionMastery: &ChampionMasteryClient{c: base},
		Summoner:        &SummonerClient{c: base},
		Champion:        &ChampionClient{c: base},
		League:          &LeagueClient{c: base},
		Status:          &StatusClient{c: base},
		Match:           &MatchClient{c: base},
		Spectator:       &SpectatorClient{c: base},
		Tournament:      &TournamentClient{c: base},
		ThirdPartyCode:  &ThirdPartyCodeClient{c: base},
	}
}
