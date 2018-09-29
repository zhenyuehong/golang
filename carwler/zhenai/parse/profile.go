package parse

import (
	"golang/carwler/engine"
	"golang/carwler/model"
	"regexp"
	"strconv"
)

//`<td><span class = "label">年龄：</span>29岁</td>`
//`<td><span class="label">婚况：</span>未婚</td>`

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

var guessRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}

	if height, err := strconv.Atoi(extractString(contents, heightRe)); err == nil {
		profile.Height = height
	}

	//if weight, err := strconv.Atoi(extractString(contents, weightRe)); err == nil {
	//	profile.Weight = weight
	//}
	profile.Weight = extractString(contents, weightRe)
	profile.Name = name
	profile.Gender = extractString(contents, genderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		////这里m[2]是拷贝出来的，为了解决m(2) 都只指向一个人的问题
		//name := string(m[2])
		////url 同理
		//url := string(m[1])
		////result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			//Url: url,
			//ParserFunc:nil,//这里要进行下一个页面的抓取，这里为了先让他编译通过，暂时设置为nil
			//ParserFunc: func(c []byte) engine.ParseResult {
			//			//	return ParseProfile(c, url, name)
			//			//	//这里m[2] 不是马上运行，而是等到这个循环结束后才排到它，所以在这里用M(2)，早就不是指向这个人了
			//			//	//所以要把M(2)拷贝出来 name := string(m[2])
			//			//},
			//ParserFunc: ProfileParser(url, name),

			Url: string(m[1]),
			//Parser: NewProfileParser(string(m[2])),
			ParserFunc: ProfileParser(string(m[2])),
		})
	}
	return result

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return parseProfile(c, url, name)
	}
}

//
//type ProfileParser struct {
//	userName string
//}
//
//func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
//	return parseProfile(contents, url, p.userName)
//}
//
//func (p *ProfileParser) Serialize() (name string, args interface{}) {
//	return "ProfileParser", p.userName
//}
//
//func NewProfileParser(name string) *ProfileParser {
//	return &ProfileParser{
//		userName: name,
//	}
//}
