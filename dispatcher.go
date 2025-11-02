package main

type Direction int

const (
	Up Direction = iota
	Down
)

type DirectionRequest struct {
	Direction Direction
	Floor     int
}

type Dispatcher struct {
    ElevatorCount int
    Requests      map[int]Direction
    Elevators     map[int]*Elevator
}

func NewDispatcher(elevatorCount int) *Dispatcher {
    d := &Dispatcher{
        ElevatorCount: elevatorCount,
        Requests:      make(map[int]Direction),
        Elevators:     make(map[int]*Elevator),
    }

    for i := 0; i < elevatorCount; i++ {
        d.Elevators[i] = NewElevator(i)
    }
    return d
}

func (d *Dispatcher) AddRequest(req *DirectionRequest) {
    d.Requests[req.Floor] = req.Direction
    elevator := d.findBestElevator(req)
    elevator.AddRequest(req.Floor)
}

func (d *Dispatcher) findBestElevator(req *DirectionRequest) *Elevator {
    var best *Elevator
    var minScore float64 = float64(^uint(0) >> 1)

    for _, elevator := range d.Elevators {
        score := d.calculateScore(elevator, req)
        if score < minScore {
            minScore = score
            best = elevator
        }
    }

    return best
}

func (d *Dispatcher) calculateScore(elevator *Elevator, req *DirectionRequest) float64 {
    const (
        DISTANCE_WEIGHT   = 1.0
        DIRECTION_WEIGHT  = 2.0
        QUEUE_SIZE_WEIGHT = 0.5
        IDLE_BONUS        = -2.0
    )

    distance := float64(abs(elevator.Location - req.Floor))
    score := distance * DISTANCE_WEIGHT

    switch elevator.Status {
    case MovingUp:
        if req.Direction == Up {
            if req.Floor >= elevator.Location {
                score -= DIRECTION_WEIGHT
            } else {
                score += DIRECTION_WEIGHT
            }
        }
    case MovingDown:
        if req.Direction == Down {
            if req.Floor <= elevator.Location {
                score -= DIRECTION_WEIGHT
            } else {
                score += DIRECTION_WEIGHT
            }
        }
    case Idle:
        score += IDLE_BONUS
    }

    queueSize := float64(elevator.UpQueue.Len() + elevator.DownQueue.Len())
    score += queueSize * QUEUE_SIZE_WEIGHT

    return score
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}