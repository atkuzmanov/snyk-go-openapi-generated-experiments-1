package handlers
import (
    "go-echo-1/models"
    "github.com/labstack/echo"
    "net/http"
)

// GetAPIVersion - 
func (c *Container) GetAPIVersion(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}


// ListAPIVersions - 
func (c *Container) ListAPIVersions(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}
