package main

import (
    "flag"
    "fmt"
    "time"
)

var width, height int
var seedProbability float64

func init() {
    flag.IntVar(&width, "width", 10, "Board width")
    flag.IntVar(&height, "height", 10, "Board height")
    flag.Float64Var(&seedProbability, "seedProbability", 0.3, "Probability of cell to exists on initialization")
}

func main() {
    flag.Parse()

    world := NewWorld(width, height)
    world.RandomSeed(seedProbability)

    for {
        fmt.Println(world.String() + "\n")
        if world.IsFinished() {
            break
        }
        world.NextGen()
        time.Sleep(1 * time.Second)
    }
}
