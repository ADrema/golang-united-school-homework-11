package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(usersNumber int64, pool int64) (res []user) {

	userIds := make(chan int64, usersNumber)
	results := make(chan user, usersNumber)
	var i int64
	var j int64
	// iterate over channels number
	for ; i < pool; i++ {
		go worker(userIds, results)
	}
	for ; j < usersNumber; j++ {
		userIds <- j
	}
	close(userIds)

	// collect results into res var
	var a int64
	for ; a < usersNumber; a++ {
		u := <-results
		res = append(res, u)
	}
	return
}

func worker(users <-chan int64, results chan<- user) {
	for j := range users {
		results <- getOne(j)
	}
}
