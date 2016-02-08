package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

type Operation struct {
  Name string
  Cells [4] int
}

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func paintRow(matrix *[][] int, operations *[]Operation, spr int, spc int, ROWS int, COLS int) int {
  var op Operation
  rowLength := 0
  for i:=0; ((spc + i) < COLS) && ((*matrix)[spr][spc + i] == 1); i++ {
    rowLength++
  }

  if rowLength == 0 {
    op = Operation {
      Name: "PAINT_SQUARE",
      Cells: [4] int {spr, spc, 0, 0},
    }
  } else {
    op = Operation {
      Name: "PAINT_LINE",
      Cells: [4] int {spr, spc, spr, spc + rowLength},
    }
  }
  (*operations) = append(*operations, op)
  for i:=spc; i <= (spc + rowLength); i++ {
    (*matrix)[spr][spc] = 0
  }
  return rowLength
}

func paintByLines(matrix *[][]int, operations *[]Operation, ROWS int, COLS int) {
  for row,_ := range (*matrix) {
    for col,_ := range (*matrix)[row] {
      if (*matrix)[row][col] == 1 {
        paintRow(matrix, operations, row, col, ROWS, COLS)
      }
    }
  }
}

func main() {
  operations := [] Operation {}
  inputFile := "inputs/logo.in"
  // Read the file and check for errors
  dat, err := ioutil.ReadFile(inputFile)
  checkError(err)
  fileString := string(dat)
  // Lines array
  lines := strings.Split(fileString, "\n")
  specs := lines[0]
  rows_cols := strings.Split(specs, " ")
  ROWS, err := strconv.Atoi(rows_cols[0])
  checkError(err)
  COLS, err := strconv.Atoi(rows_cols[1])
  checkError(err)
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
  paintByLines(&matrix, &operations, ROWS, COLS)
  fmt.Println(len(operations))
  i := 0
  for _,value := range operations {
    fmt.Println(value)
    i++
  }
  fmt.Println(i)
  // for _,line := range matrix {
  //   fmt.Println(line)
  // }
}