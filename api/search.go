package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const SEARCH_URL = "https://api.boardgameatlas.com/api/search"

// Like a class
type BoardgameAtlas struct {
	// 'members'
	clientId string
}

type Game struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	YearPublished uint   `json:"year_published"`
	Description   string `json:"description"`
	Url           string `json:"official_url"`
	ImageUrl      string `json:"image_url"`
	RulesUrl      string `json:"rules_url"`
}

type SearchResult struct {
	Games []Game `json:"games"`
	Count uint   `json:"count"`
}

// Functions as a constructor
func New(clientId string) BoardgameAtlas {
	return BoardgameAtlas{clientId: clientId}
}

// 'Method' in BoardgameAtlas
// b BoardgameAtlas -> denotes the receiver
// Return result => (*SearchResult, error)
// return nil, fmt.Errorf("Cannot create HTTP client: %v", err)

func (b BoardgameAtlas) Search(ctx context.Context, query string, limit uint, skip uint) (*SearchResult, error) {
	// Create Http Client
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SEARCH_URL, nil)

	if nil != err {
		// return an error object
		return nil, fmt.Errorf("Cannot create HTTP client: %v", err)
	}
	// Get the query string object
	qs := req.URL.Query()
	// populate the url with query params
	qs.Add("name", query)
	qs.Add("limit", fmt.Sprintf("%d", limit))
	qs.Add("skip", strconv.Itoa(int(skip)))
	qs.Add("client_id", b.clientId)

	// Encode the query params add
	req.URL.RawQuery = qs.Encode()

	// fmt.Printf("URL=%s\n", req.URL.String())

	// Make the call
	resp, err := http.DefaultClient.Do(req)

	if nil != err {
		return nil, fmt.Errorf("cannot create HTTP client for invocation: %v", err)
	}

	// HTTP status code >= 400 is error
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error HTTP status: %s", resp.Status)
	}

	var result SearchResult
	// Deserialize the JSON payload to Struct
	if err := json.NewDecoder(resp.Body).Decode(&result); nil != err {
		return nil, fmt.Errorf("cannot deserialize the JSON payload: %s", resp.Status)
	}

	return &result, nil
}
