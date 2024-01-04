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

func getUserId(local interface{}) (id primitive.ObjectID, err error) {
	user_id := local.(string)
	user_bson, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return user_bson, nil
}

func WorkspaceKbns(c *fiber.Ctx) error {
	auth_id, err := getUserId(c.Locals("user_id"))
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	var wsCol *mongo.Collection = config.GetColl("workspace")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Match stage
	matchStage := bson.M{
		"$match": bson.M{
			"members": auth_id,
		},
	}

	// Lookup stage
	lookupStage := bson.M{
		"$lookup": bson.M{
			"from":         "kanban",
			"localField":   "_id",
			"foreignField": "workspace",
			"as":           "kanban",
		},
	}

	// AddFields stage
	addFieldsStage := bson.M{
		"$addFields": bson.M{
			"kanban": bson.M{
				"$filter": bson.M{
					"input": "$kanban",
					"as":    "k",
					"cond": bson.M{
						"$or": bson.A{
							bson.M{
								"$and": bson.A{
									bson.M{
										"$eq": bson.A{
											"$$k.visibility",
											"private",
										},
									},
									bson.M{
										"$in": bson.A{
											auth_id,
											"$$k.members",
										},
									},
								},
							},
							bson.M{
								"$eq": bson.A{
									"$$k.visibility",
									"workspace",
								},
							},
						},
					},
				},
			},
		},
	}

	pipeline := []bson.M{
		matchStage,
		lookupStage,
		addFieldsStage,
	}
	cur, err := wsCol.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Cannot fetch workspace",
		})
	}
	var ws []models.WsKbnRsp
	err = cur.All(ctx, &ws)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Cannot fetch workspace",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": false,
		"msg":     "Member added successfully",
		"ws":      ws,
	})
}
func AddWsMember(c *fiber.Ctx) error {
	auth_id, err := getUserId(c.Locals("user_id"))
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	var wsCol *mongo.Collection = config.GetColl("workspace")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	type addWsMember struct {
		Ws_id  primitive.ObjectID `json:"ws_id"`
		Mem_id primitive.ObjectID `json:"mem_id"`
	}
	var addWs addWsMember
	err = c.BodyParser(&addWs)
	fmt.Println("body ->", addWs)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	if addWs.Ws_id.IsZero() || addWs.Mem_id.IsZero() {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"msg":     "ws_id and mem_id are required",
		})
	}
	update := bson.M{"$addToSet": bson.M{"members": addWs.Mem_id}}
	filter := bson.M{"members": auth_id, "_id": addWs.Ws_id}
	updateRes, err := wsCol.UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		fmt.Println("no ws :)")
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! No Workspace found",
		})
	}
	if err != nil {
		fmt.Println(err)
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": false,
		"msg":     "Member added successfully",
		"data":    updateRes,
	})
}
func FindWorkspace(c *fiber.Ctx) error {
	auth_id, err := getUserId(c.Locals("user_id"))
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	fmt.Println(auth_id)
	var wsCol *mongo.Collection = config.GetColl("workspace")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var ws []models.WsRsp
	filter := bson.M{"members": auth_id}
	cursor, err := wsCol.Find(ctx, filter)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	defer cursor.Close(ctx)
	fmt.Println("--- ws ---")
	fmt.Println(ws)
	err = cursor.All(ctx, &ws)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"msg":     "Workspaces fetched successfully",
		"ws":      ws,
	})

}
func CreateWorkspace(c *fiber.Ctx) error {
	fmt.Println("-- Create workspace called --")
	auth_id, err := getUserId(c.Locals("user_id"))
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	var wsCol *mongo.Collection = config.GetColl("workspace")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var ws models.Workspace
	defer cancel()
	err = c.BodyParser(&ws)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! cannot load user",
		})
	}
	mem := []primitive.ObjectID{}
	mem = append(mem, auth_id)
	newWs := models.Workspace{
		Name:    ws.Name,
		Creator: auth_id,
		Members: mem,
	}
	_, err = wsCol.InsertOne(ctx, newWs)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{
			"success": false,
			"msg":     "Error! Cannot create workspace",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"msg":     "Workspace created successfully",
	})

}
