package main

import (
	"log"
	"os"
	"strconv"
	"time"

	database "ideacraft/lib"

	fiber "github.com/gofiber/fiber/v2"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	db := database.GetDatabase()
	db.Init()

	port := os.Getenv("PORT")
	//port := "8080"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := fiber.New()

	// Возвращает список идей
	app.Get("/idea", func(c *fiber.Ctx) error {
		startDateStr := c.Query("startDate")

		startDate, err := time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			startDate = time.Now()
		}

		countStr := c.Query("count")
		count, err := strconv.Atoi(countStr)
		if err != nil {
			count = 10
		}

		isFavoriteStr := c.Query("isFavorite")
		isFavorite := isFavoriteStr == "true"

		var userId *int = nil
		userIdStr := c.Query("userId")
		vuserId, err := strconv.Atoi(userIdStr)
		if err != nil {
			userId = &vuserId
		}

		ideas := db.GetIdeas(database.GetIdeasFilter{
			StartDate:  startDate,
			Count:      count,
			IsFavorite: isFavorite,
			UserId:     userId,
		})
		return c.JSON(ideas)
	})

	app.Get("/idea/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		ideas := db.GetIdeaById(id)
		return c.JSON(ideas)
	})

	app.Listen(":" + port)
}
