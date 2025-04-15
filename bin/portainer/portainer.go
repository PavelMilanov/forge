// Package portainer реализует взаимодействие с Portainer API.
package portainer

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PavelMilanov/forge/config"
)

type PortainerClient struct {
	Url      string
	Username string
	Password string
	Token    string
}

func NewPortainerClient(config *config.Env) *PortainerClient {
	return &PortainerClient{
		Url:      config.Portainer.Url,
		Username: config.Portainer.Login,
		Password: config.Portainer.Password,
	}
}

func (pc *PortainerClient) getToken() error {
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	type authRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	authData := authRequest{
		Username: pc.Username,
		Password: pc.Password,
	}
	jsonData, err := json.Marshal(authData)
	if err != nil {
		return err
	}
	client := &http.Client{Transport: customTransport}
	req, err := http.NewRequest("POST", pc.Url+"/api/auth", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	type authResponse struct {
		Jwt string `json:"jwt"`
	}
	var authResp authResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return err
	}
	pc.Token = authResp.Jwt
	return nil
}

func (pc *PortainerClient) GetEnvironments() error {
	if err := pc.getToken(); err != nil {
		return err
	}
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: customTransport}
	req, err := http.NewRequest("GET", pc.Url+"/api/endpoints", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", pc.Token))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	type endpointResponse struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	var endpoints []endpointResponse
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", endpoints)
	return nil
}
