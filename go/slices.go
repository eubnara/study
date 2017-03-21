package main

import "fmt"

func main() {
    s := make([]int, 3)


    s[0] = 1;
    s[1] = 2;
    s[2] = 3;

   fmt.Println(s)
   fmt.Println(s[2])

   s = append(s, 4)
   fmt.Println(s)

   c := make([]int, len(s))
   copy(c, s)
   fmt.Println("cpy: ", c)

   l := s[2:5]
   fmt.Println("sl1:", l)

   // t := make([]string, 3)
   // t[0] = "g"
   // t[1] = "h"
   // t[2] = "i"
   t := []string{"g", "h", "i"}
   fmt.Println("dcl:", t)

   twoD := make([][]int, 3)
   for i := 0; i<3; i++ {
       innerLen := i + 1
       twoD[i] = make([]int, innerLen)
       for j := 0; j < innerLen; j++ {
           twoD[i][j] = i + j
       }
   }
   fmt.Println("2d: ", twoD)
}
