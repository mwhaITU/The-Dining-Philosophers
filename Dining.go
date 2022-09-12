package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	name         string
	think        bool
	eat          bool
	useLeftFork  bool
	useRightFork bool
	timesEaten   int
	doneEating   bool
}

type Fork struct {
	forkIsActive bool
}

var philosophers []Philosopher
var forks []Fork

// Function to create philolopher goroutines.
func aPhilosopher(index int, leftFork int, rightFork int, ch chan bool) {
	// Acces the specific philosopher in the philosophers slice.
	philosopher := &philosophers[index]

	// If the philosopher has eaten 3 times it set philosopher.doneEating to true
	if philosopher.timesEaten == 3 {
		philosopher.doneEating = true
		fmt.Println(philosopher.name + " done")
	}

	// Logic to control when the philosopher should eat or think.
	if !philosopher.doneEating {

		fmt.Println(philosopher.name + " thinking")

		if !forks[leftFork].forkIsActive {
			philosopher.useLeftFork = true
			ch <- true
		}

		if !forks[rightFork].forkIsActive {
			philosopher.useRightFork = true
			ch <- true
		}

		if philosopher.useLeftFork == true && philosopher.useRightFork == true {
			philosopher.eat = true
			fmt.Println(philosopher.name + " eating")
			philosopher.timesEaten++
		}

		if philosopher.eat && philosopher.useLeftFork == true {
			philosopher.useLeftFork = false
			ch <- false
		}

		if philosopher.eat && philosopher.useRightFork == true {
			philosopher.useRightFork = false
			ch <- false
		}

		if philosopher.eat && philosopher.useLeftFork == true && philosopher.useRightFork == true {
			philosopher.eat = false
			fmt.Println(philosopher.name + " thinking")
		}
	}
}

// Function to create fork goroutines.
func aFork(index int, ch chan bool) {
	// Acces the specific fork in the forks slice.
	fork := &forks[index]

	// Channel from aPhilosopher which tells aFork if its active.
	boolVal := <-ch
	fork.forkIsActive = boolVal
}

func main() {
	// Two parallel slices:
	// philosopher Bob
	//                     fork 0
	// philosopher Joe
	//                     fork 1
	// philosopher Ben
	//                     fork 2
	// philosopher Jack
	//                     fork 3
	// philosopher Steve
	// 					   fork 4

	// Philosophers is being created.
	philosophers = append(philosophers, Philosopher{name: "Bob", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})
	philosophers = append(philosophers, Philosopher{name: "Joe", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})
	philosophers = append(philosophers, Philosopher{name: "Ben", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})
	philosophers = append(philosophers, Philosopher{name: "Jack", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})
	philosophers = append(philosophers, Philosopher{name: "Steve", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})

	// Forks is beining created.
	forks = append(forks, Fork{forkIsActive: false})
	forks = append(forks, Fork{forkIsActive: false})
	forks = append(forks, Fork{forkIsActive: false})
	forks = append(forks, Fork{forkIsActive: false})
	forks = append(forks, Fork{forkIsActive: false})

	ch := make(chan bool)

	// For-loop runs while the philosophers is not done eating.
	for !philosophers[0].doneEating && !philosophers[1].doneEating && !philosophers[2].doneEating && !philosophers[3].doneEating && !philosophers[4].doneEating {
		if !philosophers[0].doneEating {
			go aPhilosopher(0, 4, 0, ch)
		}
		if !philosophers[1].doneEating {
			go aPhilosopher(1, 0, 1, ch)
		}
		if !philosophers[2].doneEating {
			go aPhilosopher(2, 1, 2, ch)
		}
		if !philosophers[3].doneEating {
			go aPhilosopher(3, 2, 3, ch)
		}
		if !philosophers[4].doneEating {
			go aPhilosopher(4, 3, 4, ch)
		}

		go aFork(0, ch)
		go aFork(1, ch)
		go aFork(2, ch)
		go aFork(3, ch)
		go aFork(4, ch)

		time.Sleep(time.Millisecond)
	}

	fmt.Println(philosophers[0].name+" has eaten", philosophers[0].timesEaten, "times")
	fmt.Println(philosophers[1].name+" has eaten", philosophers[1].timesEaten, "times")
	fmt.Println(philosophers[2].name+" has eaten", philosophers[2].timesEaten, "times")
	fmt.Println(philosophers[3].name+" has eaten", philosophers[3].timesEaten, "times")
	fmt.Println(philosophers[4].name+" has eaten", philosophers[4].timesEaten, "times")
}
