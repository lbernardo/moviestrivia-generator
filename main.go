package main

import (
	"encoding/json"
	"github.com/lbernardo/movidestrivia-generator/pkg/logic"
	"github.com/lbernardo/movidestrivia-generator/pkg/types"
	"os"
)

func main() {
	var cfg types.Config
	content, _ := os.ReadFile("config.json")
	json.Unmarshal(content, &cfg)
	logic.NewGetMovies(cfg)
}
