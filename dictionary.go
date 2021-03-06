package gwordsfilter

import (
	"strings"
)

type DictionaryList []DictionaryIndex

type DictionaryIndex struct {
	Key        string
	Index      int
	StartIndex int
}

type DictionaryItem struct {
	List []string
}

func NewDictionaryItem() *DictionaryItem {
	d := new(DictionaryItem)
	d.List = make([]string, 0)
	return d
}

func (t *DictionaryItem) Append(content string) {
	for _, v := range t.List {
		if v == content {
			return
		}
	}
	t.List = append(t.List, content)
}

func (t *DictionaryItem) CheckWord(words []rune, startIdx int) (bool, []DictionaryIndex) {
	for _, v := range t.List {
		indexList := make(DictionaryList, 0)
		checkWords := []rune(v)
		if len(checkWords) > len(words) {
			continue
		}
		index := 0
		for idx, word := range words {
			checkwordItem := checkWords[index]
			if strings.Index(string(checkwordItem), string(word)) >= 0 {
				indexList = append(indexList, DictionaryIndex{
					Key:        string(checkwordItem),
					Index:      idx,
					StartIndex: startIdx,
				})
				index++
				if len(indexList) >= len(checkWords) {
					return false, indexList
				}
			}
		}
	}
	return true, nil
}