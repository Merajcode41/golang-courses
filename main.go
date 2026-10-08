package main

import(
  "fmt"
)


// we hav two way for decleration variable
// firs way Global scop
var name string

func main() {
  // using decler variable and initial variable
  name = "Ali"
  // Println a build-tin function on standard library in golang for print message
  fmt.Println(name)

  // second way decleration variable an initialation
  // this variable is local scop
  family := "Rostami"

  fmt.Println(family)
}