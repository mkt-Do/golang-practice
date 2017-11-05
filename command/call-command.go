package main

import (
  "log"
  "fmt"
  "os/exec"
)

func main() {
  log.Print("This is the program to execute command\n");
  cmd := exec.Command("sleep", "5")
  error := cmd.Start()

  if error != nil {
    panic(fmt.Sprintf("error1: %v", error))
  }

  log.Print("start command")

  error = cmd.Wait()

  if error != nil {
    panic(fmt.Sprintf("error2: %v", error))
  }

  log.Print("finish command")
}
