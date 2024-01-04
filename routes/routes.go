package routes

import (
	"github.com/aunjaffery/teamsprint/handler"
	"github.com/aunjaffery/teamsprint/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {

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
	//made this route for fetching cards with kbn id. after check visibility
	kbn.Get("/fetchKbnCards/:kbn_id", utils.DeserializeUser, handler.FetchKbnCards)

	// card routes
	crd := app.Group("/crd")
	crd.Post("/createcard", utils.DeserializeUser, handler.CreateCard)
	crd.Get("/cardstatus", utils.DeserializeUser, handler.CardStatus)
}
