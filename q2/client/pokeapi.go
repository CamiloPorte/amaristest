package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PokeApi interface {
	Get(requestData int) (string, error)
}

type pokeApi struct {
	URL     string
	Timeout time.Duration
}

func NewPokeApi(url string) PokeApi {
	return &pokeApi{
		URL: url,
	}
}

func (p *pokeApi) Get(requestData int) (string, error) {
	Url := fmt.Sprintf("%s%d", p.URL, requestData)
	req, err := http.NewRequest(http.MethodGet, Url, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", err
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var answ responseReq
	json.Unmarshal(resBody, &answ)
	return answ.Pokemon.Name, nil
}
