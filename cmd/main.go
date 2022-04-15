package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	c "github.com/bhanupbalusu/custpreordermsv002/api/controller"
	h "github.com/bhanupbalusu/custpreordermsv002/api/handler/product_details"
	a "github.com/bhanupbalusu/custpreordermsv002/domain/app"
	i "github.com/bhanupbalusu/custpreordermsv002/pkg/initialize"
	db "github.com/bhanupbalusu/custpreordermsv002/pkg/initialize/product_details"
)

func main() {
	conn := db.NewDBConnection()
	service := a.NewProductDetailsService(conn)
	controller := c.NewProductDetailsController(service)
	routes := h.NewProductDetailsRoutesHandler(controller)

	app := i.NewFiberApp()
	routes.Install(app)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8080")
		errs <- app.Listen(":8080")
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}
