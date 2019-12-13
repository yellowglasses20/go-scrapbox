package scrapbox

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPageTitle(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"limit": "100",
		})
		fmt.Fprint(w, `
      {
      	"id": "58e67688d0a4fe0011a0249c",
      	"title": "API",
      	"image": "https://gyazo.com/5bf55bb6223a62bf4477f07a9aad39b8/raw",
      	"descriptions": [
      		"[https://gyazo.com/5bf55bb6223a62bf4477f07a9aad39b8]",
      		"あくまで内部APIです。APIは予告なく変更を行います。",
      		"ページデータを取得するAPI",
      		"ページリスト",
      		"/api/pages/:projectName"
      	],
      	"user": {
      		"id": "5724627723541f110097c291",
      		"name": "shokai",
      		"displayName": "Sho Hashimoto",
      		"photo": "https://lh3.googleusercontent.com/a-/AAuE7mDKK7El_Vok1GPDG61ZLPJzyr4AVLqTh1jEKosc5Q"
      	},
      	"pin": 0,
      	"views": 8661,
      	"linked": 1,
      	"commitId": "5d135304ff5b5d0017cd83cc",
      	"created": 1491498636,
      	"updated": 1561547524,
      	"accessed": 1576195133,
      	"snapshotCreated": 1561586966,
      	"persistent": true,
      	"lines": [{
      		"id": "58e67688d0a4fe0011a0249c",
      		"text": "API",
      		"userId": "5724627723541f110097c291",
      		"created": 1491498636,
      		"updated": 1491499158
      	}],
      	"links": [
      		"ページを作る"
      	],
      	"icons": {},
      	"relatedPages": {
      		"links1hop": [{
      			"id": "58ae7a998f0baf001189e1ea",
      			"title": "ページを作る",
      			"titleLc": "ページを作る",
      			"image": "https://gyazo.com/984e915a6716b9d2f787c3e4688057c0/raw",
      			"descriptions": [
      				"[https://gyazo.com/984e915a6716b9d2f787c3e4688057c0]",
      				"ページリスト上部にある新規作成ボタンから、ページを作成できます。",
      				"タイトル行は空行のままでも問題ありません。",
      				"2行目以降からタイトル候補を探して、ページリスト画面で表示します。",
      				"タイトルは重複しても大丈夫です。"
      			],
      			"linksLc": [
      				"ページをリンクする",
      				"ページをインポート・エクスポートする",
      				"リンク",
      				"記法",
      				"bookmarkletから投稿する",
      				"ページ",
      				"編集方法"
      			],
      			"updated": 1537354933,
      			"accessed": 1576174372
      		}],
      		"links2hop": [],
      		"icons1hop": []
      	},
      	"collaborators": [{
      		"id": "5714528c2ec4dd1100efa18e",
      		"name": "rakusai",
      		"displayName": "Isshu Rakusai",
      		"photo": "https://lh3.googleusercontent.com/a-/AAuE7mBemsbJUyngbd1R2-CINIQKS3lfFhesBLi4NBgw"
      	}],
      	"lastAccessed": 1576053005
      }
    `)
	})

	actual, _, err := client.PageTitle("API")

	if err != nil {
		t.Fatalf("Testlist returned error: %v", err)
	}

	want := &PageText{
		ID:    "58e67688d0a4fe0011a0249c",
		Title: "API",
		Image: "https://gyazo.com/5bf55bb6223a62bf4477f07a9aad39b8/raw",
		Descriptions: []string{
			"[https://gyazo.com/5bf55bb6223a62bf4477f07a9aad39b8]",
			"あくまで内部APIです。APIは予告なく変更を行います。",
			"ページデータを取得するAPI",
			"ページリスト",
			"/api/pages/:projectName"},
		User: UserID{
			ID:          "5724627723541f110097c291",
			Name:        "shokai",
			DisplayName: "Sho Hashimoto",
			Photo:       "https://lh3.googleusercontent.com/a-/AAuE7mDKK7El_Vok1GPDG61ZLPJzyr4AVLqTh1jEKosc5Q",
		},
		Pin:             0,
		Views:           8661,
		Linked:          1,
		CommitID:        "5d135304ff5b5d0017cd83cc",
		Created:         1491498636,
		Updated:         1561547524,
		Accessed:        1576195133,
		SnapshotCreated: 1561586966,
		Persistent:      true,
		Lines: []Line{{
			ID:      "58e67688d0a4fe0011a0249c",
			Text:    "API",
			UserID:  "5724627723541f110097c291",
			Created: 1491498636,
			Updated: 1491499158,
		}},
		Links: []string{
			"ページを作る",
		},
		RelatedPages: RelatedPage{
			Links1hop: []Links1{{
				ID:      "58ae7a998f0baf001189e1ea",
				Title:   "ページを作る",
				TitleLc: "ページを作る",
				Image:   "https://gyazo.com/984e915a6716b9d2f787c3e4688057c0/raw",
				Descriptions: []string{
					"[https://gyazo.com/984e915a6716b9d2f787c3e4688057c0]",
					"ページリスト上部にある新規作成ボタンから、ページを作成できます。",
					"タイトル行は空行のままでも問題ありません。",
					"2行目以降からタイトル候補を探して、ページリスト画面で表示します。",
					"タイトルは重複しても大丈夫です。"},
				LinksLc: []string{
					"ページをリンクする",
					"ページをインポート・エクスポートする",
					"リンク",
					"記法",
					"bookmarkletから投稿する",
					"ページ",
					"編集方法"},
				Updated:  1537354933,
				Accessed: 1576174372,
			}},
			Links2hop: []interface{}{},
			Icons1hop: []interface{}{},
		},
		Collaborators: []Collaborator{{
			ID:          "5714528c2ec4dd1100efa18e",
			Name:        "rakusai",
			DisplayName: "Isshu Rakusai",
			Photo:       "https://lh3.googleusercontent.com/a-/AAuE7mBemsbJUyngbd1R2-CINIQKS3lfFhesBLi4NBgw",
		}},
		LastAccessed: 1576053005,
	}

	if !reflect.DeepEqual(actual, want) {
		t.Errorf("Issues.List returned %+v, want %+v", actual, want)
	}
}

func TestPageText(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"limit": "100",
		})
		fmt.Fprint(w, `
[https://gyazo.com/5bf55bb6223a62bf4477f07a9aad39b8]

あくまで内部APIです。APIは予告なく変更を行います。

ページデータを取得するAPI
ページリスト
/api/pages/:projectName
パラメータ
skip 何番目から取得するかを指定。デフォルトは0
limit 取得するページ数。デフォルトは100
ページ
/api/pages/:projectName/:pageTitle
ページ本文
/api/pages/:projectName/:pageTitle/text
ページタイトル画像
/api/pages/:projectName/:pageTitle/icon

ページデータを書き込むAPI
作成中
[ページを作る#58ae7c9a97c29100005b886b]などの手段もあります`)
	})
	actual, _, err := client.PageText("API")

	t.Logf("%s", actual)

	if err != nil {
		t.Fatalf("Testlist returned error: %v", err)
	}

	want := []byte(`
[https://gyazo.com/5bf55bb6223a62bf4477f07a9aad39b8]

あくまで内部APIです。APIは予告なく変更を行います。

ページデータを取得するAPI
ページリスト
/api/pages/:projectName
パラメータ
skip 何番目から取得するかを指定。デフォルトは0
limit 取得するページ数。デフォルトは100
ページ
/api/pages/:projectName/:pageTitle
ページ本文
/api/pages/:projectName/:pageTitle/text
ページタイトル画像
/api/pages/:projectName/:pageTitle/icon

ページデータを書き込むAPI
作成中
[ページを作る#58ae7c9a97c29100005b886b]などの手段もあります`)

	if !reflect.DeepEqual(actual, want) {
		t.Errorf("Issues.List returned %+v, want %+v", actual, want)
	}
}

// func TestPageIcon(t *testing.T) {
// 	ac, _ := NewClient("help-jp")
// 	actual, _, err := ac.PageIcon("API")

// 	t.Logf("%s", actual)

// 	if err != nil {
// 		t.Fatalf("Testlist returned error: %v", err)
// 	}
// }
