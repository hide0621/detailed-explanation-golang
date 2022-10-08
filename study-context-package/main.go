package main

import (
	"fmt"
	"sync"
	"time"
)

func shortProcess(done <-chan interface{}) (bool, error) {
	shortWork := time.NewTicker(1 * time.Second)

	select {
	case <-done:
		return false, fmt.Errorf("cancel")
	case <-shortWork.C:
	}
	return true, nil
}

func longProcess(done <-chan interface{}) (bool, error) {
	longWork := time.NewTicker(5 * time.Second)

	select {
	case <-done:
		return false, fmt.Errorf("キャンセル")
	case <-longWork.C:
	}
	return true, nil
}

func main() {
	var wg sync.WaitGroup
	done := make(chan interface{})

	defer close(done)

	//チャネルをキャンセルさせてdoneの時の結果を返す
	// go func() {
	// 	time.Sleep(1 * time.Millisecond)
	// 	close(done)
	// }()

	//サブルーチン作成
	wg.Add(1)
	go func() {
		defer wg.Done()
		if isDone, err := shortProcess(done); err != nil {
			fmt.Printf("shortProcess: %v\n", err)
			fmt.Println(isDone)
			return
		}
		fmt.Println("shortProcess is Done")
	}()

	//サブルーチン作成
	wg.Add(1)
	go func() {
		defer wg.Done()
		if isDone, err := longProcess(done); err != nil {
			fmt.Printf("longProcess: %v\n", err)
			fmt.Println(isDone)
			return
		}
		fmt.Println("longProcess is Done")
	}()

	//メインルーチンへの合流
	wg.Wait()
}
