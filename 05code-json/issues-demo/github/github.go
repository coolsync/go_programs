// Package github provides a GO API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.

package github

import (
	"time"
)

// let httpUrl = "https://free-api.heweather.net/s6/weather/now?location=广州&key=3c497450d8e44c5280421ceaba1db581"
//  wsarecv: An existing connection was forcibly closed by the remote host
// const IssuesURL = "https://api.github.com/search/issues"
const IssuesURL = "https://free-api.heweather.net/s6/weather/now"



type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}