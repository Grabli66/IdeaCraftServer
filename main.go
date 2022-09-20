package main

import (
	"log"
	"os"
	"strconv"
	"time"

	lib "ideacraft/lib"

	fiber "github.com/gofiber/fiber/v2"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	db := lib.GetDatabase()
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

		ideas := db.GetIdeas(lib.GetIdeasFilter{
			StartDate:  startDate,
			Count:      count,
			IsFavorite: isFavorite,
			UserId:     userId,
		})
		return c.JSON(ideas)
	})

	// Создаёт новую идею
	app.Post("/idea", func(c *fiber.Ctx) error {
		type AddIdeaRequest struct {
			Title       string `json:"title"`
			Desctiption string `json:"desctiption"`
			UserId      int    `json:"userid"`
		}

		req := AddIdeaRequest{}

		if err := c.BodyParser(&req); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		lib.GetDatabase().AddIdea(lib.IdeaDB{})
		return nil
	})

	// Редактирует идею
	app.Put("/idea", func(c *fiber.Ctx) error {
		lib.GetDatabase().EditIdea(lib.IdeaDB{})
		return nil
	})

	// Удаляет идею
	app.Delete("/idea", func(c *fiber.Ctx) error {
		return nil
	})

	// Возвращает идею по идентификатору
	app.Get("/idea/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		ideas := db.GetIdeaById(id)
		return c.JSON(ideas)
	})

	// Возвращает список комментариевпо идентификатору идеи
	app.Get("/comment/:ideaId", func(c *fiber.Ctx) error {
		return nil
	})

	// Создаёт новый комментарий
	app.Post("/comment", func(c *fiber.Ctx) error {
		return nil
	})

	app.Listen(":" + port)
}
