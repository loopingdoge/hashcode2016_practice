package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "math"
)

type Operation struct {
  Name string
  Cells [4] int
}

type Cells struct {
  Matrix [][]int
  Rows int
  Cols int
}

type Point struct {
  X int
  Y int
}

func checkError(e error) {
  if e != nil {
    panic(e)
  }
}

func checkSquare(cells Cells, p Point, length int) bool {
  squarable := true
  if length % 2 == 0 {
    panic("you can use only an odd length")
  }
  for i := -(length/2); i <= length/2; i++ {
    for j := -(length/2); j <= length/2; j++ {
      if (p.X+i < cells.Rows) && (p.X+i >= 0) && (p.Y+j < cells.Cols) && (p.Y+j >= 0) {
        if cells.Matrix[p.X+i][p.Y+j] != 1 {
          squarable = false
          return squarable
        }
      } else {
        squarable = false
        return squarable
      }
    }
  }
  return squarable
}

func paintSquare(cells Cells, operations []Operation, p Point, size int) (Cells, []Operation) {
  op := Operation {
    Name: "PAINT_SQUARE",
    Cells: [4] int {p.X, p.Y, size/2, 0},
  }
  for i := -(size/2); i <= size/2; i++ {
    for j := -(size/2); j <= size/2; j++ {
      if (p.X+i < cells.Rows) && (p.X+i >= 0) && (p.Y+j < cells.Cols) && (p.Y+j >= 0) {
        cells.Matrix[p.X+i][p.Y+j] = 0
      }
    }
  }
  operations = append(operations, op)
  return cells, operations
}

func rowLength(cells Cells, p Point) int {
  var i int
  for i=0; (p.Y+i < cells.Cols) && (cells.Matrix[p.X][p.Y+i] == 1); i++ {}
  return i
}

func colLength(cells Cells, p Point) int {
  var i int
  for i=0; (p.X+i < cells.Rows) && (cells.Matrix[p.X+i][p.Y] == 1); i++ {}
  return i
}

func paintRow(cells Cells, operations []Operation, p Point) (Cells, []Operation) {
  var op Operation
  rowLength := rowLength(cells, p)
  if rowLength == 1 {
    op = Operation {
      Name: "PAINT_SQUARE",
      Cells: [4] int {p.X, p.Y, 0, 0},
    }
  } else {
    op = Operation {
      Name: "PAINT_LINE",
      Cells: [4] int {p.X, p.Y, p.X, p.Y + (rowLength-1)},
    }
  }
  operations = append(operations, op)
  for i := p.Y; i < (p.Y + rowLength); i++ {
    cells.Matrix[p.X][i] = 0
  }
  return cells, operations
}

func paintCol(cells Cells, operations []Operation, p Point) (Cells, []Operation) {
  var op Operation
  colLength := colLength(cells, p)
  if colLength == 1 {
    op = Operation {
      Name: "PAINT_SQUARE",
      Cells: [4] int {p.X, p.Y, 0, 0},
    }
  } else {
    op = Operation {
      Name: "PAINT_LINE",
      Cells: [4] int {p.X, p.Y, p.X + (colLength-1), p.Y},
    }
  }
  operations = append(operations, op)
  for i := p.X; i < (p.X + colLength); i++ {
    cells.Matrix[i][p.Y] = 0
  }
  return cells, operations
}

func paintByLines(cells Cells, operations []Operation) (Cells, []Operation) {
  for row,_ := range cells.Matrix {
    for col,_ := range cells.Matrix[row] {
      if cells.Matrix[row][col] == 1 {
        cells, operations = paintRow(cells, operations, Point{X: row, Y: col})
      }
    }
  }
  return cells, operations
}

func paintBySquares(cells Cells, operations []Operation) (Cells, []Operation) {
  for row,_ := range cells.Matrix {
    for col,_ := range cells.Matrix[row] {
      if cells.Matrix[row][col] == 1 {
        numIter := -1
        squareSize := 3
        canPaint := false
        for i:=0; checkSquare(cells, Point{X:row+i, Y:col+i}, squareSize); i += 2 {
          canPaint = true
          squareSize += i
          numIter++
        }
        if canPaint {
          cells, operations = paintSquare(cells, operations, Point{X: row+numIter, Y: col+numIter}, squareSize)
        }
      }
    }
  }
  return cells, operations
}

func paintByRows(cells Cells, operations []Operation, minLength int) (Cells, []Operation) {
  for row,_ := range cells.Matrix {
    for col,_ := range cells.Matrix[row] {
      if cells.Matrix[row][col] == 1 {
        p := Point{X: row, Y: col}
        if rowLength(cells, p) > minLength {
          cells, operations = paintRow(cells, operations, p)
        }
      }
    }
  }
  return cells, operations
}

func paintByCols(cells Cells, operations []Operation, minLength int) (Cells, []Operation) {
    for col,_ := range cells.Matrix[0] {
      for row,_ := range cells.Matrix {
        if cells.Matrix[row][col] == 1 {
          p := Point{X: row, Y: col}
          if colLength(cells, p) > minLength {
            cells, operations = paintCol(cells, operations, p)
          }
        }
      }
    }
  return cells, operations
}

func printOutput(operations []Operation) {
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
  cells := Cells {
    Matrix: matrix,
    Rows: ROWS,
    Cols: COLS,
  }
  
  //cells, operations = paintBySquares(cells, operations)

  for minLength := int(math.Max(float64(cells.Rows), float64(cells.Cols)) - 1); minLength >= 0; minLength-- {
    cells, operations = paintByRows(cells, operations, minLength)
    cells, operations = paintByCols(cells, operations, minLength)
  }

  fmt.Println(len(operations))
  printOutput(operations)
  // for _,line := range matrix {
  //   fmt.Println(line)
  // }
}