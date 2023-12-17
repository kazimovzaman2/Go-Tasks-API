package controllers

import (
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"github.com/kazimovzaman2/tasks-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTodos(ctx *fiber.Ctx) {
	collection := mgm.Coll(&models.Todo{})
	todos := []models.Todo{}
	err := collection.SimpleFind(&todos, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})
}

func GetTodoByID(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

func CreateTodo(ctx *fiber.Ctx) {
	params := new(struct {
		Title       string
		Description string
	})

	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title or description not specified.",
		})
		return
	}

	todo := models.CreateTodo(params.Title, params.Description)
	err := mgm.Coll(todo).Create(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

func ToggleTodoStatus(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	todo.Done = !todo.Done

	err = collection.Update(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

func DeleteTodo(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	err = collection.Delete(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}
