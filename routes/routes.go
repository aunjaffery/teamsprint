package routes

import (
	"fmt"

	"github.com/aunjaffery/teamsprint/handler"
	"github.com/aunjaffery/teamsprint/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"success": true, "msg": "App is alive and healthy!"})

	})
	// user routes
	user := app.Group("/user")
	user.Post("/signup", handler.SignUp)
	user.Post("/login", handler.Login)
	user.Get("/validation", utils.DeserializeUser, handler.TokenValidation)

	// workspace routes
	ws := app.Group("/ws")
	ws.Post("/createworkspace", utils.DeserializeUser, handler.CreateWorkspace)
	ws.Get("/findworkspace", utils.DeserializeUser, handler.FindWorkspace)
	ws.Get("/workspaceKbns", utils.DeserializeUser, handler.WorkspaceKbns)
	ws.Post("/addWsMem", utils.DeserializeUser, handler.AddWsMember)

	// kanban routes
	kbn := app.Group("/kbn")
	kbn.Post("/createKbn", utils.DeserializeUser, handler.CreateKanban)
	kbn.Get("/fetchKbns/:ws_id", utils.DeserializeUser, handler.FetchKanban)

	// card routes
	crd := app.Group("/crd")
	crd.Post("/createcard", utils.DeserializeUser, handler.CreateCard)
	crd.Get("/fetchcards/:kbn_id", utils.DeserializeUser, handler.FetchCards)
	crd.Get("/cardstatus", utils.DeserializeUser, handler.CardStatus)

	//testing routes
	ws.Get("/check", utils.DeserializeUser, func(c *fiber.Ctx) error {
		user_id, er := c.Locals("user_id").(string)
		if !er {
			return c.Status(501).JSON(fiber.Map{
				"success": false,
				"msg":     "Error! emptywe load user",
			})
		}
		user_bson, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			return c.Status(501).JSON(fiber.Map{
				"success": false,
				"msg":     "Error! cannot load user",
			})
		}
		fmt.Println("-- locals --")
		fmt.Println(user_id)
		fmt.Println(user_bson)
		fmt.Printf("%T\n", user_bson)
		fmt.Println("-- locals --")
		return c.Status(200).JSON(fiber.Map{"success": true, "msg": "checking ws"})

	})
}
