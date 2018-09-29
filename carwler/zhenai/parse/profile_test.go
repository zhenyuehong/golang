package parse

import (
	"golang/carwler/engine"
	"golang/carwler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	bytes, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := parseProfile(bytes, "http://album.zhenai.com/u/1434875416", "提线木偶")

	if len(result.Items) != 1 {
		t.Errorf("item should contain 1 element; but was %v", result.Items)
	}

	//fmt.Println(string(bytes))

	profile := result.Items[0]

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

	if profile != expected {
		t.Errorf("expected %v ,but was %v", expected, profile)
	}
}
