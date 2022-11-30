package handlers
import (
    "go-echo-1/models"
    "github.com/labstack/echo"
    "net/http"
)

// DeleteOrgsTarget - Delete target by target ID
func (c *Container) DeleteOrgsTarget(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}


// GetOrgsTarget - Get target by org ID
func (c *Container) GetOrgsTarget(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}


// GetOrgsTargets - Get targets by org ID
func (c *Container) GetOrgsTargets(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}
