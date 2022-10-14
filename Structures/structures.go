package structures

import "encoding/json"

type LoginResponse struct {
	StatusCode        int    `json:"status_code"`
	StatusDescription string `json:"status_description"`
	Data              string `json:"data"`
}

type LoginRequest struct {
	Msisdn string `json:"msisdn"`
	Pwd    string `json:"pwd"`
}

type LeaderBoardResponse struct {
	StatusCode        int    `json:"status_code"`
	StatusDescription string `json:"status_description"`
	Data              struct {
		Profiles []struct {
			Msisdn    string  `json:"msisdn"`
			FirstName string  `json:"first_name"`
			LastName  string  `json:"last_name"`
			BetsCount float64 `json:"bets_count"`
			Pos       int     `json:"pos"`
		} `json:"profiles"`
		Meta struct {
			StartsOn string `json:"starts_on"`
			EndsOn   string `json:"ends_on"`
			Gold     int    `json:"gold"`
			Bronze   int    `json:"bronze"`
			Silver   int    `json:"silver"`
			Total    int    `json:"total"`
		} `json:"meta"`
	} `json:"data"`
}

type AccountData struct {
	Token string `json:"token"`
	User  struct {
		ProfileID          string      `json:"profile_id"`
		Msisdn             string      `json:"msisdn"`
		FirstName          string      `json:"first_name"`
		LastName           string      `json:"last_name"`
		Created            string      `json:"created"`
		Balance            string      `json:"balance"`
		BonusBalance       string      `json:"bonus_balance"`
		PointsBalance      json.Number `json:"points_balance"`
		CasinoBalance      string      `json:"casino_balance"`
		CasinoBonusBalance string      `json:"casino_bonus_balance"`
		Freebet            int         `json:"freebet"`
		OpenBets           int         `json:"open_bets"`
		Theme              string      `json:"theme"`
		ReferralCode       interface{} `json:"referral_code"`
		ProfileIDStp       interface{} `json:"profile_id_stp"`
		TokenStp           interface{} `json:"token_stp"`
		Status             int         `json:"status"`
		Sessions           int         `json:"sessions"`
	} `json:"user"`
}
