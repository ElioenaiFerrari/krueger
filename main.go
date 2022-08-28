/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/ElioenaiFerrari/krueger/cmd"
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
	cmd.Execute()
}
