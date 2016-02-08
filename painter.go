package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  // "strconv"
)

type Matrix [][] int

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  inputFile := "inputs/logo.in"
  // Read the file and check for errors
  dat, err := ioutil.ReadFile(inputFile)
  checkError(err)
  fileString := string(dat)
  // Lines array
  lines := strings.Split(fileString, "\n")
  // specs := lines[0]
  // rows_cols := strings.Split(specs, " ")
  // ROWS, err := strconv.Atoi(rows_cols[0])
  // COLS, err := strconv.Atoi(rows_cols[1])
  // checkError(err)
  lines = lines[1:]
  matrix := [][] int {}
  for _,line := range lines {
    lineToAppend := [] int {}
    for _, el := range line {
      var a int = 0
      if strings.Compare(string(el), ".") == 0 {
        a = 0
      } else if strings.Compare(string(el), "#") == 0 {
        a = 1
      } else {
        panic("unexpected char")
      }
      lineToAppend = append(lineToAppend, a)
    }
    matrix = append(matrix, lineToAppend)
  }
  for _,line := range matrix {
    fmt.Println(line)
  }
  // // Print how many legs have each animal
  // for _,animal := range animals {
  //   fmt.Printf("%s has %d legs\n", animal.Name, animal.Legs)
  // }
}