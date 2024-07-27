package main

import (
	"fmt"
	"sync"
	"time"
)

var mtx sync.RWMutex = sync.RWMutex{}
var wg sync.WaitGroup = sync.WaitGroup{}
var dbdata []string = []string{"db1", "db2", "db3", "db4", "db5"}
var results []string = []string{}


func main() {
	t := time.Now()
	for i := range dbdata {
		wg.Add(1)
		go dbcall(i)
		
	}
	wg.Wait()
	fmt.Println("Time taken: ", time.Since(t))
	fmt.Println("The results are: ", results)
}

func dbcall(i int) {
	//delay := rand.Float32()*2000
	//time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbdata[i])
	writeToDB("test",i)
	readFromDB()
	wg.Done()
}

func writeToDB(s string, i int) {
	mtx.Lock()
	dbdata[i] = s
	mtx.Unlock()
}

func readFromDB() {
	mtx.RLock()
	fmt.Println("dbdata is: ", dbdata)
	mtx.RUnlock()
}