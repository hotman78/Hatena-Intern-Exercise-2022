package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var res=make(map[string]int)
	for _,str:= range strings.Fields(s){
		if _,ok:=res[str];!ok{
			res[str]=0
		}
		res[str]++
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
