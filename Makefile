build:
	docker build -t backend:1.0 .

push:
	docker tag backend:1.0
	docker push backend:1.0

mock:
	prism mock -h localhost -p 4010 api/docs.yaml

starters:
	sudo nomad agent -dev -config=nomad.hcl

status:
	nomad status


jaegar:
	docker run -d --name jaeger \
		-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
		-p 5775:5775/udp \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 5778:5778 \
		-p 16686:16686 \
		-p 14268:14268 \
		-p 14250:14250 \
		-p 9411:9411 \
		jaegertracing/all-in-one:latest --log-level=debug


prom:
	docker pill prom/prometheus

run-prom:
	docker run --name prometheus -d -p 9000:9000 prom/prometheus

grafana:
	docker pull grafana/grafana

run-grafana:
	docker run -d --name=grafana -p 3000:3000 grafana/grafana