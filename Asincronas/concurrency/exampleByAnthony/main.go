package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	userID := 10
	respch := make(chan string, 128)
	wg := &sync.WaitGroup{}

	go fetchUserData(userID, respch, wg)
	go fetchUserRecomendations(userID, respch, wg)
	go fetchUserLikes(userID, respch, wg)
	wg.Add(3)

	wg.Wait()
	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	respch <- "userData"
	wg.Done()
}
func fetchUserRecomendations(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	respch <- "userRecomndations"
	wg.Done()
}
func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)
	respch <- "userLikes"
	wg.Done()
}
