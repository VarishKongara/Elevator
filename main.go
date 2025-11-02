package main

import (
	"fmt"
)

func main() {
    dispatcher := NewDispatcher(3)

    // sample requests
    dispatcher.AddRequest(&DirectionRequest{Direction: Up, Floor: 5})
    dispatcher.AddRequest(&DirectionRequest{Direction: Down, Floor: 2})
    dispatcher.AddRequest(&DirectionRequest{Direction: Up, Floor: 8})

    // simple simulation loop
    for tick := 0; tick < 20; tick++ {
        for _, e := range dispatcher.Elevators {
            e.Move()
            fmt.Printf("tick=%02d elevator=%d floor=%d status=%v\n", tick, e.ID, e.Location, e.Status)
        }
    }
}
