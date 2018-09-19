package model

type SearchResult struct {
	Hits  int64 // total
	Start int   //从第几个开始
	//Items []engine.Item
	Items    []interface{}
	Query    string
	PreFrom  int
	NextFrom int
}
