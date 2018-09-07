package parse

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	//为了防止网址不可用，所以我们使用拷贝的内容来进行测试
	contents, err := ioutil.ReadFile("city_test_data.html")

	if err != nil {
		panic(err)
	}
	//fmt.Println(string(contents))

	//ParseCityList(contents)
	//verify data
	result := ParseCityList(contents)

	//验证数目
	expectedSize := 470
	if len(result.Requests) != expectedSize {
		t.Errorf("result should have %d requests,but have %d", expectedSize, len(result.Requests))
	}

	if len(result.Requests) != expectedSize {
		t.Errorf("result should have %d items,but have %d", expectedSize, len(result.Items))
	}
	//验证结果
	expectUrl := []string{"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng"}
	expectItems := []string{"阿坝", "阿克苏", "阿拉善盟"}
	for i, val := range expectUrl {
		if result.Requests[i].Url != val {
			t.Errorf("expected url %d : %s ; but was %s",
				i, val, result.Requests[i].Url)
		}
	}

	for i, val := range expectItems {
		if result.Items[i].(string) != val {
			t.Errorf("expected city %d : %s ; but was %s",
				i, val, result.Items[i].(string))
		}
	}
}
