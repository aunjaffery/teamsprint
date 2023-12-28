package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aunjaffery/teamsprint/config"
	"github.com/aunjaffery/teamsprint/models"
	"github.com/aunjaffery/teamsprint/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TokenValidation(c *fiber.Ctx) error {
	auth_id, err := getUserId(c.Locals("user_id"))
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	var userCol *mongo.Collection = config.GetColl("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": auth_id}
	var result models.User
	err = userCol.FindOne(ctx, filter).Decode(&result)
	fmt.Printf("%+v\n", result)
	if err != nil {
		log.Println("db query error")
		log.Println(err)
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Invalid credentials",
		})
	}
	token, err := utils.Generate_JWT(result.ID.Hex())
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Invalid credentials",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"msg":     "Login successful!",
		"data":    result,
		"token":   token,
	})

}
func Login(c *fiber.Ctx) error {
	fmt.Println("-- Login user called --")
	var userCol *mongo.Collection = config.GetColl("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Login
	defer cancel()
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	filter := bson.M{"email": user.Email}
	var result models.User
	err = userCol.FindOne(ctx, filter).Decode(&result)
	fmt.Printf("%+v\n", result)
	if err != nil {
		log.Println("db query error")
		log.Println(err)
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Invalid credentials",
		})
	}
	if result.Password != user.Password {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Invalid credentials",
		})
	}
	token, err := utils.Generate_JWT(result.ID.Hex())
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Invalid credentials",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"msg":     "Login successful!",
		"data":    result,
		"token":   token,
	})

}
func SignUp(c *fiber.Ctx) error {
	fmt.Println("-- Create user called --")
	var userCol *mongo.Collection = config.GetColl("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Signup
	defer cancel()
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	newUser := models.Signup{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	fmt.Printf("%+v", user)
	fmt.Printf("%+v", newUser)
	_, err = userCol.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"msg":     "you have signup successfully",
	})

}
