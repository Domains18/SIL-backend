package main

import (
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Domains18/food-delivery/config"
	"github.com/Domains18/food-delivery/internal/core/repository"
	"github.com/Domains18/food-delivery/internal/routes"
	"github.com/Domains18/food-delivery/pkg"
	"github.com/Domains18/food-delivery/pkg/authenticator"
	"github.com/Domains18/food-delivery/pkg/utils"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

const (
	service     = "foodDelivery"
	environment = "development"
	id          = 1
)

func init() {
	gob.Register(map[string]interface{}{})
}

func main() {
	// Initialize OpenTelemetry
	tp, err := tracerProvider("http://localhost:14269/api/traces")
	if err != nil {
		log.Fatalf("Failed to initialize tracer provider: %v", err)
	}
	defer func() {
		// Shutdown the tracer provider to flush any remaining spans
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatalf("Failed to shutdown tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)

	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	_, err = config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	err = godotenv.Load(".env_example")

	if err != nil {
		log.Fatal("Error loading .env_example file")
	}

	pkg.InitDB()

	orderRepo := repository.NewOrderRepo(pkg.DB)
	customerRepo := repository.NewCustomerRepo(pkg.DB)

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := http.NewServeMux()
	routes.RegisterRoutes(rtr, orderRepo, customerRepo, auth)

	log.Print("Server listening on http://localhost:3001/")
	if err := http.ListenAndServe("0.0.0.0:3001", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}

}

func tracerProvider(url string) (*trace.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}
