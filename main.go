package main

import "fmt"

var resultChan = make(chan string)
var done = make(chan bool)

func main() {
	go func() {
		for i := 1; i < 101; i++ {
			if i%2 == 0 {
				resultChan <- fmt.Sprintf("success")
			} else {
				resultChan <- fmt.Sprintf("数字%d不是偶数", i)
			}
		}
	}()

	length := 0
	var err []string
	go func() {
		for {
			select {
			case r := <-resultChan:
				length += 1
				if r == "success" {
					if length == 100 {
						close(resultChan)
						done <- true
					}
				} else if r == "" {
					fmt.Println("读到空字符串")
				} else {
					err = append(err, r)
				}
			}
		}
	}()

	<-done

	if len(err) > 0 {
		for _, v := range err {
			fmt.Println(v)
		}
	}
}

