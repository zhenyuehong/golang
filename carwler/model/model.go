package model

import "encoding/json"

//用户信息字段
type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     string
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hukou      string //籍贯
	Xingzuo    string
	House      string
	Car        string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
