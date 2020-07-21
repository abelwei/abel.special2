package requests

import (
	"compress/gzip"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

type Requests struct {
	hearders map[string]string
	JarCks   http.CookieJar //需要在init初始化设置
}


var (
	UnKnown = errors.New("unknown error")
	NetWork = errors.New("network error")
)

var JarCks http.CookieJar

func (self *Requests) IntiRequests() {
	self.JarCks, _ = cookiejar.New(nil)
	self.setDefaultMapHeader()
}


func (self *Requests) setDefaultMapHeader() {
	self.hearders = make(map[string]string)
	self.hearders["Connection"] = "keep-alive"
	self.hearders["Cache-Control"] = "max-age=0"
	self.hearders["Upgrade-Insecure-Requests"] = "1"
	self.hearders["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"
	self.hearders["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	self.hearders["Accept-Encoding"] = "gzip, deflate, br"
	self.hearders["Accept-Language"] = "zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7,fr;q=0.6,zh-TW;q=0.5"
}


func (self *Requests) SetHeader(sKey, sValue string) {
	self.hearders[sKey] = sValue
}

func (self *Requests) handle(method, url string, postVal url.Values) (strHtml string, err error) {
	var (
		postData *strings.Reader
		request	*http.Request
		errRequest error
	)

	if postVal == nil {
		request, errRequest = http.NewRequest(method, url, nil)
	}else{
		postData = strings.NewReader(postVal.Encode())
		request, errRequest = http.NewRequest(method, url, postData)
	}
	//logrus.Info(postVal.Encode()) //将Values转为aa=bb&cc=dd这样格式的字符串
	if errRequest == nil {
		//self.setDefaultMapHeader()
		client := &http.Client{Jar: self.JarCks}
		response, errDo := client.Do(request)
		//defer response.Body.Close()
		if errDo == nil {
			var sBody string
			if response.StatusCode == 200 {
				if response.Header.Get("Content-Encoding") == "gzip" {
					reader, _ := gzip.NewReader(response.Body)
					for {
						buf := make([]byte, 1024)
						n, errBuf := reader.Read(buf)
						if errBuf != nil && errBuf != io.EOF {
							//panic(errBuf)
							err = errBuf
							logrus.Error(errBuf)
							return
						}
						if n == 0 {
							break
						}
						sBody += string(buf)
					}
					strHtml = sBody
				} else {
					bBody, _ := ioutil.ReadAll(response.Body)
					strHtml = string(bBody)
				}
			} else {
				NetWork = errors.New(NetWork.Error()+" StatusCode: " + response.Status)
				err = NetWork
			}
		} else {
			err = errDo
			return
		}
		//defer response.Body.Close()
	} else {
		err = errRequest
		return
	}
	return
}

func (self *Requests) PostData(data map[string]string) (postVal url.Values) {
	postVal = url.Values{}
	for key, val := range data {
		postVal.Add(key, val)
	}
	return
}

func (self *Requests) Post(url string, data map[string]string) (strHtml string, err error) {
	return self.handle("POST", url, self.PostData(data))
}

func (self *Requests) Get(url string) (strHtml string, err error) {
	return self.handle("GET", url, nil)
}

func (self *Requests) Download(url string, savePath string) (err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	file, err := os.Create(savePath)
	if err != nil {
		return
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return
	}
	return
}