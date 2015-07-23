package main

import "testing"

func TestField_Width(t *testing.T) {
    field := NewField(10, 20)
    if field.Width() != 10 {
        t.Error("expected width to equal 10")
    }
}

func TestField_Height(t *testing.T) {
    field := NewField(10, 20)
    if field.Height() != 20 {
        t.Error("expected width to equal 20")
    }
}

func TestField_CellAt(t *testing.T) {
    field := NewField(5, 5)
    if field.CellAt(-1, 0) != nil {
        t.Error("expected to return nil when X is out of range")
    }

    if field.CellAt(5, 0) != nil {
        t.Error("expected to return nil when X is out of range")
    }

    if field.CellAt(0, -1) != nil {
        t.Error("expected to return nil when Y is out of range")
    }

    if field.CellAt(0, 5) != nil {
        t.Error("expected to return nil when Y is out of range")
    }

    if field.CellAt(0, 0) == nil {
        t.Error("expected to return a cell")
    }
}

func TestField_AreAllDead(t *testing.T) {
    field := NewField(5, 5)
    if !field.AreAllDead() {
        t.Error("expected to all be dead")
    }

    field.CellAt(0, 0).MakeAlive()
    if field.AreAllDead() {
        t.Error("expected to not all be dead")
    }
}

func TestField_Equals(t *testing.T) {
    field1 := NewField(5, 5)
    field2 := NewField(5, 5)

    if !field1.Equals(field2) {
        t.Error("expected field1 and field2 to be equals")
    }

    // Equality should be commutative
    if !field2.Equals(field1) {
        t.Error("expected field1 and field2 to be equals")
    }

    field1.CellAt(0, 0).MakeAlive()

    if field1.Equals(field2) {
        t.Error("expected field1 and field2 to not be equal")
    }

    // Equality should be commutative
    if field2.Equals(field1) {
        t.Error("expected field1 and field2 to not be equal")
    }
}

func TestField_AliveCellsAround(t *testing.T) {
    field := NewField(3, 3)

    // x - -
    // - - -
    // - - -
    if field.AliveCellsAround(0, 0) != 0 {
        t.Error("expected 0 alive cells around")
    }

    // x * -
    // - - -
    // - - -
    field.CellAt(1, 0).MakeAlive()
    if field.AliveCellsAround(0, 0) != 1 {
        t.Error("expected 1 alive cells around")
    }

    // x * -
    // - * -
    // - - -
    field.CellAt(1, 1).MakeAlive()
    if field.AliveCellsAround(0, 0) != 2 {
        t.Error("expected 2 alive cells around")
    }

    // x * -
    // - * -
    // - - *
    field.CellAt(2, 2).MakeAlive()
    if field.AliveCellsAround(0, 0) != 2 {
        t.Error("expected 2 alive cells around")
    }

    // - * -
    // - * -
    // x - *
    if field.AliveCellsAround(0, 2) != 1 {
        t.Error("expected 1 alive cells around")
    }

    // - * -
    // - * -
    // - x *
    if field.AliveCellsAround(1, 2) != 2 {
        t.Error("expected 2 alive cells around")
    }
}

func TestField_String(t *testing.T) {
    field := NewField(3, 3)

    if field.String() != "   \n" +
                         "   \n" +
                         "   " {
        t.Error()
    }

    field.CellAt(0, 0).MakeAlive()
    if field.String() != "*  \n" +
                         "   \n" +
                         "   " {
        t.Error()
    }

    field.CellAt(2, 2).MakeAlive()
    if field.String() != "*  \n" +
                         "   \n" +
                         "  *" {
        t.Error()
    }

    field.CellAt(1, 1).MakeAlive()
    if field.String() != "*  \n" +
                         " * \n" +
                         "  *" {
        t.Error()
    }
}
