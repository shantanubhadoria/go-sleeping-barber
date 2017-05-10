# barbershop
--
    import "github.com/shantanubhadoria/go-sleeping-barber/barbershop"

Package sleepingbarber package for emulating the sleeping barber problem
https://en.wikipedia.org/wiki/Sleeping_barber_problem

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
clients channel is closed the barber goes home(returns).

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
