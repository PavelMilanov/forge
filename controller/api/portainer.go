package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/PavelMilanov/forge/config"
)

/*
PortainerAPI предоставляет интерфейс для взаимодействия с Portainer.
*/
type Portainer struct {
	Url string
	Key string
}

type Stack struct {
	ID       int    `json:"Id"`
	Name     string `json:"Name"`
	Endpoint int    `json:"EndpointId"`
	File     string `json:"StackFileContent"`
}

// /api/stacks
// /api/stacks/11/file
// /api/stacks/11?endpointId=8

func NewPortainer(env *config.Env) (*Portainer, error) {
	return &Portainer{Url: env.Portainer.Url, Key: env.Portainer.Key}, nil
}

func (p *Portainer) GetStacks() ([]Stack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.Url+"/api/stacks", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	var stacks []Stack

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return nil, err
	}
	err = json.Unmarshal(body, &stacks)
	if err != nil {
		return nil, err
	}
	return stacks, nil
}

func (p *Portainer) GetStackFile(stack Stack) (*Stack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.Url+"/api/stacks/"+strconv.Itoa(stack.ID)+"/file", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return nil, err
	}
	err = json.Unmarshal(body, &stack)
	if err != nil {
		return nil, err
	}
	return &stack, nil
}

func (p *Portainer) UpdateStack(stack *Stack) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	params := url.Values{}
	params.Add("endpointId", strconv.Itoa(stack.Endpoint))

	data := struct {
		Env       []string `json:"env"`
		Prune     bool     `json:"prune"`
		PullImage bool     `json:"pullImage"`
		File      string   `json:"stackFileContent"`
	}{
		Env:       []string{},
		Prune:     true,
		PullImage: true,
		File:      stack.File,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	requestBody := bytes.NewReader(jsonData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, p.Url+"/api/stacks/"+strconv.Itoa(stack.ID), requestBody)
	if err != nil {
		return "", err
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return "", err
	}
	return string(body), nil
}
