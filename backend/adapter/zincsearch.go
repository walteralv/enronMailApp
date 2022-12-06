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
	zincSearchHost = "http://localhost:4080"
)

type SearchDocumentsResponse struct {
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

// SearchDocuments searches documents with the Search ZincSearch API
func SearchDocuments(indexName, term string) (*SearchDocumentsResponse, error) {
	response := &SearchDocumentsResponse{}

	path := fmt.Sprintf("%s/api/%s/_search", zincSearchHost, indexName)

	query := fmt.Sprintf(`{
        "search_type": "match",
        "query":
        {
            "term": "%s"
        },
        "from": 0,
        "max_results": 5,
        "_source": []
    }`, term)

	req, err := http.NewRequest("POST", path, strings.NewReader(query))
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
