# Purpose

The consumer service implementation that just logs the incoming messages.

# Configuration

The service is configurable using the environment variables:

| Variable               | Example value | Description                                                     |
|------------------------|---------------|-----------------------------------------------------------------|
| API_PORT               | `8080`        | gRPC API port                                                   |
| LOG_LEVEL              | `-4`          | [Logging level](https://pkg.go.dev/golang.org/x/exp/slog#Level) |

# Deployment

## Local

### Build

```make build```

### Run

```shell
API_PORT=8083 ./consumer-log
```

## Docker

```shell
make run
```

## K8s

Create a helm package from the sources:
```shell
helm package helm/consumer-log/
```

Install the helm chart:
```shell
helm install consumer-log ./consumer-log-<CHART_VERSION>.tgz
```

where
* `<CHART_VERSION>` is the helm chart version

# Usage

The service provides a gRPC API for routing a message.

Example command:
```shell
grpcurl \
  -plaintext \
  -proto api/grpc/service.proto \
  -d @ \
  localhost:8080 \
  consumer.Service/Submit
```
Payload is basically a Cloud Event:
```json
{
  "id": "3426d090-1b8a-4a09-ac9c-41f2de24d5ac",
  "type": "example.type",
  "source": "example/uri",
  "spec_version": "1.0",
  "attributes": {
    "awakarisubscription": {
      "ce_string": "f7102c87-3ce4-4bb0-8527-b4644f685b13"
    },
    "awakaridestination": {
      "ce_string": "starwars"
    },
    "subject": {
      "ce_string": "Obi-Wan Kenobi"
    },
    "time": {
      "ce_timestamp": "1985-04-12T23:20:50.52Z"
    }
  },
  "text_data": "I felt a great disturbance in the force"
}
```
