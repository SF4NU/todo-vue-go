package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sf4nu/todo-fiber-prova-server/models"
	"github.com/sf4nu/todo-fiber-prova-server/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://ryfljouh:dhohPc3uydQ006hRhGsulapMesD4MLFd@dumbo.db.elephantsql.com/ryfljouh" //l'url con username e password per permettere al server di collegarsi al db su elephant sql
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})                                        // questo codice fornito da GORM apre le una connessione con il database postgresql grazie al driver fornite sempre da gorm
	if err != nil {                                                                                 //controlla se c'è un errore e se ne trova uno la connessione fallisce
		panic("db not connected")
	}
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Task{}) //crea automaticamente le tabelle se non sono già presenti

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
	}) //tutta questa app.Use è il cors, quindi serve ad ammettere le connessioni esterne all'api (in pratica tutte, non filtra alcuna connessione esterna)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Todo-api")
	})

	app.Get("/categories", func(c *fiber.Ctx) error {
		var categories []models.Category
		db.Find(&categories)
		return c.Status(fiber.StatusOK).JSON(categories)
	}) //questo endpoint non lo uso più in frontend ma serviva all'inizio per vedere tutte le categorie

	app.Get("/categories/:id/tasks", func(c *fiber.Ctx) error {
		categoryID := c.Params("id")

		var category models.Category
		if err := db.Preload("Tasks").First(&category, categoryID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Category not found")
		}

		return c.JSON(category.Tasks)
	}) //cerca le task relative all'id della categoria selezionata

	app.Get("/users/:userID/categories", func(c *fiber.Ctx) error {
		userID := c.Params("userID")

		var categories []models.Category
		if err := db.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Category not found")
		}

		return c.JSON(categories)
	}) //cerca le categorie relative all'id dell'user selezionato, restituisce un errore 404 se non trova niente oppure le categorie se esistono, dovrebbe restituire sempre 404 per utenti nuovi che non hanno categorie di default

	app.Post("/tasks", func(c *fiber.Ctx) error {
		var task models.Task
		if err := c.BodyParser(&task); err != nil {
			return err
		}

		if err := db.Create(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(task)
	}) //anche questo endpoint lo usavo all'inizio in fase di sviluppo, non sono sicuro se è una buona pratica toglierlo o meno quindi per ora l'ho lasciato qui

	app.Post("/registration", func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		if err := db.Where("username = ?", user.Username).First(&user).Error; err == nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Username already exists")
		}

		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Password can't be hashed")
		}

		user.Password = hashedPassword

		if err := db.Create(&user).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	}) //endpoint che serve a registrare un nuovo utente, prima controlla se esistono altri user con lo stesso nome dopodiché se non ci sono conflitti viene criptata la password con una funzione che si trova nel folder utils

	app.Post("/login", func(c *fiber.Ctx) error {
		var loginRequest models.LoginRequest
		if err := c.BodyParser(&loginRequest); err != nil {
			return err
		}

		var user models.User
		if err := db.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
		}

		if err := utils.CheckPassword(loginRequest.Password, user.Password); err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
		}

		return c.JSON(fiber.Map{"user_id": user.ID})
	}) //simile alla registrazione ma in questo caso non risulta in un conflitto trovare un altro user, la password inserita in fase di login viene confrontata con quella criptata senza essere criptata a sua volta, viene usata la funzione CompareHashAndPassword() di bcrypt

	app.Post("/categories", func(c *fiber.Ctx) error {
		var category models.Category
		if err := c.BodyParser(&category); err != nil {
			return err
		}

		if err := db.Create(&category).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(category)
	}) //endpoint che serve ad aggiungere nuove categorie, visto che l'id viene inserito nel client non c'è bisogno di fare un endpoint separato per collegarlo all'id degli utenti, utilizzo semplicemente l'id dell'utente loggato per mandarlo nel database insieme alla nuova tabella così gorm li connette in automatico

	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var task models.Task
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("task not found")
		}

		var updatedTask models.Task
		if err := c.BodyParser(&updatedTask); err != nil {
			return err
		}

		if updatedTask.Description != "" {
			task.Description = updatedTask.Description
		}
		if updatedTask.Completed != task.Completed {
			task.Completed = updatedTask.Completed
		}
		if updatedTask.LastModified != "" {
			task.LastModified = updatedTask.LastModified
		}

		if err := db.Save(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(task)
	}) //il metodo di update per modificare le task in base al loro id, controlla i 3 campi (descrizione, completato, e lastmodified) solo se trova una task con quell'id

	app.Put("/categories/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var category models.Category
		if err := db.First(&category, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("category not found")
		}

		var updatedCategory models.Category
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
	}) //endpoint per aggiornare le categorie con la stessa logica di quello delle task sopra quindi ricopio: per modificare le categorie in base al loro id, controlla i 2 campi (titolo e lastmodified), solo se trova una categoria con quell'id

	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var task models.Task
		if err := db.First(&task, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("task not found")
		}

		var category models.Category
		if err := db.First(&category, task.CategoryID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Category related not found")
		}

		if task.CategoryID != category.ID {
			return c.Status(fiber.StatusNotFound).SendString("task does not belong to the category")
		}

		if err := db.Delete(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusAccepted).SendString("task deleted")
	}) //simile agli aggiornamenti ma in questo caso rimuove la task in base all'id

	app.Delete("/categories/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var category models.Category
		if err := db.First(&category, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("category not found")
		}

		var tasks []models.Task
		if err := db.Where("category_id = ?", category.ID).Find(&tasks).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("no match")
		}

		for _, task := range tasks {
			if err := db.Delete(&task).Error; err != nil {
				return err
			}
		}

		if err := db.Delete(&category).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusAccepted).SendString("category deleted")
	}) //simile agli aggiornamenti ma in questo caso rimuove la categoria in base all'id

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port) //questa parte del port durante la fase di sviluppo non aveva lo 0.0.0.0 ma senza non riuscivo a fare il hosting sul sito
}
