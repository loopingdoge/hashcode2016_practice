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

func paintRow(matrix [][] int, operations []Operation, spr int, spc int, ROWS int, COLS int) ([][]int, []Operation) {
  var op Operation
  rowLength := 0
  for i:=0; ((spc + i) < COLS) && (matrix[spr][spc + i] == 1); i++ {
    rowLength++
  }

  if rowLength == 1 {
    op = Operation {
      Name: "PAINT_SQUARE",
      Cells: [4] int {spr, spc, 0, 0},
    }
  } else {
    op = Operation {
      Name: "PAINT_LINE",
      Cells: [4] int {spr, spc, spr, spc + (rowLength-1)},
    }
  }
  operations = append(operations, op)
  for i := spc; i < (spc + rowLength); i++ {
    matrix[spr][i] = 0
  }
  return matrix, operations
}

func paintByLines(matrix [][]int, operations []Operation, ROWS int, COLS int) ([][]int, []Operation){
  for row,_ := range matrix {
    for col,_ := range matrix[row] {
      if matrix[row][col] == 1 {
        matrix, operations = paintRow(matrix, operations, row, col, ROWS, COLS)
      }
    }
  }
  return matrix, operations
}

func main() {
  operations := [] Operation {}
  inputFile := "inputs/right_angle.in"
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
  matrix, operations = paintByLines(matrix, operations, ROWS, COLS)
  fmt.Println(len(operations))
  for _,value := range operations {
    switch {
    case strings.Compare(value.Name, "PAINT_LINE") == 0:
      fmt.Printf("%s %d %d %d %d\n", value.Name, value.Cells[0], value.Cells[1], value.Cells[2], value.Cells[3])
    case strings.Compare(value.Name, "PAINT_SQUARE") == 0:
      fmt.Printf("%s %d %d %d\n", value.Name, value.Cells[0], value.Cells[1], value.Cells[2])
    case strings.Compare(value.Name, "ERASE_CELL") == 0:
        fmt.Printf("%s %d %d\n", value.Name, value.Cells[0], value.Cells[1])
    }
  }
  // for _,line := range matrix {
  //   fmt.Println(line)
  // }
}