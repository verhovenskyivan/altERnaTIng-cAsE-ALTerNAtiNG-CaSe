package kata
package kata

import "unicode"

func ToAlternatingCase(str string) string {
  result := []rune{}
  for _, ch:=range str{
    if unicode.IsUpper(ch){
      result = append(result, unicode.ToLower(ch))
    } else if unicode.IsLower(ch){
       result = append(result, unicode.ToUpper(ch))
    } else {
      result = append(result, ch)
    }
  }
  
  return string(result)
}