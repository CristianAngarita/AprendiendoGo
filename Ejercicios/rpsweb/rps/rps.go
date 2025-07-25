package rps

import (
	"math/rand/v2"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Bien hecho!",
	"Buen trabajo!",
	"Deberías comprar un boleto de lotería",
}

var loseMessages = []string{
	"Qué lastima!",
	"Intentaño de nuevo!",
	"Hoy no es tu día",
}

var drawMessages = []string{
	"Las grandes mentes piensan igual",
	"Oh intentalo de nuevo",
	"Nadie gana, pero puedes intentarlo de nuevo",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.IntN(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "Computadora eligio Piedra"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "Computadora eligio Papel"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "Computadora eligio Tijeras"
	}

	messageInt := rand.IntN(3)
	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate!"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
