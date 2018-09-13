package persist

import (
	"context"
	"encoding/json"
	"golang/carwler/engine"
	"golang/carwler/model"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestItemSaver(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1434875416",
		Type: "zhenai",
		Id:   "1434875416",
		Payload: model.Profile{
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

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"

	//存数据
	err = saveItem(client, index, expected)
	if err != nil {
		panic(err)
	}

	//todo try to start up elastic search
	//here using docker go client
	//获取数据

	resp, err := client.Get().Index(index).
		Type(expected.Type).Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual engine.Item
	//bytes, _ := resp.Source.MarshalJSON()
	//err = json.Unmarshal(bytes, &actual)
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	//Verify result
	if actual != expected {
		t.Errorf("got %v , expected %v", actual, expected)
	}
}
