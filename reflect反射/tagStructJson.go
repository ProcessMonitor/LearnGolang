package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Actors []string `json:"actors"`
}

func main() {
	m := Movie{"重庆森林", 1989, []string{"金城武", "梁朝伟"}}
	//fmt.Println("movie", m)

	jsonstr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	//fmt.Println("json_str :", jsonstr) //会失败
	fmt.Println("json_str :", string(jsonstr))

	// jsonstr : {"title":"重庆森林","year":1989,"actors":["金城武","梁朝伟"]}
	new_m := Movie{}
	err = json.Unmarshal(jsonstr, &new_m)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Print("new_m:", new_m)
}
