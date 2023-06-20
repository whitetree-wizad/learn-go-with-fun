package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left = command(1)
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r:=&RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos:=image.Point{X:0,Y:0}
	direction:=image.Point{X: 1,Y: 0}
	updateInterval:=250*time.Millisecond
	nextMove:=time.After(updateInterval)
	for  {
		select {
		case c:=<-r.commandc:
			switch c {
			case right:
				direction = image.Point{
				X: -direction.Y,
				Y: direction.X,
				}
			case left:
				direction=image.Point{
				X: direction.Y,
				Y: -direction.X,
				}
			}
			log.Printf("new direction %v",direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("move to %v",pos)
			nextMove=time.After(updateInterval)
		}
	}
}
