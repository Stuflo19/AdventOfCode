package day14

import (
	"image"
	"strconv"
	"strings"
)

type Robot struct {
	pos      image.Point
	velocity image.Point
}

func parseXY(location string) image.Point {
	xySplit := strings.Split(location, ",")

	intX, _ := strconv.Atoi(xySplit[0])
	intY, _ := strconv.Atoi(xySplit[1])

	return image.Point{
		X: intX,
		Y: intY,
	}
}

func NewRobot(line string) Robot {
	splitPV := strings.Split(line, " ")
	position := image.Point{}
	velocity := image.Point{}

	for _, v := range splitPV {
		if pos, ok := strings.CutPrefix(v, "p="); ok {
			position = parseXY(pos)
		} else if vel, ok := strings.CutPrefix(v, "v="); ok {
			velocity = parseXY(vel)
		}
	}

	return Robot{
		pos:      position,
		velocity: velocity,
	}
}

func (r Robot) GetPos() image.Point {
	return r.pos
}

func (r *Robot) MoveTimes(times int, maxX, maxY int) {
	r.pos.X += (r.velocity.X * times) % maxX
	r.pos.Y += (r.velocity.Y * times) % maxY

	if r.pos.X >= maxX {
		r.pos.X = r.pos.X - maxX
	}
	if r.pos.Y >= maxY {
		r.pos.Y = r.pos.Y - maxY
	}

	if r.pos.X < 0 {
		r.pos.X = maxX + r.pos.X
	}
	if r.pos.Y < 0 {
		r.pos.Y = maxY + r.pos.Y
	}
}

func (r Robot) GetQuadrant(maxX, maxY int) int {
	if r.pos.X < maxX/2 && r.pos.Y < maxY/2 {
		return 1
	} else if r.pos.X > maxX/2 && r.pos.Y < maxY/2 {
		return 2
	} else if r.pos.X > maxX/2 && r.pos.Y > maxY/2 {
		return 3
	} else if r.pos.X < maxX/2 && r.pos.Y > maxY/2 {
		return 4
	}

	return -1
}
