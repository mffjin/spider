package parser

import (
"regexp"
"engine"
"strconv"
"model"
)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

var guessRe = regexp.MustCompile(`<a class="exp-user-name" target="_blank" [^>]*href="(http://album.zhenai.com/u/[0-9]+)">([^<]+)</a>`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)


func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}	

	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil{
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil{
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil{
		profile.Weight = weight
	}
	profile.Marriage = extractString(contents,marriageRe)
	profile.Income = extractString(contents,incomeRe)
	profile.Gender = extractString(contents,genderRe)
	profile.Education = extractString(contents,educationRe)
	profile.Occupation = extractString(contents,occupationRe)
	profile.Hukou = extractString(contents,hukouRe)
	profile.Xingzuo = extractString(contents,xingzuoRe)
	profile.House = extractString(contents,houseRe)
	profile.Car = extractString(contents, carRe)
	
	//Type: "zhenai",
	result := engine.ParseResult {
		Items: []engine.Item{
			{
				Url: url,
				Id: extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	//guess
	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		result.Requests = append(result.Requests,
			engine.Request{
				Url: url,
				ParserFunc: ProfileParser(name),
			})
	}
	return result
}

func extractString(contents []byte, re * regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}else{
		return ""
	}
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
