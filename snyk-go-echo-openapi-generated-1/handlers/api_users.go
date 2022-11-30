package handlers
import (
    "go-echo-1/models"
    "github.com/labstack/echo"
    "net/http"
)

// GetUser - Get user by ID
func (c *Container) GetUser(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}
