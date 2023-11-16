package main

import (
	"fmt"
	"sync"
)

type Plate struct {
	fruit1 string
	fruit2 string
	mutex  sync.Mutex
}

// PutFruit 临界资源
func (p *Plate) PutFruit(person string, fruit string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.fruit1 == "" {
		p.fruit1 = fruit
		if p.fruit1 == "🍎" {
			fmt.Printf("%s 🫱 %s\n", person, fruit)
		} else if p.fruit1 == "🍊" {
			fmt.Printf("%s 🫱 %s\n", person, fruit)
		}

	} else if p.fruit2 == "" {
		if p.fruit2 == "🍎" {
			fmt.Printf("%s 🫱 %s\n", person, fruit)
		} else if p.fruit2 == "🍊" {
			fmt.Printf("%s 🫱 %s\n", person, fruit)
		}
	} else {
		fmt.Println("盘子已满")
	}
}

// TakeFruit 临界资源
func (p *Plate) TakeFruit(person string, fruit string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.fruit1 == fruit {
		p.fruit1 = ""
		fmt.Printf("%s 🫳 %s\n", person, fruit)
	} else if p.fruit2 == fruit {
		p.fruit2 = ""
		fmt.Printf("%s 🫳 %s\n", person, fruit)
	} else {
		fmt.Printf("%s ❌  %s\n", person, fruit)
	}
}

func father(plate *Plate) {
	plate.PutFruit("👨", "🍎")
}

func mother(plate *Plate) {
	plate.PutFruit("👩", "🍊")
}

func son(plate *Plate) {
	plate.TakeFruit("👦", "🍊")
}

func daughter(plate *Plate) {
	plate.TakeFruit("👧", "🍎")
}

// WaitGroup 内部通过一个计数器来统计有多少协程被等待。这个计数器的值在我们启动 goroutine 之前先写入（使用 Add 方法），
// 然后在 goroutine 结束的时候，将这个计数器减 1（使用 Done 方法）。除此之外，在启动这些 goroutine 的协程中，
// 会调用 Wait 来进行等待，在 Wait 调用的地方会阻塞，直到 WaitGroup 内部的计数器减到 0。
func main() {
	plate := &Plate{}
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		father(plate)
	}()
	go func() {
		defer wg.Done()
		mother(plate)
	}()
	go func() {
		defer wg.Done()
		son(plate)
	}()
	go func() {
		defer wg.Done()
		daughter(plate)
	}()
	wg.Wait()
}
