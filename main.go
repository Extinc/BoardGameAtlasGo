package main

import (
	"context"
	"flag"
	"fmt"
	"github/LearningGoLang/api"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {

	// boardgameatlas --query "ticket to ride" --clientId abc123 --skip 1- --limit 10
	// Define the command line arguments
	query := flag.String("query", "", "Boardgame name to search")
	clientId := flag.String("clientId", "", "Boardgame atlas clientId")
	skip := flag.Uint("skip", 10, "Skip the number of results provided")
	limit := flag.Uint("limit", 5, "limit the number of results returned")
	timeout := flag.Uint("timeout", 10, "Timeout")
	// Parse the command line arguments
	flag.Parse()

	if isNull(*query) {
		log.Fatalln("Please use --query to set the boardgame name to search")
	}

	if isNull(*clientId) {
		log.Fatalln("Please use --clientId to set the Boardgame Atlas client_id")
	}
	// fmt.Printf("query=%s, clientId=%s, limit=%d, skip=%d ", *query, *clientId, *limit, *skip)

	// Instantiate a BoardgameAtlas struct
	bga := api.New(*clientId)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout*uint(time.Second)))
	defer cancel()

	// time.Sleep(3 * time.Second)

	result, err := bga.Search(ctx, *query, *limit, *skip)

	if nil != err {
		log.Fatalf("Cannot search for boardgame: %v", err)
	}

	// Colors
	boldGreen := color.New(color.Bold).Add(color.FgHiGreen).SprintFunc()

	for _, g := range result.Games {
		fmt.Printf("%s: %s\n", boldGreen("Name"), g.Name)
		fmt.Printf("%s: %s\n", boldGreen("Description"), g.Description)
		fmt.Printf("%s: %s\n\n", boldGreen("URL"), g.Url)
	}

}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
