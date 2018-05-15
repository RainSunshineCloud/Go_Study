package main

import (
  "fmt"
)
var count = 0;
func test (s string) {
   for i := 0;i < 4;i++ {
     fmt.Println(s)
     count++;
  }
}

func main () {
  fmt.Println("php");
  go test("js");
  
  for  {
     if (count > 3){
	break;
     }
  }

  fmt.Println("php");
}
