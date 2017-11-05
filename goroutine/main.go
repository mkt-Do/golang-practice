package main

import (
  "log"
  "time"
)

func main() {

  c := make(chan string)

  log.Print("Start")

  go oneSec(c)
  go twoSec(c)
  go  threeSec(c)

  r1, r2, r3 := <-c, <-c, <-c

  log.Println(r1, r2, r3)

  log.Print("all finished")
}

func oneSec(c chan string) {
  log.Print("sleep1 started")
  time.Sleep(1 * time.Second)
  c <- "sleep1 finished"
}

func twoSec(c chan string) {
  log.Print("sleep2 started")
  time.Sleep(2 * time.Second)
  c <- "sleep2 finished"
}
 
func threeSec(c chan string) {
  log.Print("sleep3 started")
  time.Sleep(3 * time.Second)
  c <- "sleep3 finished"
}
