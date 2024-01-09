package handlers

import (
	"strconv"

	"github.com/AaronSaikovski/golotteryservice/internal/pkg/lottery"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// @Summary Lottery number generator
// @Description Generates the lottery numbers and returns JSON - The request responds to a url matching:  /lottery?maxgames=XX&randnum=XX&maxnumspergame=XX
// @Tags Lottery
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /lottery [get]
func HandleGetLotteryNumbers(c *fiber.Ctx) error {

	log.Debug().Msg("calling HandleGetLotteryNumbers()")

	//get input params
	maxgames, err := strconv.Atoi(c.Query("maxgames"))
	if err != nil {
		panic(err)
	}

	randnum, err := strconv.Atoi(c.Query("randnum"))
	if err != nil {
		panic(err)
	}

	maxnumspergame, err := strconv.Atoi(c.Query("maxnumspergame"))
	if err != nil {
		panic(err)
	}

	// Get the Generated lottery games from the package
	gameResults := lottery.GenerateLotteryGames(maxgames, randnum, maxnumspergame)

	// Return JSON results
	return c.Status(200).JSON(gameResults)
}
