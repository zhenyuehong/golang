package persist

import (
	"context"
	"encoding/json"
	"golang/carwler/model"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestItemSaver(t *testing.T) {
	profile := model.Profile{
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
	}
	//存数据
	id, err := saveItem(profile)
	if err != nil {
		panic(err)
	}

	//todo try to start up elastic search
	//here using docker go client
	//获取数据
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual model.Profile
	//bytes, _ := resp.Source.MarshalJSON()
	//err = json.Unmarshal(bytes, &actual)
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	if actual != profile {
		t.Errorf("got %v , expected %v", actual, profile)
	}
}
