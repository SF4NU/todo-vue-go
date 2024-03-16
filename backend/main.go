package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID           uint   `gorm:"primaryKey"`
	Title        string `json:"title"`
	LastModified string `json:"last_modified"`
	Todos        []Task `gorm:"foreignKey:CategoryID"`
}
type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Deadline    string `json:"deadline"`
	CategoryID  uint   `json:"category_id"`
} //la relazione è one-to-many cioè uno a molti perché la tabella categorie è assocciata a più tabelle task

func main() {
	dsn := "postgres://ryfljouh:dhohPc3uydQ006hRhGsulapMesD4MLFd@dumbo.db.elephantsql.com/ryfljouh" //l'url con username e password per permettere al server di collegarsi al db su elephant sql
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db not connected")
	}
	db.AutoMigrate(&Category{}, &Task{}) //crea automaticamente le tabelle se non sono già presenti

	app := fiber.New() //fa partire il server

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == "OPTIONS" {
			c.Status(fiber.StatusOK)
			return nil
		}
		return c.Next()
	}) //tutta questa app.use è il cors, quindi serve ad ammettere le connessioni esterne all'api

	app.Get("/", func(c *fiber.Ctx) error { //il metodo get per ottenere la lista dei vari todo
		return c.SendString("Todo-api")
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		var tasks []Task
		db.Find(&tasks)
		return c.Status(fiber.StatusOK).JSON(tasks)
	})

	app.Get("/categories", func(c *fiber.Ctx) error {
		var categories []Category
		db.Find(&categories)
		return c.Status(fiber.StatusOK).JSON(categories)
	})

	app.Post("/tasks", func(c *fiber.Ctx) error {
		var task Task
		if err := c.BodyParser(&task); err != nil {
			return err
		}

		if err := db.Create(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(task)
	})

	app.Post("/categories", func(c *fiber.Ctx) error {
		var category Category
		if err := c.BodyParser(&category); err != nil {
			return err
		}

		if err := db.Create(&category).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(category)
	})

	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var task Task
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("task not found")
		}

		var updatedTask Task
		if err := c.BodyParser(&updatedTask); err != nil {
			return err
		}

		if updatedTask.Description != "" {
			task.Description = updatedTask.Description
		}
		if updatedTask.Completed != task.Completed {
			task.Completed = updatedTask.Completed
		}
		if updatedTask.Deadline != task.Deadline {
			task.Deadline = updatedTask.Deadline
		}

		if err := db.Save(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(task)
	})

	app.Put("/categories/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var category Category
		if err := db.First(&category, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("category not found")
		}

		var updatedCategory Category
		if err := c.BodyParser(&updatedCategory); err != nil {
			return err
		}

		if updatedCategory.Title != "" {
			category.Title = updatedCategory.Title
		}
		if updatedCategory.LastModified != "" {
			category.LastModified = updatedCategory.LastModified
		}

		if err := db.Save(&category).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(category)
	})

	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var task Task
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("task not found")
		}

		if err := db.Delete(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusAccepted).SendString("task deleted")
	})

	app.Delete("/categories/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var category Category
		if err := db.First(&category, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("category not found")
		}

		if err := db.Delete(&category).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusAccepted).SendString("category deleted")
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}
