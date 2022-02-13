package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Thalisonh/star-wars-api/models"
	"github.com/Thalisonh/star-wars-api/util"
)

type ISwapi interface {
	TotalFilms(planetName string) (int, int, error)
}

type SwapiService struct {
	client util.HttpClient
}

func NewSwapiService(client util.HttpClient) ISwapi {
	return &SwapiService{client}
}

func (s *SwapiService) TotalFilms(planetName string) (int, int, error) {
	url := fmt.Sprintf("https://swapi.dev/api/planets?search=%s", planetName)

	totalFilms := 0

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return http.StatusInternalServerError, totalFilms, err
	}
	req.Close = true

	resp, err := s.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, totalFilms, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, totalFilms, err
	}

	swapi := &models.Swapi{}
	if err := json.Unmarshal([]byte(body), swapi); err != nil {
		return http.StatusInternalServerError, totalFilms, err
	}

	for i := 0; i < len(swapi.Results); i++ {
		if strings.EqualFold(swapi.Results[i].Name, planetName) {
			totalFilms = len(swapi.Results[i].Films)
		}
	}

	return http.StatusOK, totalFilms, nil
}
