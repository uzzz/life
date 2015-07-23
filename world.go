package main

import (
    "math/rand"
    "time"
)

type World struct {
    currentGenField *Field
    nextGenField    *Field
    sameAsPrevGen   bool
}

func NewWorld(width, height int) *World {
    world := new(World)
    world.currentGenField = NewField(width, height)
    world.nextGenField = NewField(width, height)
    world.sameAsPrevGen = false
    return world
}

func (w *World) RandomSeed(probability float64) {
    rand.Seed(time.Now().Unix())

    for x := 0; x < w.currentGenField.Width(); x++ {
        for y := 0; y < w.currentGenField.Height(); y++ {
            cell := w.CellAt(x, y)
            if rand.Float64() >= (1 - probability)  {
                cell.MakeAlive()
            }
        }
    }
}

func (w *World) CellAt(x, y int) *Cell {
    return w.currentGenField.CellAt(x, y)
}

func (w *World) NextGen() {
    for x := 0; x < w.currentGenField.Width(); x++ {
        for y := 0; y < w.currentGenField.Height(); y++ {
            currentGenCell := w.currentGenField.CellAt(x, y)
            nextGenCell := w.nextGenField.CellAt(x, y)

            aliveCellsAround := w.currentGenField.AliveCellsAround(x, y)
            if aliveCellsAround == 3 {
                nextGenCell.MakeAlive()
            } else if aliveCellsAround == 2 && currentGenCell.IsAlive() {
                nextGenCell.MakeAlive()
            } else {
                nextGenCell.MakeDead()
            }
        }
    }

    if w.currentGenField.Equals(w.nextGenField) {
        w.sameAsPrevGen = true
    }

    tmp := w.currentGenField
    w.currentGenField = w.nextGenField
    w.nextGenField = tmp
}

func (w *World) IsFinished() bool {
    return w.AreAllDead() || w.SameAsPrevGen()
}

func (w *World) AreAllDead() bool {
    return w.currentGenField.AreAllDead()
}

func (w *World) SameAsPrevGen() bool {
    return w.sameAsPrevGen
}

func (w *World) String() string {
    return w.currentGenField.String()
}
