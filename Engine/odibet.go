package engine

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"odibet/Structures"
)

func GetLeaderboard() (*structures.LeaderBoardResponse, error) {
	var lbr structures.LeaderBoardResponse
	response, err := http.DefaultClient.Get("https://odibets.com/api/leaderboard")
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBody, &lbr)
	if err != nil {
		return nil, err
	}
	return &lbr, nil
}


func Login(phone, password string) (*structures.AccountData, error) {
	var loginResp structures.LoginResponse
	loginReq := structures.LoginRequest{
		Msisdn: phone,
		Pwd:    password,
	}
	odiClient := http.Client{}
	payloadBytes, err := json.Marshal(loginReq)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	odiRequest, err := http.NewRequest(http.MethodPost, "https://odibets.com/api/va", body)
	if err != nil {
		return nil, err
	}
	odiRequest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	response, err := odiClient.Do(odiRequest)
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	err = json.Unmarshal(responseBody, &loginResp)
	if err != nil {
		return nil, err
	}

	

	var accountData structures.AccountData
	accdata, err := base64.StdEncoding.DecodeString(loginResp.Data[2:])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(accdata, &accountData)
	if err != nil {
		return nil, err
	}
	return &accountData, nil
}
