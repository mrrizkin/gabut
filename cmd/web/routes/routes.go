package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/nugrhrizki/gabut/web"
)

type JSONRPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type JSONRPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
	ID      int         `json:"id"`
}

func Setup(app *fiber.App) *fiber.App {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	/**
	 * TODO: Complete JSON-RPC implementation
	 */
	rpc := v1.Group("/rpc")
	rpc.Post("/", func(ctx *fiber.Ctx) error {
		var req JSONRPCRequest

		if err := ctx.BodyParser(&req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(JSONRPCResponse{
				Jsonrpc: req.Jsonrpc,
				Result:  nil,
				Error:   err.Error(),
				ID:      req.ID,
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(JSONRPCResponse{
			Jsonrpc: req.Jsonrpc,
			Result:  "Hello World",
			Error:   nil,
			ID:      req.ID,
		})
	})

	app.Get("/*", filesystem.New(filesystem.Config{
		Root:         web.Dist(),
		NotFoundFile: "index.html",
		Index:        "index.html",
	}))

	return app
}
