package scrapbox

import "net/http"

// ListOptions is an option for getting PageList
type ListOptions struct {
	Skip  int `url:"skip,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// PageList is a collection of pagelist elements
type PageList struct {
	ProjectName string `json:"projectName"`
	Skip        int    `json:"skip"`
	Limit       int    `json:"limit"`
	Count       int    `json:"count"`
	Pages       []Page `json:"pages"`
}

// Page is a collection of pagelist elements
type Page struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Image           string   `json:"image"`
	Descriptions    []string `json:"descriptions"`
	User            UserID   `json:"user"`
	Pin             int      `json:"pin"`
	Views           int      `json:"views"`
	Linked          int      `json:"linked"`
	CommitID        string   `json:"commitId"`
	Created         int      `json:"created"`
	Updated         int      `json:"updated"`
	Accessed        int      `json:"accessed"`
	SnapshotCreated int      `json:"snapshotCreated"`
}

// UserID is a collection of pagelist elements
type UserID struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Photo       string `json:"photo"`
}

// PageList return pagelist from scrapbox API
func (s *Client) PageList(lo *ListOptions) (*PageList, *http.Response, error) {
	if lo == nil {
		lo = &ListOptions{Skip: defaultSkip, Limit: defaultLimit}
	}
	u, err := addOptions(s.ProjectName, lo)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	lists := new(PageList)
	resp, err := s.Do(req, lists)
	if err != nil {
		return nil, resp, err
	}
	return lists, resp, err
}
