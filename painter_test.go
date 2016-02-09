package main

import "testing"

func TestCheckSquare(t *testing.T) {
  var actual bool
  expected := true
  m := [][]int {
    {1,1,1},
    {1,1,1},
    {1,1,1},
  }
  p := Point{X:1, Y:1}
  c := Cells {
    Matrix: m,
    Rows: len(m),
    Cols: len(m[0]),
  }
  actual = checkSquare(c, p, 3)
  if actual != expected {
    t.Errorf("expected %v to be a square", m)
  }
  // --------------------------------------
  m2 := [][]int {
    {0, 0, 0, 0, 0},
    {0, 0, 1, 1, 0},
    {0, 1, 1, 1, 0},
    {0, 1, 1, 1, 0},
    {0, 0, 0, 0, 0},
  }
  p = Point{X:1, Y:1}
  c = Cells {
    Matrix: m2,
    Rows: len(m2),
    Cols: len(m2[0]),
  }
  expected = false
  actual = checkSquare(c, p, 3)
  if actual != expected {
    t.Errorf("expected %v to not be a square", m)
  }
}

func TestRowLength(t *testing.T) {
  var actual int
  expected := 5
  r := [][]int {
    {0, 1, 1, 1, 1, 1, 0},
  }
  c := Cells {
    Matrix: r,
    Rows: len(r),
    Cols: len(r[0]),
  }
  p := Point{X: 0, Y: 1}
  actual = rowLength(c, p)
  if actual != expected {
    t.Errorf("expected %v to be a long %d, got %d", r, expected, actual)
  }
}

func TestColLength(t *testing.T) {
  var actual int
  expected := 3
  r := [][]int {
    {0, 1},
    {0, 1},
    {0, 1},
    {1, 0},
    {0, 1},
  }
  c := Cells {
    Matrix: r,
    Rows: len(r),
    Cols: len(r[0]),
  }
  p := Point{X: 0, Y: 1}
  actual = colLength(c, p)
  if actual != expected {
    t.Errorf("expected %v to be a long %d, got %d", r, expected, actual)
  }
}

func TestPaintSquare(t *testing.T) {
  actualOperations := [] Operation {}
  m := [][]int {
    {0, 0, 0, 0, 0},
    {0, 1, 1, 1, 0},
    {0, 1, 1, 1, 0},
    {0, 1, 1, 1, 0},
    {0, 0, 0, 0, 0},
  }
  actualCells := Cells {
    Matrix: m,
    Rows: len(m),
    Cols: len(m[0]),
  }
  p := Point{X: 2, Y: 2}
  actualCells, actualOperations = paintSquare(actualCells, actualOperations, p, 3)
  for _,row := range m {
    for _,value := range row {
      if value != 0 {
        t.Errorf("expected the matrix %v to contain only zeroes", m)
      }
    }
  }
}