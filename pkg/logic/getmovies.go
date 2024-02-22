package logic

import (
	"encoding/json"
	"fmt"
	"github.com/lbernardo/movidestrivia-generator/pkg/service"
	"github.com/lbernardo/movidestrivia-generator/pkg/types"
	"os"
	"strconv"
	"strings"
)

func NewGetMovies(config types.Config) {
	svc := service.NewService(config)
	if svc.Auth() {
		fmt.Println("Authenticated")
	}
	var movies = []types.Movie{}
	for i := 1; i < 20; i++ {
		m := svc.GetMovies(i)
		for _, movie := range m {
			s, _ := strconv.Atoi(movie.ReleaseDate[:4])
			movie.Year = s
			movie.ID = strings.ReplaceAll(movie.ReleaseDate, "-", "")
			movies = append(movies, movie)
		}
	}
	content, _ := json.MarshalIndent(movies, "", "  ")
	os.WriteFile("result.json", content, 0644)
}
