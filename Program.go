// package main

// import (
// 	"time"
// )

// type Fork struct{ available bool }

// func main() {
// 	forkOne := make(chan *Fork)
// 	forkTwo := make(chan *Fork)
// 	forkThree := make(chan *Fork)
// 	forkFour := make(chan *Fork)
// 	forkFive := make(chan *Fork)
// 	go philosopher("p1", forkOne, forkTwo)

// 	forkOne <- new(Fork)
// 	forkTwo <- new(Fork)
// 	forkThree <- new(Fork)
// 	forkFour <- new(Fork)
// 	forkFive <- new(Fork)
// 	time.Sleep(5 * time.Second)
// 	<-forkOne
// 	<-forkTwo
// 	<-forkThree
// 	<-forkFour
// 	<-forkFive
// }

// func philosopher(name string, forkLeft chan *Fork, forkRight chan *Fork) {
// 	var timesEaten int = 0
// 	for timesEaten < 3 {
// 		think()
// 		interactLeftFork()
// 		interactRightFork()
// 		eat(timesEaten)

// 	}
// }

// func interactLeftFork() {

// }

// func interactRightFork() {

// }

// func eat(timesEaten int) {
// 	timesEaten++
// 	interactLeftFork()
// 	interactRightFork()
// }

// func think() {

// }
