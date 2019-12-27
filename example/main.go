package main

import (
	"fmt"
	"time"
	"gwordsfilter"
)

func main(){
	start := time.Now()
	wf := gwordsfilter.New()
	fmt.Println(wf.Read("./wordsfilter.txt"))
	fmt.Println(wf.Replace("测试测试游$$$戏&&&&开发^^^^人员测试测试"))
	fmt.Println(time.Now().Sub(start))
}