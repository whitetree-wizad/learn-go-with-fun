package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	var aspect command
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				aspect = right
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				aspect = left
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case stop:
				log.Printf("rover stop at %v", pos)
				direction = image.Point{
					X: 0,
					Y: 0,
				}
			case start:
				log.Printf("rover begin to move")
				direction = image.Point{X: 1, Y: 0}
				switch aspect {
				case right:
					direction = image.Point{
						X: -direction.Y,
						Y: direction.X,
					}
				case left:
					direction = image.Point{
						X: direction.Y,
						Y: -direction.X,
					}
				}
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)
	r.Left()
}
