package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MAXNUM = 5
const MAXEATCOUNT = 3
const MAXCONCURRENTDINERS = 2

type ChopStick struct {
	number int
	sync.Mutex
}

type Philosopher struct {
	number         int
	eatCount       int
	leftChopStick  *ChopStick
	rightChopStick *ChopStick
}

// Host buffered channel of size MAXCONCURRENTDINERS
type Host struct {
	philosopherChannel chan *Philosopher
}

// Allow wait until we get MAXCONCURRENTDINERS requests
func (h *Host) Allow(wg *sync.WaitGroup) {
	for {
		if len(h.philosopherChannel) == MAXCONCURRENTDINERS {
			<-h.philosopherChannel
			<-h.philosopherChannel
		}
	}
}

// Eat lock resources (chop sticks) and wait a random time to eat
func (p *Philosopher) Eat(wg *sync.WaitGroup, host *Host) {
	for i := 0; i < MAXEATCOUNT; i++ {
		host.philosopherChannel <- p
		if p.eatCount < 3 {
			p.leftChopStick.Lock()
			p.rightChopStick.Lock()
			fmt.Printf("Starting to eat %d\n", p.number)

			p.eatCount += 1
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // random time to eat

			fmt.Printf("Finishing eating: %d\n", p.number)
			p.leftChopStick.Unlock()
			p.rightChopStick.Unlock()
			wg.Done()
		}
	}
}

func makeAllChopSticks() []*ChopStick {
	allChopSticks := make([]*ChopStick, MAXNUM)
	for i := 0; i < len(allChopSticks); i++ {
		allChopSticks[i] = new(ChopStick)
		allChopSticks[i].number = i
	}
	return allChopSticks
}

func makeAllPhilosophers(allChopSticks []*ChopStick) []*Philosopher {
	allPhilosophers := make([]*Philosopher, MAXNUM)
	for i := 0; i < len(allPhilosophers); i++ {
		allPhilosophers[i] = &Philosopher{
			number:         i,
			leftChopStick:  allChopSticks[i],
			rightChopStick: allChopSticks[(i+1)%MAXNUM],
		}
	}
	return allPhilosophers
}

func main() {
	var wg sync.WaitGroup

	philosopherChannel := make(chan *Philosopher, MAXCONCURRENTDINERS)

	allChopSticks := makeAllChopSticks()
	allPhilosophers := makeAllPhilosophers(allChopSticks)

	fmt.Println(allPhilosophers)

	host := &Host{philosopherChannel}

	wg.Add(MAXNUM * MAXEATCOUNT)

	go host.Allow(&wg)

	for _, p := range allPhilosophers {
		go p.Eat(&wg, host)
	}

	wg.Wait()
}
