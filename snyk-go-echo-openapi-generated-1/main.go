package main

import (
	"go-echo-1/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// GetAPIVersion - 
	e.GET("/rest/openapi/:version", c.GetAPIVersion)

	// ListAPIVersions - 
	e.GET("/rest/openapi", c.ListAPIVersions)

	// GetOrgsProjects - Get projects by org ID.
	e.GET("/rest/orgs/:org_id/projects", c.GetOrgsProjects)

	// DeleteOrgsTarget - Delete target by target ID
	e.DELETE("/rest/orgs/:org_id/targets/:target_id", c.DeleteOrgsTarget)

	// GetOrgsTarget - Get target by org ID
	e.GET("/rest/orgs/:org_id/targets/:target_id", c.GetOrgsTarget)

	// GetOrgsTargets - Get targets by org ID
	e.GET("/rest/orgs/:org_id/targets", c.GetOrgsTargets)

	// GetUser - Get user by ID
	e.GET("/rest/orgs/:org_id/users/:id", c.GetUser)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}