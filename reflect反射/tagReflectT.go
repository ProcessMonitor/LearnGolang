package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	name string `info:"name" doc:"name" xml:"name.xml"`
	sex  string `info:"sex" doc:"sex" xml:"sex.xml"`
}

func findTag(ii interface{}) {
	t := reflect.TypeOf(ii).Elem()
	for i := 0; i < t.NumField(); i++ {
		cur_tag_xml := t.Field(i).Tag.Get("xml")
		cur_tag_info := t.Field(i).Tag.Get("info")
		cur_tag_doc := t.Field(i).Tag.Get("doc")
		fmt.Println("tagxml :", cur_tag_xml, "info :", cur_tag_info, "doc:", cur_tag_doc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
