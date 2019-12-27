package wordsfilter

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Wordsfilter struct {
	Dictionary map[string]*DictionaryItem
	ReplaceStr string
	/** 大于这个值 直接返回空 */
	LimitReplaceStr int
}

func New() *Wordsfilter {
	wf := new(Wordsfilter)
	wf.Dictionary = make(map[string]*DictionaryItem)
	wf.ReplaceStr = "*"
	return wf
}

func (t *Wordsfilter) Read(filename string) (err error){
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer func() {
		_ = file.Close()
	}()
	line := bufio.NewReader(file)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(content)
		if strings.Trim(str, " ") == ""{
			continue
		}
		for _, v := range str {
			if _, ok := t.Dictionary[string(v)]; ok == false {
				t.Dictionary[string(v)] = NewDictionaryItem()
			}
			t.Dictionary[string(v)].Append(str)
			break
		}
	}
	return
}

func (t *Wordsfilter) CheckWord(word string) (bool, []DictionaryIndex) {
	content := []rune(word)
	max := len([]rune(word))
	if max == 0 {
		return true, nil
	}
	indexList := make([]DictionaryIndex, 0)
	for i := 0; i < max; i++ {
		val := strings.Trim(string(content[i]), " ")
		if len(val) > 0 {
			if item, ok := t.Dictionary[string(val)]; ok {
				str := content[i:]
				flag, idxList := item.CheckWord(str, i)
				if !flag {
					indexList = append(indexList, idxList...)
					if t.LimitReplaceStr > 0 && len(indexList) < t.LimitReplaceStr{
						return false, indexList
					}
				}
			} else {
				continue
			}
		}
	}
	if len(indexList) > 0 {
		return false, indexList
	}
	return true, nil
}

func (t *Wordsfilter) Replace(word string) string {
	flag, idxList := t.CheckWord(word)
	if !flag {
		if t.LimitReplaceStr > 0 && t.LimitReplaceStr >= len(idxList){
			return ""
		}
		for _, v := range idxList {
			word = strings.Replace(word, string([]rune(word)[v.StartIndex+v.Index]), t.ReplaceStr, 1)
		}
	}
	return word
}