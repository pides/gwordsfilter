package main

import (
	"fmt"
	"time"
	"wordsfilter"
)

func main(){
	start := time.Now()
	wf := wordsfilter.New()
	wf.Read("./wordsfilter.txt")
	fmt.Println(wf.Replace("啦啦啦啦啦啦"))
	fmt.Println(time.Now().Sub(start))
}