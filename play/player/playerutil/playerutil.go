package playerutil

import (
	"fmt"
	"math/rand"
	"strings"
)

var playerNames = []string{"Spartak", "Atilla", "Hercules", "Otello", "Chegevara", "Stiven", "Ninzia", "Mask", "Jeff", "Robocop", "Gudwin", "Camel", "Adic", "Roxana", "Oboltus", "Leon", "Vook", "Mosquito",
	"Bear", "Uncle", "Tor", "Micky", "Godzilla", "Cowboy", "Rembo", "Alien", "Predator", "Pumkin", "Supermen", "Leon", "Tiger", "Cobra", "Sun", "Ferum", "Genesis", "Diktator", "Alcheymer", "AlKaida", "Donkihot",
	"Terminator", "Lucky", "Nicotine", "Guru"}

var playerPrefix = []string{"blind", "big", "tiny", "sneaky", "dirty", "famous", "bold", "shiny", "vigilant", "special", "angry", "cheap", "slow", "dark", "silly", "smart", "cruel", "singing",
	"fast", "goofy", "quirky", "clumsy", "dopey", "inept", "bumbling"}

// Generating funny player names
// Pattern is big-name-123 ends always with 3 digits

func GenerateRandomName() string {
	prefix := strings.ToLower(playerPrefix[rand.Intn(len(playerPrefix))])
	return fmt.Sprintf("%s-%s-%d", prefix, strings.ToLower(playerNames[rand.Intn(len(playerNames))]), (rand.Intn(900) + 100))
}
