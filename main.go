package main

import "fmt"

func main() {
  var res string
  res = hello()
  fmt.Println(res)
}

func hello() string {
    return "Hello World, from GoHorse!";
}
