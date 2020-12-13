package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, ""))
	// q := url.QueryEscape(terms)
	// l := terms[0]
	// k := terms[1]

	resp, err := http.Get(IssuesURL + "?q=" + q)
	// resp, err := http.Get(IssuesURL + "?location=" + l + "&key=" + k)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
