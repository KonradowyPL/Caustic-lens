package main

import (
	"math"
)

var particles []Particle

type Particle struct {
	x float64
	y float64

	sx float64
	sy float64
}

var friction float64 = 0.99


func update() {
	//step 1: accelerate
	loop()

	//step 2: move
	move()

}

func loop() {
	for i1 := 0; i1 < len(particles); i1++ {
		thisParticle := &particles[i1]
		accelerate(thisParticle)

	}

}

func move() {
	for i := 0; i < len(particles); i++ {
		particle := &particles[i]

		particle.x += particle.sx
		particle.y += particle.sy

		particle.sx *= friction
		particle.sy *= friction

		checkBorder(particle)
	}
}

func checkBorder(particle *Particle) {
	if particle.x < 0 {
		particle.sx = math.Abs(particle.sx) 
		particle.x = 1
	}
	if particle.y < 0 {
		particle.sy = math.Abs(particle.sy)
		particle.y = 1
	}
	if particle.x > float64(screenX) {
		particle.sx = 0 - math.Abs(particle.sx) 
		particle.x = float64(screenX) - 1
	}
	if particle.y > float64(screenY) {
		particle.sy = 0 - math.Abs(particle.sy)
		particle.y = float64(screenY) - 1
	}
}

func accelerate(thisParticle *Particle) {
	var closest float64 = math.MaxFloat64
	var closestObj = Particle{0, 0, 0, 0}
	for i2 := 0; i2 < len(particles); i2++ {
		second := &particles[i2]

		if second != thisParticle {

			//step 1a calculate distance
			dx := thisParticle.x - second.x
			dy := thisParticle.y - second.y

			distance := math.Sqrt(dx*dx + dy*dy)


			if distance < closest {
				closest = distance
				closestObj = *second
			}

		}
	}
	//step 1b accelerate
	dx := thisParticle.x - closestObj.x
	dy := thisParticle.y - closestObj.y

	force := 0.0002

	if !math.IsNaN(force) && !math.IsInf(force, 1) && !math.IsInf(force, -1) {
		thisParticle.sx += force * dx
		thisParticle.sy += force * dy
	}
}
