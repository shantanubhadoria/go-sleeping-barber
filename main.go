package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/shantanubhadoria/go-sleeping-barber/barbershop"
)

func main() {
	args := os.Args
	usage := fmt.Sprintf(""+
		"Usage: \n"+
		"\t%s <waiting room capacity> <haircut time in milliseconds> <average arrival rate in milliseconds> <shop open time in seconds>", args[0])
	if len(args) != 5 {
		fmt.Println(usage)
	} else {
		seatingCapacity, err1 := strconv.Atoi(args[1])
		timePerHairCut, err2 := strconv.Atoi(args[2])
		arrivalRate, err3 := strconv.Atoi(args[3])
		openTime, err4 := strconv.Atoi(args[4])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println(err1, err2, err3, err4, usage)
		} else {
			runOperations(
				seatingCapacity,
				time.Millisecond*time.Duration(timePerHairCut),
				arrivalRate,
				time.Second*time.Duration(openTime),
			)
		}
	}
}

func runOperations(
	seatingCapacity int,
	timePerHairCut time.Duration,
	arrivalRate int,
	openTime time.Duration,
) {
	shop := barbershop.New(seatingCapacity, timePerHairCut)
	shop.AddBarber("Barber")
	i := 1

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(openTime)
		shopClosing <- true
		shop.Close()
		closed <- true
	}()

	go func() {
		for {
			// Get a random number with average arrival rate as specified
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.AddClient(strconv.Itoa(i))
				i++
			}
		}
	}()

	<-closed
}
