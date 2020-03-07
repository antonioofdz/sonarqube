package main

import (
  "fmt"
)

func main() {
  pesos := []int{1,2,3,1,2,3,1,2,3,1,3,4,4,2}

  mapPeso := make(map[int]int)
  for _, peso := range pesos {
   mapPeso[peso] = countElements(pesos, peso)
  }

  var solutions int
  for _, occ := range mapPeso {
    solutions += getSolutions(occ)
  }

  fmt.Println("SOLUCION:  ", solutions)
}

func getSolutions(n int) int {
  even := n%2==0
  if !even {
    n = n - 1
  }

  if n == 2 {
    return 1;
  }

  return n/2;
}

func countElements(list []int, value int) int {
  var total int
  for _, item := range list {
    if item == value {
      total++
    }
  }
  return total
}
