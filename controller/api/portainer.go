package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
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

// /api/stacks
// /api/stacks/11/file
// /api/stacks/11?endpointId=8

func NewPortainer(env *config.Env) (*Portainer, error) {
	return &Portainer{Url: env.Portainer.Url, Key: env.Portainer.Key}, nil
}

func (p *Portainer) GetStacks() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.Url+"/api/stacks", nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-API-Key", p.Key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return err
	}
	fmt.Println(string(body))
	return nil
}
