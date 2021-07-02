package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testfreelancer/src"

	"github.com/labstack/echo"
)

func call(c echo.Context, d *src.Dispatcher) {

}
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		cancel()
	}()
	e := echo.New()
	p := src.NewPrinter()
	d := src.NewDispatcher(p, 10, 1000)
	d.Start(ctx)
	// Routes
	handler := func(c echo.Context) error {
		return src.CreateJob(c, d)
	}
	e.POST("/task", handler)
	e.Logger.Fatal(e.Start(":8000"))
	d.Wait()

}
