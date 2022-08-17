package main

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
)

//Mydataを設定
type Mydata struct{
	Name string
	Mail string
	Tel string
}
func (m *Mydata) Str() string{
	return "<\"" + m.Name +"\" " + m.Mail + ","+ m.Tel + ">"
}

func main(){
	p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"
	re,er := http.Get(p)
	if er != nil{
		panic(er)
	}
	defer re.Body.Close()

	s, er := ioutil.ReadAll(re.Body)
	if er != nil{
		panic(er)
	}

	var items []Mydata
	er = json.Unmarshal(s,&items)//取り出す
	if er != nil{
		panic(er)
	}

	for i,im := range items{
		println(i,im.Str())
	}
	/*
	for i, im := range data{
		m := im.(map[string]interface{})
		fmt.Println(i,m["name"].(string),m["mail"].(string),m["tel"].(string))
	}
	*/
}