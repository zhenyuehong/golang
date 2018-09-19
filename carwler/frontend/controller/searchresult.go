package controller

import (
	"context"
	"golang/carwler/engine"
	"golang/carwler/frontend/model"
	"golang/carwler/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//TOD fill in query string
//support search button
//rewrite query string
//support paging
//add start page
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

////localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w, "q = %s, from = %d", q, from)
	page, err := h.getSearchRequest(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//注意，這裡要加這行，不然打印出來的就是html代碼，This is because it received no content type from the server. You need to set it in the header.
	w.Header().Set("Content-Type", "text/html")
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateSearchResultHandler(temple string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(temple),
		client: client,
	}

}

func (h SearchResultHandler) getSearchRequest(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PreFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
