package main

import "testing"

func TestWorld_NextGen(t *testing.T) {
    var world *World

    // * - - - -
    // - - - - -
    // - - - - -
    // - - - - -
    world = NewWorld(5, 5)
    world.CellAt(0, 0).MakeAlive()
    world.NextGen()
    if world.CellAt(0, 0).IsAlive() {
        t.Error("cell at 0 0 expected to die")
    }

    // * * - - -
    // * * - - -
    // - - - - -
    // - - - - -
    world = NewWorld(5, 5)
    world.CellAt(0, 0).MakeAlive()
    world.CellAt(1, 0).MakeAlive()
    world.CellAt(0, 1).MakeAlive()
    world.CellAt(1, 1).MakeAlive()
    world.NextGen()
    if !world.CellAt(0, 0).IsAlive() || !world.CellAt(1, 0).IsAlive() || !world.CellAt(0, 1).IsAlive() || !world.CellAt(1, 1).IsAlive() {
        t.Errorf("figure should remain unchanged, got \n%s", world.String())
    }

    // - * - - -
    // - * - - -
    // - * - - -
    // - - - - -
    world = NewWorld(5, 5)
    world.CellAt(1, 0).MakeAlive()
    world.CellAt(1, 1).MakeAlive()
    world.CellAt(1, 2).MakeAlive()
    world.NextGen()
    if !world.CellAt(0, 1).IsAlive() || !world.CellAt(1, 1).IsAlive() || !world.CellAt(2, 1).IsAlive() {
        t.Errorf("wrong figure, got \n%s", world.String())
    }
    world.NextGen()
    if !world.CellAt(1, 0).IsAlive() || !world.CellAt(1, 1).IsAlive() || !world.CellAt(1, 2).IsAlive() {
        t.Errorf("wrong figure, got \n%s", world.String())
    }
}

func TestWorld_IsFinished(t *testing.T) {
    var world *World

    // all cells dies
    world = NewWorld(5, 5)
    world.CellAt(0, 0).MakeAlive()
    world.NextGen()
    if !world.IsFinished() {
        t.Error("expect game to be finished")
    }

    // stable figure
    world = NewWorld(5, 5)
    world.CellAt(0, 0).MakeAlive()
    world.CellAt(1, 0).MakeAlive()
    world.CellAt(0, 1).MakeAlive()
    world.CellAt(1, 1).MakeAlive()
    world.NextGen()
    if !world.IsFinished() {
        t.Error("expect game to be finished")
    }
}
