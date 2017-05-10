// Package sleepingbarber package for emulating the sleeping barber problem https://en.wikipedia.org/wiki/Sleeping_barber_problem
package barbershop

import (
	"fmt"
	"time"
)

// BarberShop struct, use New to create a instance, do not init directly

type BarberShop struct {
	seatingCapacity int
	hairCutDuration time.Duration
	barbersCount    int
	barbersDone     chan bool // barbers indicate done status at end of day using these channels
	clients         chan string
	open            bool
}

// New returns a instance of a BarberShop
func New(seatingCapacity int, hairCutDuration time.Duration) BarberShop {
	shop := BarberShop{seatingCapacity: seatingCapacity, hairCutDuration: hairCutDuration}
	shop.clients = make(chan string, shop.seatingCapacity)
	shop.barbersCount = 0
	shop.barbersDone = make(chan bool)
	shop.open = true
	return shop
}

// Close cleans up and closes the shop once all the barbers are done.
func (shop *BarberShop) Close() {
	shop.open = false
	fmt.Println("Closing shop")
	close(shop.clients)
	for a := 1; a <= shop.barbersCount; a++ {
		<-shop.barbersDone
	}
	close(shop.barbersDone)
	fmt.Println("Shop closed")
}

// AddClient tries to add a client to the shop.
func (shop *BarberShop) AddClient(client string) {
	if shop.open {
		select {
		case shop.clients <- client:
			fmt.Printf("Adding Client %s\n", client)
		default:
			fmt.Println("shop's full, please come back later")
		}
	} else {
		fmt.Println("Shop already closed, cannot accept any more clients")
	}
}

// AddBarber processes clients and goes to sleep when shop is empty.
// Once the clients channel is closed the barber goes home(returns).
func (shop *BarberShop) AddBarber(barber string) {
	shop.barbersCount++
	go func() {
		for {
			if len(shop.clients) == 0 {
				fmt.Printf("%s is sleeping: Zzzzzzzzzzzzz ...\n", barber)
			}
			client, shopOpen := <-shop.clients
			if shopOpen {
				shop.cutHair(barber, client)
			} else {
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

// sendBarberHome makes the barber go home
func (shop *BarberShop) sendBarberHome(barber string) {
	fmt.Printf("%s is going home\n", barber)
	shop.barbersDone <- true
}

// cutHair makes the barber cut the given clients hair
func (shop *BarberShop) cutHair(barber string, client string) {
	fmt.Printf("%s is cutting client %s's hair\n", barber, client)
	time.Sleep(shop.hairCutDuration)
	fmt.Printf("%s is finished cutting client %s's hair\n", barber, client)
}
