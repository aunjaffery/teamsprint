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

func FetchKbnCards(c *fiber.Ctx) error {
	// check if user is member of workspace or kanban.
	// according to visibilty of kanban...
	auth_id, err := getUserId(c.Locals("user_id"))
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	kbn_id := c.Params("kbn_id")
	fmt.Println("kbnCardsparams -->", kbn_id)
	var kbnCol *mongo.Collection = config.GetColl("kanban")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	obj_kbn_id, err := primitive.ObjectIDFromHex(kbn_id)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Invalid kanban id",
		})
	}
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"_id": obj_kbn_id,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "workspace",
				"localField":   "workspace",
				"foreignField": "_id",
				"as":           "ws",
			},
		},
		{
			"$match": bson.M{
				"$or": []bson.M{
					{
						"$and": []bson.M{
							{"visibility": "workspace"},
							{"ws.members": auth_id},
						},
					},
					{
						"$and": []bson.M{
							{"visibility": "private"},
							{"members": auth_id},
						},
					},
				},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "cards",
				"localField":   "_id",
				"foreignField": "kanban",
				"as":           "cards",
			},
		},
		{
			"$project": bson.M{
				"_id":        1,
				"title":      1,
				"lists":      1,
				"visibility": 1,
				"cards": bson.M{
					"_id":    1,
					"title":  1,
					"status": 1,
				},
			},
		},
	}
	cursor, err := kbnCol.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println(err)
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Cannot fetch kanban",
		})
	}
	kbnWithCards := models.KanbanWithCards{}
	for cursor.Next(ctx) {
		err := cursor.Decode(&kbnWithCards)
		if err != nil {
			return c.Status(501).JSON(fiber.Map{
				"success": false,
				"msg":     "Error! cannot load cards",
			})
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"success": false,
		"msg":     "Cards fetched successfully",
		"data":    kbnWithCards,
	})
}
func FetchKanban(c *fiber.Ctx) error {
	auth_id, err := getUserId(c.Locals("user_id"))
	ws_id := c.Params("ws_id")
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	fmt.Println("params -->", auth_id, ws_id)
	var kbnCol *mongo.Collection = config.GetColl("kanban")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	obj_ws_id, err := primitive.ObjectIDFromHex(ws_id)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Invalid workspace id",
		})
	}
	filter := bson.M{
		"workspace": obj_ws_id,
		"members":   auth_id,
	}
	cursor, err := kbnCol.Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Cannot fetch kanban",
		})
	}
	kbns := []models.KanbanRsp{}
	err = cursor.All(ctx, &kbns)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": false,
		"msg":     "Kanban fetched successfully",
		"data":    kbns,
	})
}

func CreateKanban(c *fiber.Ctx) error {
	time.Sleep(time.Second)
	auth_id, err := getUserId(c.Locals("user_id"))
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	var kbnCol *mongo.Collection = config.GetColl("kanban")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var kbn models.Kanban
	err = c.BodyParser(&kbn)
	fmt.Println("body ->")
	fmt.Printf("%+v\n", kbn)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load body",
		})
	}
	mem := []primitive.ObjectID{}
	mem = append(mem, auth_id)
	preMadeList := []models.List{
		{ID: primitive.NewObjectID(), Title: "todo"},
		{ID: primitive.NewObjectID(), Title: "inprogress"},
		{ID: primitive.NewObjectID(), Title: "done"},
	}
	newKbn := models.Kanban{
		Title:      kbn.Title,
		Creator:    auth_id,
		Members:    mem,
		Lists:      preMadeList,
		Visibility: kbn.Visibility,
		Workspace:  kbn.Workspace,
	}
	fmt.Printf("%+v\n", newKbn)
	// kbnCol.Find(ctx, bson.M{})
	_, err = kbnCol.InsertOne(ctx, newKbn)
	if err != nil {
		fmt.Println(err)
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Cannot create kanban",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"msg":     "Kanban added successfully",
	})
}
