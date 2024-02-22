package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lbernardo/movidestrivia-generator/pkg/types"
	"net/http"
)

type Service struct {
	cfg    types.Config
	client *http.Client
}

func NewService(config types.Config) *Service {
	return &Service{
		cfg:    config,
		client: new(http.Client),
	}
}

func (s *Service) Auth() bool {
	req, _ := http.NewRequest(http.MethodGet, s.cfg.BaseUrl+"/authentication", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.cfg.Token)
	res, _ := s.client.Do(req)
	if res.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func (s *Service) GetMovies(page int) []types.Movie {
	var data = struct {
		Results []types.Movie `json:"results"`
	}{}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(s.cfg.BaseUrl+"/discover/movie?include_adult=false&include_video=false&language=pt-BR&sort_by=popularity.desc&page=%v&primary_release_date.gte=1995-01-01&primary_release_date.lte=2020-12-31&with_original_language=en", page), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.cfg.Token)
	res, _ := s.client.Do(req)
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	_ = json.Unmarshal(buf.Bytes(), &data)
	return data.Results
}
