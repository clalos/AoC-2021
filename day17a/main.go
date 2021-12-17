/**
https://adventofcode.com/2021/day/17
*/
package main

import "fmt"

type coord struct {
	x, y int
}

type velocity struct {
	x, y int
}

func main() {
	targetX := [2]int{81, 129}
	targetY := [2]int{-150, -108}

	velocities := generateVelocities(200)
	var trajectories [][]coord
	var validVelocities []velocity
	for index, speed := range velocities {
		point := coord{0, 0}
		var trajectory []coord
		var targetHit bool
		for steps := 0; steps < 500; steps++ {
			trajectory = append(trajectory, point)

			if point.x >= targetX[0] && point.x <= targetX[1] &&
				point.y >= targetY[0] && point.y <= targetY[1] {
				// Target hit.
				targetHit = true
				break
			}
			point.x += speed.x
			point.y += speed.y

			if speed.x > 0 {
				speed.x -= 1
			}
			if speed.x < 0 {
				speed.x += 1
			}
			speed.y -= 1
		}
		if targetHit {
			trajectories = append(trajectories, trajectory)
			validVelocities = append(validVelocities, velocities[index])
		}
	}

	fmt.Println(len(trajectories), " trajectories found")

	if len(trajectories) == 0 {
		return
	}

	maximumY := trajectories[0][0].y
	trajectoryIndex := 0
	for index, trajectory := range trajectories {
		for _, coordinate := range trajectory {
			if coordinate.y > maximumY {
				maximumY = coordinate.y
				trajectoryIndex = index
			}
		}
	}

	fmt.Println("Best velocity: ", validVelocities[trajectoryIndex].x, ",", validVelocities[trajectoryIndex].y)
	fmt.Println("Highest Y: ", maximumY)
}

// generateVelocities generate all possible velocities in the given range.
func generateVelocities(n int) []velocity {
	var velocities []velocity
	for i := -n; i <= n; i++ {
		for j := -n; j <= n; j++ {
			velocities = append(velocities, velocity{i, j})
		}
	}
	return velocities
}
