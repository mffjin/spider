package parser 
import (
"regexp"
"engine"
)

var (
profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity (contents []byte, _ string) engine.ParseResult {
    matches := profileRe.FindAllSubmatch(contents,-1)

    result := engine.ParseResult{}
    for _,m := range matches {
		url := string(m[1])
		name := string(m[2])
        //result.Items = append(result.Items,"User " + name)
        result.Requests = append(result.Requests,engine.Request{
            Url: url,
            ParserFunc: ProfileParser(name),
        })
    }
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),	
			ParserFunc: ParseCity,
		})
	}
    return result	
}
