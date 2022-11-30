package handlers
import (
    "go-echo-1/models"
    "github.com/labstack/echo"
    "net/http"
)

// GetOrgsProjects - Get projects by org ID.
func (c *Container) GetOrgsProjects(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}
