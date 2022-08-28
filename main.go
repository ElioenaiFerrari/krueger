/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ElioenaiFerrari/krueger/cmd"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println(`
	⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥⬜⬜🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥⬜⬜🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟩🟩🟩⬜⬜⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟩🟩🟩🟩🟩🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟩🟩🟩🟩🟩🟩🟩🟩🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜🟩🟩🟩🟩🟩🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟩⬜⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜🟩🟩🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟩🟩🟩🟩🟩🟩🟩⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟩🟩🟩🟩🟩🟩🟩🟩🟩🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜🟩🟩🟩🟩🟩🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟩🟩🟩🟩⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟥🟥🟥🟥⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜🟩🟩🟩🟩🟩🟩🟩🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜
  ⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜
  ⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜
  ⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜
  ⬜⬜⬜⬜⬜⬜🟥🟥🟥🟥⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜
  ⬜⬜⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜🟥🟥🟥🟥⬜
  ⬜⬜⬜⬜⬜🟥🟥🟥🟥⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜🟥🟥🟥🟥⬜
  ⬜⬜⬜⬜🟥🟥🟥🟥🟥⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜🟥🟥🟥🟥⬜
  ⬜⬜⬜⬜🟥🟥🟥🟥🟥⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜🟥🟥🟥⬜
  ⬜⬜⬜🟥🟥🟥🟥🟥⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜⬜🟥🟥🟥⬜
  ⬜⬜⬜⬜🟥🟥🟥🟥⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜🟥🟥🟥🟥
  ⬜⬜⬜🟥🟥🟥🟥🟥⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜🟥🟥🟥🟥
  ⬜⬜⬜🟥🟥🟥🟥⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜⬜🟥🟥🟥⬜
  ⬜⬜⬜🟥🟥🟥🟥⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜🟥🟥🟥🟥🟥
  ⬜⬜⬜🟥🟥🟥⬜⬜⬜🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬜🟥🟥🟥🟥⬜
  ⬜⬜🟥🟥🟥🟥⬜⬛⬜⬜⬛⬜⬛⬛⬛⬜⬛⬜⬛⬜⬛⬛⬛⬜⬛⬛⬛⬜🟥🟥🟥🟥⬜
  ⬜⬜🟥🟥🟥⬜⬜⬛⬛⬜⬛⬜⬛⬜⬜⬜⬛⬜⬛⬜⬛⬜⬜⬜⬛⬜⬛🟥⬜🟥🟥🟥⬜
  ⬜🟥🟥🟥🟥⬜⬜⬛⬛⬛⬛⬜⬛⬛⬜⬜⬛⬜⬛⬜⬛⬛⬜⬜⬛⬛⬜⬜⬜🟥🟥🟥⬜
  🟥🟥🟥🟥⬜⬜⬜⬛⬜⬛⬛⬜⬛⬜⬜⬜⬜⬛⬜⬜⬛⬜⬜⬜⬛⬜⬛⬜⬜🟥🟥⬜⬜
  🟥🟥🟥🟥🟥⬜⬜⬛⬜⬜⬛⬜⬛⬛⬛⬜⬜⬛⬜⬜⬛⬛⬛⬜⬛⬜⬛⬜🟥⬜⬜⬜⬜
  🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  🟥🟥🟥🟥🟥⬜⬜⬜⬛⬛⬜⬜⬛⬜⬜⬜⬛⬛⬛⬜⬛⬛⬛⬜⬛⬛⬛⬜⬜⬜⬜⬜⬜
  🟥🟥🟥🟥🟥⬜⬜⬛⬜⬜⬛⬜⬛⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬛⬜⬜⬜⬜⬜
  🟥🟥🟥⬜🟥⬜⬜⬛⬜⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬛⬜⬜⬜⬜⬜
  🟥🟥🟥🟥⬜⬜⬜⬜⬛⬛⬜⬜⬛⬜⬜⬜⬛⬛⬜⬜⬛⬛⬜⬜⬛⬛⬛⬜⬜⬜⬜⬜⬜
  🟥🟥🟥🟥⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜
  🟥⬜🟥🟥🟥⬜⬜⬛⬜⬜⬛⬜⬛⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜
  🟥⬜🟥🟥⬜🟥🟥⬜⬛⬛⬜⬜⬛⬛⬛⬜⬛⬛⬛⬜⬛⬛⬛⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜
  🟥⬜⬜🟥🟥⬜⬜🟥⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
  🟥⬜⬜🟥⬜🟥🟥⬜⬜⬜⬛⬛⬜⬜⬜⬛⬛⬜⬜⬜⬛⬛⬜⬜⬛⬜⬛⬜⬜⬛⬜⬜⬜
  🟥⬜⬜⬜🟥⬜⬜🟥⬜⬛⬜⬜⬛⬜⬛⬜⬜⬛⬜⬛⬜⬜⬛⬜⬛⬜⬛⬛⬜⬛⬜⬜⬜
  ⬜🟥⬜⬜⬜🟥⬜⬜⬜⬛⬜⬜⬛⬜⬛⬜⬜⬜⬜⬛⬜⬜⬛⬜⬛⬜⬛⬛⬛⬛⬜⬜⬜
  ⬜🟥⬜⬜⬜⬜⬜⬜⬜⬛⬛⬛⬛⬜⬛⬜⬛⬛⬜⬛⬛⬛⬛⬜⬛⬜⬛⬜⬛⬛⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜⬛⬜⬛⬜⬜⬛⬜⬛⬜⬜⬛⬜⬛⬜⬛⬜⬜⬛⬜⬜⬜
  ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜⬛⬜⬜⬛⬛⬜⬜⬛⬜⬜⬛⬜⬛⬜⬛⬜⬜⬛⬜⬜⬜`)
	fmt.Println()
	fmt.Println()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	kruegerFile, err := ioutil.ReadFile("krueger.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config cmd.Config
	if err := yaml.Unmarshal(kruegerFile, &config); err != nil {
		log.Fatal(err)
	}

	go cmd.FetchApps(config)

	<-ch
	log.Println("shutdown...")
}
