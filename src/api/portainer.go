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

	"github.com/PavelMilanov/forge/models"
)

/*
PortainerAPI предоставляет интерфейс для взаимодействия с Portainer.
*/
type Portainer struct {
	Url   string
	Key   string
	Teams []int
}

type StackMode string

const (
	StackModeCreate StackMode = "create"
	StackModeUpdate StackMode = "update"
	StackModeUpsert StackMode = "upsert"
)

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
func NewPortainer(url, key string, teams []int) (*Portainer, error) {
	if url == "" {
		return nil, fmt.Errorf("portainer url is empty")
	}
	if key == "" {
		return nil, fmt.Errorf("portainer api key is empty")
	}
	return &Portainer{Url: url, Key: key, Teams: teams}, nil
}

/*
GetStacks получает список стеков из Portainer API.

Returns

	[]Stack - список стеков.
	error - ошибка запроса или декодирования ответа.
*/
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

/*
GetStackFile получает файл конфигурации для указанного стека.

Params

	stack - метаданные стека (используется ID).

Returns

	*Stack - стек с заполненным полем StackFile.
	error - ошибка запроса или декодирования ответа.
*/
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

/*
UpdateStack обновляет стек его текущим содержимым.

Params

	stack - стек с заполненными ID, Endpoint и StackFile.

Returns

	error - ошибка обновления.
*/
func (p *Portainer) UpdateStack(stack *Stack) error {
	return p.UpdateStackContent(stack.ID, stack.Endpoint, stack.StackFile, true, true)
}

/*
UpdateStackContent обновляет содержимое стека указанным контентом.

Params

	stackID - идентификатор стека в Portainer.
	endpointID - идентификатор endpoint-а, к которому привязан стек.
	content - содержимое stack-файла.
	prune - удалять неиспользуемые ресурсы.
	pullImage - принудительно подтягивать образы.

Returns

	error - ошибка обновления.
*/
func (p *Portainer) UpdateStackContent(stackID, endpointID int, content string, prune, pullImage bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	params := url.Values{}
	params.Add("endpointId", strconv.Itoa(endpointID))
	data := struct {
		Prune     bool   `json:"prune"`
		PullImage bool   `json:"pullImage"`
		File      string `json:"stackFileContent"`
	}{
		Prune:     prune,
		PullImage: pullImage,
		File:      content,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	requestBody := bytes.NewReader(jsonData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, p.Url+"/api/stacks/"+strconv.Itoa(stackID), requestBody)
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

/*
GetTemplates получает список custom templates из Portainer.

Returns

	[]Stack - список template-объектов.
	error - ошибка запроса или декодирования ответа.
*/
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

/*
GetTemplateFile получает содержимое файла custom template.

Params

	stack - template-объект (используется ID).

Returns

	*Stack - template-объект с заполненным полем TemplateFile.
	error - ошибка запроса или декодирования ответа.
*/
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

/*
CreateStack создает новый standalone stack из строкового контента.

Params

	endpointId - идентификатор endpoint-а.
	name - имя стека.
	content - содержимое stack-файла.

Returns

	error - ошибка создания стека.
*/
func (p *Portainer) CreateStack(endpointId int, name, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	params := url.Values{}
	params.Add("endpointId", strconv.Itoa(endpointId)) // dev
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
	var postData struct {
		ResourceControl struct {
			Id int `json:"Id"`
		} `json:"ResourceControl"`
	}
	err = json.Unmarshal(body, &postData)
	if err != nil {
		return err
	}
	if len(p.Teams) == 0 {
		return nil
	}
	if err := p.updateResourceControl(postData.ResourceControl.Id); err != nil {
		return err
	}
	return nil
}

/*
GetEndpoints получает список endpoint-ов Portainer.

Returns

	[]models.PortainerEndpoint - список доступных endpoint-ов.
	error - ошибка запроса или декодирования ответа.
*/
func (p *Portainer) GetEndpoints() ([]models.PortainerEndpoint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.Url+"/api/endpoints", nil)
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
	var data []models.PortainerEndpoint
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
GetEndpointByName ищет endpoint по имени.

Params

	name - имя endpoint-а.

Returns

	*models.PortainerEndpoint - найденный endpoint.
	bool - флаг, найден ли endpoint.
	error - ошибка запроса endpoint-ов.
*/
func (p *Portainer) GetEndpointByName(name string) (*models.PortainerEndpoint, bool, error) {
	endpoints, err := p.GetEndpoints()
	if err != nil {
		return nil, false, err
	}
	for i := range endpoints {
		if endpoints[i].Name == name {
			return &endpoints[i], true, nil
		}
	}
	return nil, false, nil
}

/*
GetStackByName ищет стек по имени.

Params

	name - имя стека.

Returns

	*Stack - найденный стек.
	bool - флаг, найден ли стек.
	error - ошибка запроса списка стеков.
*/
func (p *Portainer) GetStackByName(name string) (*Stack, bool, error) {
	stacks, err := p.GetStacks()
	if err != nil {
		return nil, false, err
	}
	for i := range stacks {
		if stacks[i].StackName == name {
			return &stacks[i], true, nil
		}
	}
	return nil, false, nil
}

/*
DeployStackFromContent выполняет деплой стека из готового контента в режимах create/update/upsert.

Params

	endpointID - идентификатор endpoint-а.
	stackName - имя стека.
	content - содержимое stack-файла.
	mode - режим деплоя: create | update | upsert.
	prune - удалять неиспользуемые ресурсы при update.
	pullImage - принудительно подтягивать образы при update.

Returns

	string - результат операции: "created" или "updated".
	error - ошибка валидации или вызова Portainer API.
*/
func (p *Portainer) DeployStackFromContent(endpointID int, stackName, content string, mode StackMode, prune, pullImage bool) (string, error) {
	if stackName == "" {
		return "", fmt.Errorf("stack name is required")
	}
	if content == "" {
		return "", fmt.Errorf("stack content is empty")
	}
	switch mode {
	case StackModeCreate, StackModeUpdate, StackModeUpsert:
	default:
		return "", fmt.Errorf("unknown deploy mode %q", mode)
	}

	existing, found, err := p.GetStackByName(stackName)
	if err != nil {
		return "", err
	}

	switch mode {
	case StackModeCreate:
		if found {
			return "", fmt.Errorf("stack %q already exists", stackName)
		}
		if err := p.CreateStack(endpointID, stackName, content); err != nil {
			return "", err
		}
		return "created", nil
	case StackModeUpdate:
		if !found {
			return "", fmt.Errorf("stack %q not found", stackName)
		}
		if existing.Endpoint != endpointID {
			return "", fmt.Errorf("stack %q exists on endpointId=%d, not on selected endpointId=%d", stackName, existing.Endpoint, endpointID)
		}
		if err := p.UpdateStackContent(existing.ID, endpointID, content, prune, pullImage); err != nil {
			return "", err
		}
		return "updated", nil
	default: // upsert
		if found {
			if existing.Endpoint != endpointID {
				return "", fmt.Errorf("stack %q exists on endpointId=%d, not on selected endpointId=%d", stackName, existing.Endpoint, endpointID)
			}
			if err := p.UpdateStackContent(existing.ID, endpointID, content, prune, pullImage); err != nil {
				return "", err
			}
			return "updated", nil
		}
		if err := p.CreateStack(endpointID, stackName, content); err != nil {
			return "", err
		}
		return "created", nil
	}
}

/*
RefreshStack повторно применяет текущий stack-файл существующего стека.

Params

	endpointID - идентификатор endpoint-а.
	stackName - имя существующего стека.
	prune - удалять неиспользуемые ресурсы.
	pullImage - принудительно подтягивать образы.

Returns

	error - ошибка получения или обновления стека.
*/
func (p *Portainer) RefreshStack(endpointID int, stackName string, prune, pullImage bool) error {
	if stackName == "" {
		return fmt.Errorf("stack name is required")
	}
	existing, found, err := p.GetStackByName(stackName)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("stack %q not found", stackName)
	}
	if existing.Endpoint != endpointID {
		return fmt.Errorf("stack %q exists on endpointId=%d, not on selected endpointId=%d", stackName, existing.Endpoint, endpointID)
	}
	project, err := p.GetStackFile(*existing)
	if err != nil {
		return err
	}
	return p.UpdateStackContent(existing.ID, endpointID, project.StackFile, prune, pullImage)
}

/*
GetTemplateByName ищет custom template по имени.

Params

	name - имя template.

Returns

	*Stack - найденный template-объект.
	bool - флаг, найден ли template.
	error - ошибка запроса списка шаблонов.
*/
func (p *Portainer) GetTemplateByName(name string) (*Stack, bool, error) {
	tmpls, err := p.GetTemplates()
	if err != nil {
		return nil, false, err
	}
	for i := range tmpls {
		if tmpls[i].TemplateName == name {
			return &tmpls[i], true, nil
		}
	}
	return nil, false, nil
}

/*
updateResourceControl обновляет права доступа для созданного стека.

Params

	id - идентификатор ResourceControl.

Returns

	error - ошибка запроса обновления прав.
*/
func (p *Portainer) updateResourceControl(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	data := struct {
		AdministratorsOnly bool  `json:"AdministratorsOnly"`
		Teams              []int `json:"teams"`
	}{
		AdministratorsOnly: false,
		Teams:              p.Teams,
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
	return nil
}
