package encode

import (
	"bytes"
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func GbkToUtf8(s []byte) []byte{
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	bRead, err := ioutil.ReadAll(reader)
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}
	return bRead
}

func Utf8ToGbk(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	bRead, err := ioutil.ReadAll(reader)
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}
	return bRead
}

func GbkToUtf84Str(strGbk string) string {
	bGbk := []byte(strGbk)
	bUtf8 := GbkToUtf8(bGbk)
	return string(bUtf8)
}

func Utf8ToGbk4Str(strUtf8 string) string {
	bUtf8 := []byte(strUtf8)
	bGbk := Utf8ToGbk(bUtf8)
	return string(bGbk)
}

func Base64(str string, isStr2Base64 bool) string {
	var result string
	if isStr2Base64 {
		input := []byte(str)
		result = base64.StdEncoding.EncodeToString(input)
	}else{
		decodeBytes, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			logrus.Error(err)
		}
		result = string(decodeBytes)
	}
	return result
}
