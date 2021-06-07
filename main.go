package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	gestures = [5]Gesture{}
)

const (
	min = 1
	max = 6
)

type Gesture struct {
	Name       string `json:"name"`
	ID         int    `json:"id"`
	WinAgainst []int  `json:"-"`
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	setUpRules()
}

func main() {
	app := fiber.New(fiber.Config{})
	app.Use(recover.New(), cors.New())

	SetupRoute(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8084"
	}
	log.Fatal(app.Listen(":" + port))
}

// Route Start
// SetupRoute func
func SetupRoute(app *fiber.App) {
	app.Get("/choices", getChoices)
	app.Get("/choice", getChoice)
	app.Post("/play", play)
	app.Post("/multiplayer", multiplayer)
}

func getChoices(c *fiber.Ctx) error {
	return c.JSON(gestures)
}
func getChoice(c *fiber.Ctx) error {
	chioce := min + rand.Intn(max-min)
	rnd := chioce - 1
	return c.JSON(gestures[rnd])
}

type PlayerRequest struct {
	ChoiceId int `json:"player"`
}

func play(c *fiber.Ctx) error {

	pr := new(PlayerRequest)
	if err := c.BodyParser(pr); err != nil {
		// log error to file or elk
		log.Println("unable to parse request", err.Error(), "main::play")
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "unable to parse request",
			"error":   err.Error(),
		})
	}
	if pr.ChoiceId > 5 || pr.ChoiceId < 1 {
		log.Println("invalid request", pr.ChoiceId, "main::play")
		return c.Status(422).JSON(fiber.Map{
			"status":  "failed",
			"message": "player choiceId cannot be greater than 5 or less than 1",
			"error":   fmt.Errorf("invalid request").Error(),
		})
	}
	cID := min + rand.Intn(max-min)

	pg := pr.ChoiceId - 1
	cg := cID - 1
	results := RPSLSEngine(gestures[pg], gestures[cg])
	return c.JSON(fiber.Map{
		"results":  results,
		"player":   pr.ChoiceId,
		"computer": cID,
	})
}

type MultiPlayerRequest struct {
	ChoiceId1 int `json:"player1"`
	ChoiceId2 int `json:"player2"`
}

func multiplayer(c *fiber.Ctx) error {

	pr := new(MultiPlayerRequest)
	if err := c.BodyParser(pr); err != nil {
		// log error to file or elk
		log.Println("unable to parse request", err.Error(), "main::play")
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "unable to parse request",
			"error":   err.Error(),
		})
	}
	if pr.ChoiceId1 > 5 || pr.ChoiceId1 < 1 || pr.ChoiceId2 > 5 || pr.ChoiceId2 < 1 {
		log.Println("invalid request", pr.ChoiceId1, pr.ChoiceId2, "main::play")
		return c.Status(422).JSON(fiber.Map{
			"status":  "failed",
			"message": "player choiceId cannot be greater than 5 or less than 1",
			"error":   fmt.Errorf("invalid request").Error(),
		})
	}
	p1g := pr.ChoiceId1 - 1 // player 1 geuesture (p1g)
	p2g := pr.ChoiceId2 - 1 // player 2 geuesture (p2g)

	results := RPSLSEngine(gestures[p1g], gestures[p2g])
	return c.JSON(fiber.Map{
		"results": results,
		"player1": pr.ChoiceId1,
		"player2": pr.ChoiceId2,
	})
}

// Route ends

func setUpRules() {
	patterns := [5]string{"rock", "paper", "scissors", "lizard", "spock"}
	winAgainst := map[int][]int{0: {3, 4}, 1: {1, 5}, 2: {3, 4}, 3: {2, 4}, 4: {3}}
	for i, gesture := range patterns {
		gestures[i] = Gesture{
			Name:       gesture,
			ID:         i + 1,
			WinAgainst: winAgainst[i],
		}
	}
}

func RPSLSEngine(player, computer Gesture) string {
	if player.ID == computer.ID {
		fmt.Printf("You played %s & the computer played %s. You tie \n", player.Name, computer.Name)
		return "tie"
	}
	switch contains(player.WinAgainst, computer.ID) {
	case true:
		fmt.Printf("You played %s & the computer played %s. You win \n", player.Name, computer.Name)
		return "win"
	default:
		fmt.Printf("You played %s & the computer played %s. You loss \n", player.Name, computer.Name)
		return "lose"
	}
}

func contains(winAgainst []int, id int) bool {
	for _, v := range winAgainst {
		if v == id {
			return true
		}
	}
	return false
}
