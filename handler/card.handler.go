package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/aunjaffery/teamsprint/config"
	"github.com/aunjaffery/teamsprint/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CardStatus(c *fiber.Ctx) error {
	// check if user is member of workspace or kanban.
	// according to visibilty of kanban...
	card := c.Query("card")
	status := c.Query("status")
	fmt.Println("quries -->", card, status)
	if card == "" || status == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Invalid request",
		})
	}
	var cardCol *mongo.Collection = config.GetColl("cards")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	card_id, err := primitive.ObjectIDFromHex(card)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Invalid kanban id",
		})
	}
	filter := bson.M{
		"_id": card_id,
	}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err = cardCol.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Cannot update card",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": false,
		"msg":     "Card updated successfully",
	})
}
func CreateCard(c *fiber.Ctx) error {
	var cardCol *mongo.Collection = config.GetColl("cards")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var crd models.CreateCard
	err := c.BodyParser(&crd)
	fmt.Println("body ->")
	fmt.Printf("%+v\n", crd)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load body",
		})
	}
	newCard := models.CreateCard{
		Title:  crd.Title,
		Kanban: crd.Kanban,
		Status: crd.Status,
	}
	fmt.Printf("%+v\n", newCard)
	_, err = cardCol.InsertOne(ctx, newCard)
	if err != nil {
		fmt.Println(err)
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Cannot create kanban",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": false,
		"msg":     "Card added successfully",
	})

}
