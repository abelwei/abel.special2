package requests

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"strings"
)

type Jquery struct {
	Dom *goquery.Document
	Err error
}

func (self *Jquery) New(sHtml string) {
	var err error
	if self.Dom, err = goquery.NewDocumentFromReader(strings.NewReader(sHtml)); err != nil {
		logrus.Error(err)
	}

}

func (self *Jquery) Parse2Str(sRule string, bHtml bool) string {
	var sResult string
	var err error
	self.Dom.Find(sRule).EachWithBreak(func(i int, selection *goquery.Selection) bool {
		if bHtml {
			if sResult, err = selection.Html(); err != nil {
				logrus.Error(err)
			}
		} else {
			sResult = selection.Text()
		}
		if i == 0 {
			return true
		} else {
			return false
		}

	})
	return sResult
}

func (self *Jquery) Parse2Sli(sRule string, bHtml bool) []string {
	var sliResult []string
	self.Dom.Find(sRule).Each(func(i int, selection *goquery.Selection) {
		if bHtml {
			if sResult, err := selection.Html(); err == nil {
				sliResult = append(sliResult, sResult)
			} else {
				logrus.Error(err)
			}
		} else {
			sResult := selection.Text()
			sliResult = append(sliResult, sResult)
		}

	})
	return sliResult
}
