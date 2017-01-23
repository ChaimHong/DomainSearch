package ds

import (
	"io/ioutil"
	"net/http"
)

type Domain struct {
	Name  string
	Price float64
}

type ISearch interface {
	GetUrl(key string) string
	ParseBody(body []byte) []Domain
}

type SearchServer struct {
	Apis       []ISearch
	Chars      []rune   // 字母组合
	CharNum    int      // 几位数
	Suffix     []string // 后缀
	availables []Domain
}

func (ss *SearchServer) Do() []Domain {
	ss.availables = ss.availables[0:0]
	combines := GetCombineMatch(ss.Chars, ss.CharNum)
	for i := 0; i < len(ss.Apis); i++ {
		s := ss.Apis[i]

		for _, v := range combines {
			// go func(v string) {
			resp, err := http.Get(s.GetUrl(v + ".xyz"))
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()
			body, err2 := ioutil.ReadAll(resp.Body)
			if err2 != nil {
				panic(err2)
			}

			availables := s.ParseBody(body)
			ss.availables = append(ss.availables, availables...)
			// }(v)
		}
	}

	return ss.availables
}

func (ss *SearchServer) One(key string) []Domain {
	ss.availables = ss.availables[0:0]
	for i := 0; i < len(ss.Apis); i++ {
		s := ss.Apis[i]

		resp, err := http.Get(s.GetUrl(key + ".xyz"))

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			panic(err2)
		}

		availables := s.ParseBody(body)
		ss.availables = append(ss.availables, availables...)
	}

	return ss.availables
}

func (ss *SearchServer) loop() {

}
