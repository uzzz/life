package main

type Cell struct {
    alive bool
}

func (c *Cell) IsAlive() bool {
    return c.alive == true
}

func (c *Cell) MakeAlive() {
    c.alive = true
}

func (c *Cell) MakeDead() {
    c.alive = false
}

func (c *Cell) String() string {
    if c.alive {
        return "*"
    } else {
        return " "
    }
}
