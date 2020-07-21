package requests

import (
	"github.com/antchfx/htmlquery"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"strings"
)

type Xhtml struct {
	Dom *html.Node
}

func (self *Xhtml) New(sHtml string) {
	var err error
	if self.Dom, err = htmlquery.Parse(strings.NewReader(sHtml)); err != nil {
		logrus.Debug(err)
	}
}

func (self *Xhtml) Parse2Str(sRule string, bHtml bool) string {
	var sResult string
	dom := htmlquery.FindOne(self.Dom, sRule)
	if dom != nil {
		if bHtml {
			sResult = htmlquery.OutputHTML(dom, true)
		} else {
			dom := htmlquery.FindOne(self.Dom, sRule)
			sResult = htmlquery.InnerText(dom)
		}
	} else {
		logrus.Debug("返回的数据空，XPATH:" + sRule)
	}

	return sResult
}

func (self *Xhtml) Parse2Sli(sRule string, bHtml bool) []string {
	var sliResult []string
	lsDom := htmlquery.Find(self.Dom, sRule)
	if lsDom != nil {
		if bHtml {
			for _, frDom := range lsDom {
				sliResult = append(sliResult, htmlquery.OutputHTML(frDom, true))
			}
		} else {
			for _, frDom := range lsDom {
				sliResult = append(sliResult, htmlquery.InnerText(frDom))
			}
		}
	} else {
		logrus.Debug("返回的数据空，XPATH:" + sRule)
	}
	return sliResult
}

type Url struct {
	Link string
	Title string
}

func (self *Xhtml) Parse2SliUrl(sRule string) []Url {
	var reslut []Url
	sliUrl := self.Parse2Sli(sRule+"/@href", false)
	sliTitle := self.Parse2Sli(sRule+"/text()", false)
	urlCount := len(sliUrl)
	titleCount := len(sliTitle)
	if urlCount == titleCount {
		for i:=0; i < urlCount; i++ {
			url := Url{}
			url.Link = sliUrl[i]
			url.Title = sliTitle[i]
			reslut = append(reslut, url)
		}
	}else{
		logrus.Fatal("url and title unequal")
	}
	return reslut
}

func (self *Xhtml) Parse2Url(sRule string) Url {
	result := Url{}
	result.Link = self.Parse2Str(sRule+"/@href", false)
	result.Title = self.Parse2Str(sRule+"/text()", false)
	return result
}

