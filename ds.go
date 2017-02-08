package ds

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const RUN_STEP = 3

type Domain struct {
	Name  string
	Price float64
}

type ISearch interface {
	GetUrl(key string) string
	ParseBody(body []byte) []Domain
}

type DomainChan struct {
	api    ISearch
	domain string
}

type SearchServer struct {
	Apis        []ISearch
	Chars       []rune          // 字母组合
	CharNum     int             // 几位数
	Suffix      []string        // 后缀
	availables  []Domain        // 可用的域名
	waitGroup   *sync.WaitGroup // 批量跑
	concurrency int

	domainChan chan DomainChan
}

func NewSearchServer(apis []ISearch, chars []rune, charNum int) *SearchServer {
	ss := &SearchServer{
		Apis:        apis,
		Chars:       chars,
		CharNum:     charNum,
		domainChan:  make(chan DomainChan, 1000),
		concurrency: RUN_STEP,
	}

	return ss
}

func (ss *SearchServer) Do() []Domain {
	ss.availables = ss.availables[0:0]
	combines := GetCombineMatch(ss.Chars, ss.CharNum)

	for i := 0; i < len(ss.Apis); i++ {
		api := ss.Apis[i]

		comblen := len(combines)
		for j := 0; j < comblen; {
			end := j + ss.concurrency
			if end > comblen {
				end = comblen
			}

			ss.bat(combines[j:end], api)
			j = j + ss.concurrency
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

func (ss *SearchServer) bat(combines []string, api ISearch) {
	wg := new(sync.WaitGroup)
	for _, v := range combines {
		go func(v string) {
			wg.Add(1)
			resp, err := http.Get(api.GetUrl(v + ".xyz"))
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()
			body, err2 := ioutil.ReadAll(resp.Body)
			if err2 != nil {
				panic(err2)
			}

			availables := api.ParseBody(body)
			ss.availables = append(ss.availables, availables...)
			log.Printf("bat %v", availables)
			wg.Done()
		}(v)
	}

	wg.Wait()
}
