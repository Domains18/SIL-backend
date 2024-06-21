# Docker image details
IMAGE_NAME := Domains18/foodDelivery
IMAGE_TAG := 1.0

# Docker commands
DOCKER := docker

# Prism command for mock server
PRISM := prism

# Ports
JAEGER_PORTS := -p 5776:5775/udp -p 6833:6831/udp -p 6834:6832/udp -p 5779:5778 -p 16687:16686 -p 14269:14268 -p 9412:9411
PROMETHEUS_PORT := 9090
GRAFANA_PORT := 3000

# Docker images
JAEGER_IMAGE := jaegertracing/all-in-one:latest
PROMETHEUS_IMAGE := prom/prometheus
GRAFANA_IMAGE := grafana/grafana

.PHONY: build push mockserver jaeger prom run-prom grafana run-grafana

build:
	$(DOCKER) build -t $(IMAGE_NAME):$(IMAGE_TAG) .

push:
	$(DOCKER) push $(IMAGE_NAME):$(IMAGE_TAG)

mockserver:
	$(PRISM) mock -h localhost -p 4010 docs/api.yaml

jaeger:
	$(DOCKER) run -d --name jaeger-instance \
	-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
	$(JAEGER_PORTS) \
	$(JAEGER_IMAGE)

prom:
	$(DOCKER) pull $(PROMETHEUS_IMAGE)

run-prom:
	$(DOCKER) run --name prometheus -d -p $(PROMETHEUS_PORT):$(PROMETHEUS_PORT) $(PROMETHEUS_IMAGE)

grafana:
	$(DOCKER) pull $(GRAFANA_IMAGE)

run-grafana:
	$(DOCKER) run --name grafana -d -p $(GRAFANA_PORT):$(GRAFANA_PORT) $(GRAFANA_IMAGE)