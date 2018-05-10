package main

import (
    "fmt"
    "math/rand"
    "time"
    "strings"
)

var cnt int
var A string
var a, b, c bool

func main() {
  s := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","1","2","3","4","5","6","7","8","9","0"}

  fmt.Print("欲しいパスワードの文字数を入力してください : ")
  for ;; {
    fmt.Scan(&cnt)
    if 3 < cnt {
      break
    }

    fmt.Print("4桁以上で入力してください : ")
  }

  rand.Seed(time.Now().UnixNano())
  for ;; {

    i := 0
    for i < cnt {
      A = A + s[rand.Intn(62)]
      i++
    }

    i_a := 0
    for i_a < 25 {
      a = strings.Contains(A, s[i_a])
      i_a++
    }

    i_b := 26
    for i_b < 52 {
      b = strings.Contains(A, s[i_b])
      i_b++
    }

    i_c := 53
    for i_c < 62 {
      c = strings.Contains(A, s[i_c])
      i_c++
    }

    if a == true && b == true && c == true {
      break
    }

    a = false
    b = false
    c = false

    A = ""
  }

  fmt.Println(A)
}
