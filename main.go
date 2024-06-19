package main

import (
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Domains18/SIL-backend/conf"
	"github.com/Domains18/SIL-backend/internal/core/repositories"
	"github.com/Domains18/SIL-backend/internal/routes"
	authenticator "github.com/Domains18/SIL-backend/pkg/auth"
	pkg "github.com/Domains18/SIL-backend/pkg/db"
	"github.com/Domains18/SIL-backend/pkg/resolvers"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

const (
	service    = "backend"
	enviroment = "dev"
	id         = 1
)

func init() {
	gob.Register(map[string]interface{}{})
}

func main() {
	err := godotenv.Load(".env_example")
	if err != nil {
		log.Fatal("Error loading .env_example file")
	}

	tp, err := tracerProvider("http://localhost:4317/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatalf("failed to shutdown trace provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)

	configPath := resolvers.GetConfigPath(os.Getenv("config"))
	cfgFile, err := conf.RequireConfigurations(configPath)
	if err != nil {
		log.Fatalf("unable to read configurations: %v", err)
	}
	_, err = conf.ParseConfigurations(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	orderRepo := repositories.NewOrderRepo(pkg.DB)
	customerRepo := repositories.NewCustomerRepo(pkg.DB)

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
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		log.Fatal(err)
	}
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			semconv.ServiceInstanceIDKey.Int(id),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}