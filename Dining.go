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
func aPhilosopher(index int, leftFork int, rightFork int, chL chan bool, chR chan bool) {
	// Acces the specific philosopher in the philosophers slice.
	philosopher := &philosophers[index]

	for true{
		
	
	// If the philosopher has eaten 3 times it set philosopher.doneEating to true
	if philosopher.timesEaten == 3 {
		philosopher.doneEating = true
		fmt.Println(philosopher.name + " done")
	}

	// Logic to control when the philosopher should eat or think.
	if !philosopher.doneEating {

		if !forks[leftFork].forkIsActive {
			philosopher.useLeftFork = true
			chL <- true
		}

		if !forks[rightFork].forkIsActive {
			philosopher.useRightFork = true
			chR <- true
		}

		if philosopher.useLeftFork == true && philosopher.useRightFork == true {
			philosopher.eat = true
			fmt.Println(philosopher.name + " eating")
			philosopher.timesEaten++
			fmt.Println(philosopher.name+" has eaten", philosopher.timesEaten, "times")
		}
		if philosopher.eat && philosopher.useLeftFork == true {
			philosopher.useLeftFork = false
			chL <- false
		}
		
		if philosopher.eat && philosopher.useRightFork == true {
			philosopher.useRightFork = false
			chR <- false
		}
		if(philosopher.eat){
			philosopher.eat = false
			fmt.Println(philosopher.name + " thinking")
		}
		
	}
	if(philosopher.doneEating){
		break;
	}
}
}

// Function to create fork goroutines.
func aFork(index int, ch chan bool) {
	// Acces the specific fork in the forks slice.
	fork := &forks[index]
	for(true){
		boolVal := <-ch
		fork.forkIsActive = boolVal
	}
	

	// Channel from aPhilosopher which tells aFork if its active.
	
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

	ch0 := make(chan bool)
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	ch4 := make(chan bool)
	
		go aPhilosopher(0, 4, 0, ch4, ch0)
		go aPhilosopher(1, 0, 1, ch0, ch1)
		go aPhilosopher(2, 2, 1, ch2, ch1)
		go aPhilosopher(3, 2, 3, ch2, ch3)
		go aPhilosopher(4, 3, 4, ch3, ch4)
		
		go aFork(0, ch0)
		go aFork(1, ch1)
		go aFork(2, ch2)
		go aFork(3, ch3)
		go aFork(4, ch4)

	//fmt.Println(philosophers[0].name+" has eaten", philosophers[0].timesEaten, "times")
	/*fmt.Println(philosophers[1].name+" has eaten", philosophers[1].timesEaten, "times")
	fmt.Println(philosophers[2].name+" has eaten", philosophers[2].timesEaten, "times")
	fmt.Println(philosophers[3].name+" has eaten", philosophers[3].timesEaten, "times")
	fmt.Println(philosophers[4].name+" has eaten", philosophers[4].timesEaten, "times")*/
	time.Sleep(1000*time.Millisecond)
}
