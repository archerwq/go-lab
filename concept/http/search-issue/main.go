package main

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {

}