# barbershop

    import "github.com/shantanubhadoria/go-sleeping-barber/barbershop"

Package sleepingbarber package for emulating the sleeping barber problem
https://en.wikipedia.org/wiki/Sleeping_barber_problem

## Installation

    $ go get github.com/shantanubhadoria/go-sleeping-barber/barbershop

## Running the code
Run without any parameters to see usage details

    $ go-sleeping-barber
    Usage: 
	    go-sleeping-barber <waiting room capacity> <haircut time in milliseconds> <average arrival rate in milliseconds> <shop open time in seconds>

for e.g. to run the stream for a barber 

    $ go-sleeping-barber 5 500 500 1
    Barber is sleeping: Zzzzzzzzzzzzz ...
    Adding Client 1
    Barber is cutting client 1's hair
    Barber is finished cutting client 1's hair
    Barber is sleeping: Zzzzzzzzzzzzz ...
    Adding Client 2
    Barber is cutting client 2's hair
    Closing shop
    Barber is finished cutting client 2's hair
    Barber is sleeping: Zzzzzzzzzzzzz ...
    Barber is going home
    Shop closed


## Synopsis

You may also include the module in your code and run this yourself

```go
import (
    "github.com/shantanubhadoria/go-sleeping-barber/barbershop"
)

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

```

## Usage

#### type BarberShop

```go
type BarberShop struct {
}
```


#### func  New

```go
func New(seatingCapacity int, hairCutDuration time.Duration) BarberShop
```
New returns a instance of a BarberShop

#### func (*BarberShop) AddBarber

```go
func (shop *BarberShop) AddBarber(barber string)
```
AddBarber processes clients and goes to sleep when shop is empty. Once the
clients channel is closed the barber goes home(returns). You may add more 
than one barber to address the multi barber version of this problem

#### func (*BarberShop) AddClient

```go
func (shop *BarberShop) AddClient(client string)
```
AddClient tries to add a client to the shop.

#### func (*BarberShop) Close

```go
func (shop *BarberShop) Close()
```
Close cleans up and closes the shop once all the barbers are done.
