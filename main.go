package main

import (
	"fmt"
	"math/rand"
	"time"
)

var facts = []string{
	"wheels are often made out of materials",
	"wheels sometimes have spokes in the middle",
	"a wheel is not a wheel if it's not a wheel",
	"a wheely-bin is so called because of its wheels",
	"we would not be where we are without wheels",
	"some cars actually have wheels to help the vehicle move",
	"wheeeeeel",
	"half the merriment of a wheel is that it spins",
	"did you know? wheel",
	"i wheel, you wheel, she wheels",
	"to wheel or not to wheel; that is the question",
	"standard-sized wheels may be a choking hazard to children with extremely large mouths",
	"i once saw a wheel on the bus go round and round...",
}

func main() {
	fmt.Println("Right, let's scan your code...")
	time.Sleep(1 * time.Second)
	fmt.Println("Just kidding! I don't care about your code!")
	fmt.Println("Here's a fun fact about wheels:", facts[rand.Intn(len(facts))])
}
