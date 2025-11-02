# Elevator

### Compile

`go build .`

### Assumptions and Missing Features

The elevator system makes several simplifying assumptions:

- elevators operate with constant speed without acceleration or door timing
- start from floor 1, and have no maximum floor limit or capacity constraints
- continuous operation without maintenance needs or failures.
  Key features not implemented include:
- physical constraints (door operations, weight limits, variable speeds)
- safety measures (emergency stops, overload detection, door sensors)
- advanced scheduling (express elevators, peak hour optimization, energy efficiency)
- system management (logging, metrics, error handling)
- user interfaces (displays, control panels, configuration)

The implementation focuses on basic scheduling using a dispatcher with min-heap and max-heap queues on each elevator, leaving out complexities like request cancellation, priority access, and floor restrictions. Additionally, the system lacks concurrent operation handling, proper error management, and monitoring capabilities that would be essential in a real elevator system.
