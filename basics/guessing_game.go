package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func mai() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the Guassing Game!")
	fmt.Println("I hava chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		// Check if the guess if correct
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number")
		} else {
			fmt.Println("Too high! Try guessing a lower number")
		}
	}
}
