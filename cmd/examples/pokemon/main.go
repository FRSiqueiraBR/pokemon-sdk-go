package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/frsiqueira/pokemon-sdk-go/pkg/pokemon"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	service := pokemon.NewService(pokemon.Config{})

	result, err := service.GetPokemon(ctx, "pikachu", pokemon.FormatSummary)
	if err != nil {
		log.Fatalf("fetch pokemon: %v", err)
	}

	fmt.Printf("%+v\n", result)
}
