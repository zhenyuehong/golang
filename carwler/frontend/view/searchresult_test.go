package view

import (
	"golang/carwler/engine"
	"golang/carwler/frontend/model"
	common "golang/carwler/model" //两个model会冲突，所以要重命名
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	page := model.SearchResult{}
	//err := template.Execute(os.Stdout,page)
	//现在我们建一个测试的html文件
	out, err := os.Create("template.test.html")

	page.Hits = 1234
	items := engine.Item{
		Url:  "http://album.zhenai.com/u/1434875416",
		Type: "zhenai",
		Id:   "1434875416",
		Payload: common.Profile{
			Name:       "提线木偶",
			Gender:     "男",
			Age:        29,
			Height:     168,
			Weight:     "--",
			Income:     "3001-5000元",
			Marriage:   "未婚",
			Education:  "大学本科",
			Occupation: "运营管理",
			Hukou:      "辽宁鞍山",
			Xingzuo:    "天蝎座",
			House:      "--",
			Car:        "未购车",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, items)
	}

	err = view.Render(out, page)

	if err != nil {
		panic(err)
	}

}
