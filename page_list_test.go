package scrapbox

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"skip":  "1",
			"limit": "1",
		})
		fmt.Fprint(w, `
      {
        "ProjectName": "help-jp",
        "Skip": 0,
        "Limit": 1,
        "Count": 78,
        "Pages": [{
          "ID": "57c7d72ad25ef00f00100688",
          "Title": "Scrapboxの使い方",
          "Image": "https://gyazo.com/7057219f5b20ca8afd122945b72453d3/raw",
          "Descriptions": [
            "[https://gyazo.com/7057219f5b20ca8afd122945b72453d3]",
            "Scrapbox（スクラップボックス）の使い方・活用方法についてご紹介します。",
            "[* 編集を始める]", 
            "まず最初に[ブラケティング]を読んでみましょう",
            "[* チュートリアルで雰囲気をつかむ]"
          ],
          "User": {
            "ID": "566f8b954fb08e1100af5c5b"
          },
          "Pin": 9007197717386014,
          "Views": 46459,
          "Linked": 2,
          "CommitID": "5cf9c78742fc7800179f8c19",
          "Created": 1472713402,
          "Updated": 1559271981,
          "Accessed": 1576036512,
          "SnapshotCreated": 1559315996
        }]
      }
    `)
	})

	actual, _, err := client.PageList(&ListOptions{Skip: 1, Limit: 1})

	if err != nil {
		t.Fatalf("Testlist returned error: %v", err)
	}

	want := &PageList{
		ProjectName: "help-jp",
		Skip:        0,
		Limit:       1,
		Count:       78,
		Pages: []Pages{{
			ID:    "57c7d72ad25ef00f00100688",
			Title: "Scrapboxの使い方",
			Image: "https://gyazo.com/7057219f5b20ca8afd122945b72453d3/raw",
			Descriptions: []string{
				"[https://gyazo.com/7057219f5b20ca8afd122945b72453d3]",
				"Scrapbox（スクラップボックス）の使い方・活用方法についてご紹介します。",
				"[* 編集を始める]",
				"まず最初に[ブラケティング]を読んでみましょう",
				"[* チュートリアルで雰囲気をつかむ]"},
			User: UserID{
				ID: "566f8b954fb08e1100af5c5b"},
			Pin:             9007197717386014,
			Views:           46459,
			Linked:          2,
			CommitID:        "5cf9c78742fc7800179f8c19",
			Created:         1472713402,
			Updated:         1559271981,
			Accessed:        1576036512,
			SnapshotCreated: 1559315996}}}

	if !reflect.DeepEqual(actual, want) {
		t.Errorf("Issues.List returned %+v, want %+v", actual, want)
	}
}
