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
)

/*
PortainerAPI предоставляет интерфейс для взаимодействия с Portainer.
*/
type Portainer struct {
	Url string
	Key string
}

/*
Stack представляет абстракцию при взаимодействии с json в Portainer API.
*/
type Stack struct {
	ID           int    `json:"Id"`
	StackName    string `json:"Name"`
	TemplateName string `json:"Title"`
	Endpoint     int    `json:"EndpointId"`
	StackFile    string `json:"StackFileContent"`
	TemplateFile string `json:"FileContent"`
}

// /api/stacks
// /api/stacks/11/file
// /api/stacks/11?endpointId=8

// NewPortainer создает новый экземпляр Portainer.
func NewPortainer(url, key string) (*Portainer, error) {
	return &Portainer{Url: url, Key: key}, nil
}

func (p *Portainer) GetStacks() ([]Stack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.Url+"/api/stacks", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, string(body))
	}
	defer resp.Body.Close()

	var stacks []Stack

	err = json.Unmarshal(body, &stacks)
	if err != nil {
		return nil, err
	}
	return stacks, nil
}

func (p *Portainer) GetStackFile(stack Stack) (*Stack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.Url+"/api/stacks/"+strconv.Itoa(stack.ID)+"/file", nil)
	if err != nil {
		return &stack, err
	}
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &stack, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &stack, err
	}
	if resp.StatusCode != http.StatusOK {
		return &stack, fmt.Errorf("%s: %s", resp.Status, string(body))
	}
	defer resp.Body.Close()
	err = json.Unmarshal(body, &stack)
	if err != nil {
		return &stack, err
	}
	return &stack, nil
}

func (p *Portainer) UpdateStack(stack *Stack) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	params := url.Values{}
	params.Add("endpointId", strconv.Itoa(stack.Endpoint))
	data := struct {
		Prune     bool   `json:"prune"`
		PullImage bool   `json:"pullImage"`
		File      string `json:"stackFileContent"`
	}{
		Prune:     true,
		PullImage: true,
		File:      stack.StackFile,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	requestBody := bytes.NewReader(jsonData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, p.Url+"/api/stacks/"+strconv.Itoa(stack.ID), requestBody)
	if err != nil {
		return err
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", resp.Status, string(body))
	}
	defer resp.Body.Close()
	// fmt.Println(string(body))
	return nil
}

func (p *Portainer) GetTemplates() ([]Stack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.Url+"/api/custom_templates", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, string(body))
	}
	defer resp.Body.Close()

	var stacks []Stack

	err = json.Unmarshal(body, &stacks)
	if err != nil {
		return nil, err
	}
	return stacks, nil
}

func (p *Portainer) GetTemplateFile(stack Stack) (*Stack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.Url+"/api/custom_templates/"+strconv.Itoa(stack.ID)+"/file", nil)
	if err != nil {
		return &stack, err
	}
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &stack, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &stack, err
	}
	if resp.StatusCode != http.StatusOK {
		return &stack, fmt.Errorf("%s: %s", resp.Status, string(body))
	}

	defer resp.Body.Close()
	err = json.Unmarshal(body, &stack)
	if err != nil {
		return &stack, err
	}
	return &stack, nil
}

func (p *Portainer) CreateStack(name, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	params := url.Values{}
	params.Add("endpointId", strconv.Itoa(8)) // dev
	data := struct {
		Name string `json:"name"`
		File string `json:"stackFileContent"`
	}{
		Name: name,
		File: content,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	requestBody := bytes.NewReader(jsonData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.Url+"/api/stacks/create/standalone/string", requestBody)
	if err != nil {
		return err
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", resp.Status, string(body))
	}
	defer resp.Body.Close()
	//fmt.Println(string(body))
	var postData struct {
		ResourceControl struct {
			Id int `json:"Id"`
		} `json:"ResourceControl"`
	}
	err = json.Unmarshal(body, &postData)
	if err != nil {
		return err
	}
	if err := p.updateResourceControl(postData.ResourceControl.Id); err != nil {
		return err
	}
	return nil
}

func (p *Portainer) updateResourceControl(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	data := struct {
		AdministratorsOnly bool  `json:"AdministratorsOnly"`
		Teams              []int `json:"teams"`
	}{
		AdministratorsOnly: false,
		Teams:              []int{5}, //dev team
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	requestBody := bytes.NewReader(jsonData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, p.Url+"/api/resource_controls/"+strconv.Itoa(id), requestBody)
	if err != nil {
		return err
	}
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", resp.Status, string(body))
	}
	defer resp.Body.Close()
	// fmt.Println(string(body))
	return nil
}
