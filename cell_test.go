package main

import "testing"

func TestCell_IsAlive(t *testing.T) {
    cell := new(Cell)
    if cell.IsAlive() {
        t.Error("expected cell to not be alive")
    }
    cell.alive = true
    if !cell.IsAlive() {
        t.Error("expected cell to be alive")
    }
}

func TestCell_MakeAlive(t *testing.T) {
    cell := new(Cell)
    cell.MakeAlive()
    if !cell.alive {
        t.Error("expected cell to be alive")
    }
}

func TestCell_MakeDead(t *testing.T) {
    cell := new(Cell)
    cell.alive = true
    cell.MakeDead()
    if cell.IsAlive() {
        t.Error("expected cell to not be alive")
    }
}

func TestCell_String(t *testing.T) {
    cell := new(Cell)
    if cell.String() != " " {
        t.Error("expected cell's string representation to be a space")
    }
    cell.MakeAlive()
    if cell.String() != "*" {
        t.Error("expected cell's string representation to be a *")
    }
}
