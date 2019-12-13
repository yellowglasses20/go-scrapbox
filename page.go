package scrapbox

import (
	"fmt"
	"net/http"
)

// PageText is a collection of pagelist elements
type PageText struct {
	ID              string         `json:"id"`
	Title           string         `json:"title"`
	Image           string         `json:"image"`
	Descriptions    []string       `json:"descriptions"`
	User            UserID         `json:"user"`
	Pin             int            `json:"pin"`
	Views           int            `json:"views"`
	Linked          int            `json:"linked"`
	CommitID        string         `json:"commitId"`
	Created         int            `json:"created"`
	Updated         int            `json:"updated"`
	Accessed        int            `json:"accessed"`
	SnapshotCreated int            `json:"snapshotCreated"`
	Persistent      bool           `json:"persistent"`
	Lines           []Line         `json:"lines"`
	Links           []string       `json:"links"`
	Icons           struct{}       `json:"icons"`
	RelatedPages    RelatedPage    `json:"relatedPages"`
	Collaborators   []Collaborator `json:"collaborators"`
	LastAccessed    int            `json:"lastAccessed"`
}

// Line is a collection of pagelist elements
type Line struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	UserID  string `json:"userId"`
	Created int    `json:"created"`
	Updated int    `json:"updated"`
}

// RelatedPage is a collection of pagelist elements
type RelatedPage struct {
	Links1hop []Links1 `json:"links1hop"`
	// TODO: not implement
	Links2hop []interface{} `json:"links2hop"`
	Icons1hop []interface{} `json:"icons1hop"`
}

// Collaborator is a collection of pagelist elements
type Collaborator struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Photo       string `json:"photo"`
}

// Links1 is a collection of pagelist elements
type Links1 struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	TitleLc      string   `json:"titleLc"`
	Image        string   `json:"image"`
	Descriptions []string `json:"descriptions"`
	LinksLc      []string `json:"linksLc"`
	Updated      int      `json:"updated"`
	Accessed     int      `json:"accessed"`
}

// PageTitle return title from scrapbox API
func (s *Client) PageTitle(title string) (*PageText, *http.Response, error) {
	if title == "" {
		return nil, nil, fmt.Errorf("Title is a required parameter")
	}

	lo := &ListOptions{Skip: defaultSkip, Limit: defaultLimit}
	u, err := addOptions(s.ProjectName+"/"+title, lo)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	lists := new(PageText)
	resp, err := s.Do(req, lists)
	if err != nil {
		return nil, resp, err
	}
	return lists, resp, err
}

// PageText return text from scrapbox API
func (s *Client) PageText(title string) ([]byte, *http.Response, error) {
	if title == "" {
		return nil, nil, fmt.Errorf("Title is a required parameter")
	}

	lo := &ListOptions{Skip: defaultSkip, Limit: defaultLimit}
	u, err := addOptions(s.ProjectName+"/"+title+"/text", lo)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, body, err := s.Text(req)
	if err != nil {
		return nil, resp, err
	}
	return body, resp, err
}

// PageIcon return text from scrapbox API
func (s *Client) PageIcon(title string) ([]byte, *http.Response, error) {
	if title == "" {
		return nil, nil, fmt.Errorf("Title is a required parameter")
	}

	lo := &ListOptions{Skip: defaultSkip, Limit: defaultLimit}
	u, err := addOptions(s.ProjectName+"/"+title+"/icon", lo)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, body, err := s.Text(req)
	if err != nil {
		return nil, resp, err
	}
	return body, resp, err
}
