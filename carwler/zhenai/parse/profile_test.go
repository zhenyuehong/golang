package parse

import (
	"fmt"
	"golang/carwler/fetcher"
	"testing"
)

func TestParseProfile(t *testing.T) {
	bytes, _ := fetcher.Fetch("http://album.zhenai.com/u/1434875416")
	profile := ParseProfile(bytes)
	fmt.Printf(profile.Requests[0].Url)
}
