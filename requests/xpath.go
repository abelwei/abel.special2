package requests

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/xmlpath.v2"
	"strings"
)

type Xpath struct {
	Node *xmlpath.Node
	Err  error
	Path *xmlpath.Path
}

func (self *Xpath) New(sHtml string) {
	rdHtml := strings.NewReader(sHtml)
	self.Node, self.Err = xmlpath.ParseHTML(rdHtml)
}

func (self *Xpath) Parse2Str(sXpath string) string {
	var sResult string
	if pathComp, errComp := xmlpath.Compile(sXpath); errComp == nil {
		iter := pathComp.Iter(self.Node)
		//beego.Debug(pathComp.String(self.Node))
		if iter.Next() {
			sResult = iter.Node().String()
		}
	} else {
		logrus.Error(errComp)
	}
	return sResult
}

func (self *Xpath) Parse2Sli(sXpath string) []string {
	var sliResult []string
	pathComp := xmlpath.MustCompile(sXpath)
	iter := pathComp.Iter(self.Node)
	for iter.Next() {
		sliResult = append(sliResult, iter.Node().String())
	}
	return sliResult
}
