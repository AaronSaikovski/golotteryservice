package lottery

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// GenerateLotteryGames - Generates a number of lottery games for a given range and selection of numbers
func GenerateLotteryGames(MaxNumberGames int, RandomNumbersPerGame int, NumbersPerGame int) []string {

	var strGameTmp string
	var gameResults []string

	for i := 0; i < MaxNumberGames; i++ {
		//format string
		strGameTmp = "Game: " + strconv.Itoa(i+1) + " - " + fmt.Sprintf("%v", generateRandomNumbers(RandomNumbersPerGame, NumbersPerGame))

		//add to the game results slice
		gameResults = append(gameResults, strGameTmp)
	}

	return gameResults
}

// GenerateRandomNumbers - Generates the random numbers No duplicates - returns an array of random numbers per call
func generateRandomNumbers(maxRandomNums int, numbersPerGame int) []int {

	// Create a slice to store the generated numbers
	var numbers []int

	//seed the randomiser
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate random numbers until we have numbersPerGame of them
	for len(numbers) < numbersPerGame {

		// Generate a random number between 1 and maxVal
		n := rand.Intn(maxRandomNums) + 1

		// Check if the number has already been generated
		duplicate := false
		for _, v := range numbers {
			if v == n {
				duplicate = true
				break
			}
		}

		// If the number is not a duplicate, add it to the slice
		if !duplicate {
			numbers = append(numbers, n)
		}
	}

	//return the results array
	return numbers
}
