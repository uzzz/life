package main

import "strings"

type Field struct {
    cells [][]*Cell
}

func NewField(width, height int) *Field {
    field := new(Field)
    field.cells = make([][]*Cell, height)
    for i := 0; i < height; i++ {
        field.cells[i] = make([]*Cell, width)
        for j := 0; j < width; j++ {
            field.cells[i][j] = new(Cell)
        }
    }
    return field
}

func (f *Field) CellAt(col, row int) *Cell {
    if row < 0 || col < 0 || row > f.Height() - 1 || col > f.Width() - 1 {
        return nil
    }

    return f.cells[row][col]
}

func (f *Field) AliveCellsAround(col, row int) int {
    alive := 0

    for _, neighbor := range f.neighborsOf(col, row) {
        if neighbor.IsAlive() {
            alive += 1
        }
    }

    return alive
}

func (f *Field) neighborsOf(col, row int) []*Cell {
    neighbors := make([]*Cell, 0)

    if f.CellAt(col - 1, row) != nil {
        neighbors = append(neighbors, f.CellAt(col - 1, row))
    }

    if f.CellAt(col - 1, row - 1) != nil {
        neighbors = append(neighbors, f.CellAt(col - 1, row - 1))
    }

    if f.CellAt(col - 1, row + 1) != nil {
        neighbors = append(neighbors, f.CellAt(col - 1, row + 1))
    }

    if f.CellAt(col, row - 1) != nil {
        neighbors = append(neighbors, f.CellAt(col, row - 1))
    }

    if f.CellAt(col, row + 1) != nil {
        neighbors = append(neighbors, f.CellAt(col, row + 1))
    }

    if f.CellAt(col + 1, row) != nil {
        neighbors = append(neighbors, f.CellAt(col + 1, row))
    }

    if f.CellAt(col + 1, row - 1) != nil {
        neighbors = append(neighbors, f.CellAt(col + 1, row - 1))
    }

    if f.CellAt(col + 1, row + 1) != nil {
        neighbors = append(neighbors, f.CellAt(col + 1, row + 1))
    }

    return neighbors
}

func (f *Field) Width() int {
    return len(f.cells[0])
}

func (f *Field) Height() int {
    return len(f.cells)
}

func (f *Field) AreAllDead() bool {
    for y := 0; y < f.Height(); y++ {
        for x := 0; x < f.Width(); x++ {
            if f.CellAt(x, y).IsAlive() {
                return false
            }
        }
    }

    return true
}

func (f *Field) Equals(other *Field) bool {
    for y := 0; y < f.Height(); y++ {
        for x := 0; x < f.Width(); x++ {
            if f.CellAt(x, y).IsAlive() != other.CellAt(x, y).IsAlive() {
                return false
            }
        }
    }

    return true
}

func (f *Field) String() string {
    rows := make([]string, f.Height())

    for i := 0; i < f.Height(); i++ {
        row := make([]string, f.Width())
        for j := 0; j < f.Width(); j++ {
            row[j] = f.CellAt(j, i).String()
        }
        rows[i] = strings.Join(row, "")
    }

    return strings.Join(rows, "\n")
}
