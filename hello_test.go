package main

import "testing"

func TestHello(t *testing.T){
  expected := "Hello World, from GoHorse!"
  actual := hello()
  if actual != expected {
    t.Error("Test failed, expected: '%s', got: '%s'", expected, actual)
  }
}
