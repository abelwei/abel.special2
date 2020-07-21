package datastruct

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func Yml2Map(text string) map[interface{}]interface{} {
	mapYml := make(map[interface{}]interface{})
	bText := []byte(text)
	err := yaml.Unmarshal(bText,&mapYml)
	if err != nil {
		logrus.Error(err)
	}
	return mapYml
}


func Map2Yml(mpaz map[interface{}]interface{}) string {
	//txtYml := ""
	txtYml, err := yaml.Marshal(&mpaz)
	if(err == nil){
		return string(txtYml)
	}else{
		logrus.Error(err)
		return ""
	}

}

func MapMerge(source, add map[interface{}]interface{}) map[interface{}]interface{} {
	for k, v := range add {
		source[k] = v
	}
	return source
}
