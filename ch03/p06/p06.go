package main

import (
	"errors"
	"fmt"
	"time"
)

type Species int

const (
	Dog Species = iota + 1
	Cat
)

type AnimalQueue struct {
	queue []*Animal
}

func NewAnimalQueue() *AnimalQueue {
	return &AnimalQueue{
		queue: make([]*Animal, 0),
	}
}

func (q *AnimalQueue) enqueue(element *Animal) {
	if q.queue == nil {
		q.queue = make([]*Animal, 0)
	}
	q.queue = append(q.queue, element)
}

func (q *AnimalQueue) dequeue() *Animal {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}

func (q *AnimalQueue) peek() *Animal {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	return result
}

func (q *AnimalQueue) isEmpty() bool {
	return len(q.queue) == 0
}

type Animal struct {
	spec      Species
	timestamp int64
}

type Shelter struct {
	dogs *AnimalQueue
	cats *AnimalQueue
}

func NewShelter() *Shelter {
	return &Shelter{
		dogs: NewAnimalQueue(),
		cats: NewAnimalQueue(),
	}
}

func (s *Shelter) enqueue(element Animal) error {
	element.timestamp = time.Now().Unix()

	switch element.spec {
	case Dog:
		s.dogs.enqueue(&element)
	case Cat:
		s.cats.enqueue(&element)
	default:
		return errors.New("the animal is not supported")
	}

	return nil
}

func (s *Shelter) dequeueDog() *Animal {
	if s.dogs == nil || s.dogs.isEmpty() {
		return nil
	}
	return s.dogs.dequeue()
}

func (s *Shelter) dequeueCat() *Animal {
	if s.cats == nil || s.cats.isEmpty() {
		return nil
	}
	return s.cats.dequeue()
}

func (s *Shelter) dequeueAny() *Animal {
	if s.cats.isEmpty() {
		return s.dequeueDog()
	} else if s.dogs.isEmpty() {
		return s.dequeueCat()
	}

	d := s.dogs.peek()
	c := s.cats.peek()

	if d.timestamp < c.timestamp {
		return s.dequeueDog()
	}
	return s.dequeueCat()
}

func main() {
	s := NewShelter()
	_ = s.enqueue(Animal{spec: Cat})
	_ = s.enqueue(Animal{spec: Dog})
	_ = s.enqueue(Animal{spec: Cat})
	_ = s.enqueue(Animal{spec: Cat})
	fmt.Println(s)
	a := s.dequeueAny()
	d := s.dequeueDog()
	c := s.dequeueCat()
	fmt.Println(a, d, c)
}
