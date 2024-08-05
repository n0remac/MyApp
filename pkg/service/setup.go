package service

import (
	"fmt"
	"net/http"
	"MyApp/gen/proto/product/productconnect"
	"MyApp/gen/proto/sensor/sensorconnect"

	"github.com/bufbuild/connect-go"
)

func SetupServices(apiRoot *http.ServeMux, interceptors connect.Option) []string {
	fmt.Println("Setting up services")
	productService := &ProductService{}
	apiRoot.Handle(productconnect.NewProductServiceHandler(productService, interceptors))
	sensorService := &SensorService{}
	apiRoot.Handle(sensorconnect.NewSensorServiceHandler(sensorService, interceptors))

	reflectorServices := []string{
		"products.productService",
		"sensors.sensorService",
	}

	return reflectorServices
}
