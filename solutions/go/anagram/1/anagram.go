package anagram

import (
  //  "fmt"
    "strings"
    "sort"
    )

func sortLetters(word string) string{
    letters:= strings.Split(strings.ToLower(word),"")
    sort.Strings(letters)
    return strings.Join(letters,"")
}

func Detect(subject string, candidates []string) []string {
	//fmt.Println(subject)
    //fmt.Println(candidates)

    var result []string
    subjectSorted:= sortLetters(subject)
    
    for _,word:=range candidates{
        if strings.ToLower(subject)==strings.ToLower(word){
            continue
        }
        if sortLetters(word)==subjectSorted{
            result=append(result, word)
        }
      
    }
	  return result
}
