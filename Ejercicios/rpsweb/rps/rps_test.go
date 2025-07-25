package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)
		switch i {
		case 0:
			fmt.Println("EL jugador eligio piedra")
		case 1:
			fmt.Println("EL jugador eligio papel")
		case 2:
			fmt.Println("EL jugador eligio tijera")
		}
		fmt.Printf("mesage: %s\n", round.Message)
		fmt.Printf("Cp: %s\n", round.ComputerChoice)
		fmt.Printf("RoundResult: %s\n", round.RoundResult)
		fmt.Printf("CpInt: %d\n", round.ComputerChoiceInt)
		fmt.Printf("CpScore: %s\n", round.ComputerScore)
		fmt.Printf("Playercore: %s\n", round.PlayerScore)

		fmt.Println("*******************************************************")
	}

}
