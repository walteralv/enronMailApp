package adapter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	zincSearchHost    = "http://localhost:4080"
	defaultSearchType = "matchphrase"
	defaultMaxResults = 10
)

// For more information read https://docs.zincsearch.com/api/search/search/#golang-example
// Request struct to consult ZincSearch API
type SearchEmailRequest struct {
	SearchType string                  `json:"search_type"`
	SortFields []string                `json:"sort_fields"`
	From       int                     `json:"from"`
	MaxResults int                     `json:"max_results"`
	Query      SearchEmailRequestQuery `json:"query"`
	Source     map[string]interface{}  `json:"_source"`
}

type SearchEmailRequestQuery struct {
	Term string `json:"term"`
}

// For more information read https://docs.zincsearch.com/api/search/search/#golang-example
// Response struct to ZincSearch API response
type SearchEmailResponse struct {
	Hits struct {
		Hits []struct {
			ID        string                 `json:"_id"`
			Score     float64                `json:"_score"`
			Source    map[string]interface{} `json:"_source"`
			Timestamp string                 `json:"@timestamp"`
		} `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
	} `json:"hits"`
	TimedOut bool    `json:"timed_out"`
	Took     float64 `json:"took"`
}

// set basic header to request
func setBasicHeaders(req *http.Request) {
	username := os.Getenv("ZINC_FIRST_ADMIN_USER")
	password := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	if username == "" || password == "" {
		panic("No ZINC_FIRST_ADMIN_USER and ZINC_FIRST_ADMIN_PASSWORD specified, please set ")
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

}

// For more information read https://docs.zincsearch.com/api/search/search/#golang-example
// searches emails with the Search ZincSearch API
func SearchEmail(indexName, term string) (*SearchEmailResponse, error) {
	response := &SearchEmailResponse{}

	path := fmt.Sprintf("%s/api/%s/_search", zincSearchHost, indexName)

	searchEmailRequest := SearchEmailRequest{
		SearchType: defaultSearchType,
		Query: SearchEmailRequestQuery{
			Term: term,
		},
		SortFields: []string{"@timestamp"},
		From:       0,
		MaxResults: defaultMaxResults,
	}

	query, _ := json.Marshal(searchEmailRequest)

	req, err := http.NewRequest("POST", path, strings.NewReader(string(query)))
	setBasicHeaders(req)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error searching documents: %s", err)
	}

	return response, nil
}
