package src

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateJob(c echo.Context, d *Dispatcher) error {
	// p := NewPrinter()
	// d := NewDispatcher(p, 10, 1000)
	job := new(Job)
	if err := c.Bind(job); err != nil {
		return err
	}

	d.Add(job)
	// d.Start(c.Request().Context())
	return c.JSON(http.StatusCreated, job)
}
