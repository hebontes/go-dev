package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string,error){
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name bozo")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns on of a set of greeting messages. the Returned
// message is selected at random 
func randomFormat() string {
	// define a number variable

	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message formats by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
