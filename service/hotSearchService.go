package service

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HotSearchService struct {
	DataType string
}

func (hotSearchService *HotSearchService) GetV2EX() []map[string]interface{} {
	url := "https://www.v2ex.com/?tab=hot"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".item_title").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.v2ex.com" + url})
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetITHome() []map[string]interface{} {
	url := "https://www.ithome.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".hot-list .bx ul li").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": url})
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetZhiHu() []map[string]interface{} {
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "https://www.zhihu.com/hot"
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("Cookie", `_zap=09ee8132-fd2b-43d3-9562-9d53a41a4ef5; d_c0="AGDv-acVoQ-PTvS01pG8OiR9v_9niR11ukg=|1561288241"; capsion_ticket="2|1:0|10:1561288248|14:capsion_ticket|44:NjE1ZTMxMjcxYjlhNGJkMjk5OGU4NTRlNDdkZTJhNzk=|7aefc35b3dfd27b74a087dd1d15e7a6bb9bf5c6cdbe8471bc20008feb67e7a9f"; z_c0="2|1:0|10:1561288250|4:z_c0|92:Mi4xeGZsekFBQUFBQUFBWU9fNXB4V2hEeVlBQUFCZ0FsVk5PcXo4WFFBNWFFRnhYX2h0ZFZpWTQ5T3dDMGh5ZTV1bjB3|0cee5ae41ff7053a1e39d96df2450077d37cc9924b337584cf006028b0a02f30"; q_c1=ae65e92b2bbf49e58dee5b2b29e1ffb3|1561288383000|1561288383000; tgw_l7_route=f2979fdd289e2265b2f12e4f4a478330; _xsrf=f8139fd6-b026-4f01-b860-fe219aa63543; tst=h; tshl=`)
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".HotList-list .HotItem-content").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("h2").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": url})
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetWeiBo() []map[string]interface{} {
	url := "https://s.weibo.com/top/summary"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".list_a li").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a span").Text()
		textLock := selection.Find("a em").Text()
		text = strings.Replace(text, textLock, "", -1)
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://s.weibo.com" + url})
		}
	})
	return allData[0:]

}

// 贴吧
func (hotSearchService *HotSearchService) GetTieBa() []map[string]interface{} {
	url := "http://tieba.baidu.com/hottopic/browse/topicList"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	str, _ := ioutil.ReadAll(res.Body)
	js, err2 := simplejson.NewJson(str)
	if err2 != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	i := 1
	for i < 30 {
		test := js.Get("data").Get("bang_topic").Get("topic_list").GetIndex(i).MustMap()
		allData = append(allData, map[string]interface{}{"title": test["topic_name"], "url": test["topic_url"]})
		i++
	}
	return allData

}

// 豆瓣
func (hotSearchService *HotSearchService) GetDouBan() []map[string]interface{} {
	url := "https://www.douban.com/group/explore"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Referer", `https://www.douban.com/group/explore`)
	request.Header.Add("Host", `www.douban.com`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str,_ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".channel-item").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("h3 a").Attr("href")
		text := selection.Find("h3 a").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": url})
		}
	})
	return allData
}

// 天涯
func (hotSearchService *HotSearchService) GetTianYa() []map[string]interface{} {
	url := "http://bbs.tianya.cn/list.jsp?item=funinfo&grade=3&order=1"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Referer", `http://bbs.tianya.cn/list.jsp?item=funinfo&grade=3&order=1`)
	request.Header.Add("Host", `bbs.tianya.cn`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str,_ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find("table tr").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("td a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "http://bbs.tianya.cn/" + url})
		}
	})
	return allData
}

// 虎扑
func (hotSearchService *HotSearchService) GetHuPu() []map[string]interface{} {
	url := "https://bbs.hupu.com/all-gambia"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Referer", `https://bbs.hupu.com/`)
	request.Header.Add("Host", `bbs.hupu.com`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str,_ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".bbsHotPit li").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find(".textSpan a")
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://bbs.hupu.com/" + url})
		}
	})
	return allData
}

// Github
func (hotSearchService *HotSearchService) GetGitHub() []map[string]interface{} {
	url := "https://github.com/trending"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str,_ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}

	document.Find(".Box article").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find(".lh-condensed a")
		//desc := selection.Find(".col-9 .text-gray .my-1 .pr-4")
		//descText := desc.Text()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		descText := selection.Find("p").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "desc": descText, "url": "https://github.com" + url})
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetBaiDu() []map[string]interface{} {
	url := "http://top.baidu.com/buzz?b=341&c=513&fr=topbuzz_b1"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Host", `top.baidu.com`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find("table tr").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		MyText, _ := GbkToUtf8([]byte(text))
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": string(MyText), "url": url})
		}
	})
	return allData

}

func (hotSearchService *HotSearchService) Get36Kr() []map[string]interface{} {
	url := "https://36kr.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Host", `36kr.com`)
	request.Header.Add("Referer", `https://36kr.com/`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str,_ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".hotlist-item-toptwo").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := selection.Find("ap").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://36kr.com" + url})
		}
	})
	document.Find(".hotlist-item-other-info").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://36kr.com" + url})
		}
	})
	return allData

}

func (hotSearchService *HotSearchService) GetQDaily() []map[string]interface{} {
	url := "https://www.qdaily.com/tags/29.html"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Host", `www.qdaily.com`)
	request.Header.Add("Referer", `https://www.qdaily.com/tags/30.html`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str,_ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".packery-item").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := selection.Find(".grid-article-bd h3").Text()
		if len(text) != 0 {
			if boolUrl {
				allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.qdaily.com/" + url})
			}
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetGuoKr() []map[string]interface{} {
	url := "https://www.guokr.com/scientific/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Host", `www.guokr.com`)
	request.Header.Add("Referer", `https://www.guokr.com/scientific/`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str,_ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find("div .article").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("h3 a")
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
			}
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetHuXiu() []map[string]interface{} {
	url := "https://www.huxiu.com/article"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Host", `www.guokr.com`)
	request.Header.Add("Referer", `https://www.huxiu.com/channel/107.html`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str,_ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".article-item--large__content").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Find("h5").Text()
		if len(text) != 0 {
			if boolUrl {
				allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.huxiu.com" + url})
			}
		}
	})
	document.Find(".article-item__content").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Find("h5").Text()
		if len(text) != 0 {
			if boolUrl {
				allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.huxiu.com" + url})
			}
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetDBMovie() []map[string]interface{} {
	url := "https://movie.douban.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".slide-container").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a")
		url, boolUrl := s.Attr("href")
		text := s.Find("p").Text()
		if len(text) != 0 {
			if boolUrl {
				allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.huxiu.com" + url})
			}
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetZHDaily() []map[string]interface{} {
	url := "http://daily.zhihu.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".row .box").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Find("span").Text()
		if len(text) != 0 {
			if boolUrl {
				allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://daily.zhihu.com" + url})
			}
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetSegmentfault() []map[string]interface{} {
	url := "https://segmentfault.com/hottest"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".news-list .news__item-info").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a:nth-child(2)").First()
		url, boolUrl := s.Attr("href")
		text := s.Find("h4").Text()
		if len(text) != 0 {
			if boolUrl {
				allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://segmentfault.com" + url})
			}
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetHacPai() []map[string]interface{} {
	url := "https://hacpai.com/domain/play"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".hotkey li").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("h2 a")
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
			}
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetWYNews() []map[string]interface{} {
	url := "http://news.163.com/special/0001386F/rank_whole.html"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find("table tr").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("td a").First()
		url, boolUrl := s.Attr("href")
		text, _ := GbkToUtf8([]byte(s.Text()))
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
				}
			}
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetWaterAndWood() []map[string]interface{} {
	url := "https://www.newsmth.net/nForum/mainpage?ajax"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//sss,_ := GbkToUtf8([]byte(string(str)))
	//fmt.Println(string(sss))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	// topics
	document.Find("#top10 li").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a:nth-child(2)").First()
		url, boolUrl := s.Attr("href")
		text, _ := GbkToUtf8([]byte(s.Text()))
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.newsmth.net" + url})
				}
			}
		}
	})
	document.Find(".topics").Find("li").Each(func(i int, selection *goquery.Selection) {
		if i > 10 {
			s := selection.Find("a:nth-child(2)").First()
			url, boolUrl := s.Attr("href")
			text, _ := GbkToUtf8([]byte(s.Text()))
			if len(text) != 0 {
				if boolUrl {
					if len(allData) <= 100 {
						allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.newsmth.net" + url})
					}
				}
			}
		}
	})
	return allData
}

// http://nga.cn/

func (hotSearchService *HotSearchService) GetNGA() []map[string]interface{} {
	url := "http://nga.cn/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find("h2").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
				}
			}
		}
	})
	return allData
}

func (hotSearchService *HotSearchService) GetCSDN() []map[string]interface{} {
	url := "https://www.csdn.net/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find("#feedlist_id li").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("h2 a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
				}
			}
		}
	})
	return allData
}

// https://weixin.sogou.com/?pid=sogou-wsse-721e049e9903c3a7&kw=
func (hotSearchService *HotSearchService) GetWeiXin() []map[string]interface{} {
	url := "https://weixin.sogou.com/?pid=sogou-wsse-721e049e9903c3a7&kw="
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".news-list li").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("h3 a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
				}
			}
		}
	})
	return allData
}

//

func (hotSearchService *HotSearchService) GetKD() []map[string]interface{} {
	url := "http://www.kdnet.net/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".indexside-box-hot li").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text, _ := GbkToUtf8([]byte(s.Text()))
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
				}
			}
		}
	})
	return allData
}

// http://www.mop.com/

func (hotSearchService *HotSearchService) GetMop() []map[string]interface{} {
	url := "http://www.mop.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".swiper-slide").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := selection.Find("h2").Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
				}
			}
		}
	})
	document.Find(".tabel-right").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := selection.Find("h3").Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
				}
			}
		}
	})
	return allData[:15]
}

// https://www.chiphell.com/

func (hotSearchService *HotSearchService) GetChiphell() []map[string]interface{} {
	url := "https://www.chiphell.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find("#frameZ3L5I7 li").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.chiphell.com/" + url})
				}
			}
		}
	})
	// portal_block_530_content
	document.Find("#portal_block_530_content dt").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.chiphell.com/" + url})
				}
			}
		}
	})
	// frame-tab move-span cl
	document.Find("#portal_block_560_content dt").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.chiphell.com/" + url})
				}
			}
		}
	})
	// portal_block_564_content
	document.Find("#portal_block_564_content dt").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.chiphell.com/" + url})
				}
			}
		}
	})
	// portal_block_568_content
	document.Find("#portal_block_568_content dt").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.chiphell.com/" + url})
				}
			}
		}
	})
	// portal_block_569_content
	document.Find("#portal_block_569_content dt").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.chiphell.com/" + url})
				}
			}
		}
	})
	// portal_block_570_content
	document.Find("#portal_block_570_content dt").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://www.chiphell.com/" + url})
				}
			}
		}
	})
	return allData
}

// http://jandan.net/

func (hotSearchService *HotSearchService) GetJianDan() []map[string]interface{} {
	url := "http://jandan.net/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find("h2").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if len(text) != 0 {
			if boolUrl {
				if len(allData) <= 100 {
					allData = append(allData, map[string]interface{}{"title": string(text), "url": url})
				}
			}
		}
	})
	return allData
}

// https://dig.chouti.com/

func (hotSearchService *HotSearchService) GetChouTi() []map[string]interface{} {
	url := "https://dig.chouti.com/top/24hr?_=" + strconv.FormatInt(time.Now().Unix(), 10) + "163"
	url2 := "https://dig.chouti.com/link/hot?afterTime=" + strconv.FormatInt(time.Now().Unix(), 10) + "026000" + "&_=" + strconv.FormatInt(time.Now().Unix(), 10) + "667"
	res, err := http.Get(url)
	res2, _ := http.Get(url2)
	if err != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	str, _ := ioutil.ReadAll(res.Body)
	str2, _ := ioutil.ReadAll(res2.Body)
	js, err2 := simplejson.NewJson(str)
	js2, _ := simplejson.NewJson(str2)
	if err2 != nil {
		fmt.Println("抓取" + hotSearchService.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	i := 1
	for i < 30 {
		test := js.Get("data").GetIndex(i).MustMap()
		if test["title"] != nil && test["url"] != nil {
			allData = append(allData, map[string]interface{}{"title": test["title"], "url": test["url"]})
		}
		i++
	}
	j := 1
	for j < 60 {
		test := js2.Get("data").GetIndex(j).MustMap()
		if test["title"] != nil && test["url"] != nil {
			allData = append(allData, map[string]interface{}{"title": test["title"], "url": test["url"]})
		}
		j++
	}
	return allData

}

/*
*
部分热榜标题需要转码
*/
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
