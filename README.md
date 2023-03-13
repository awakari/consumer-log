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
  consumer.Service/SubmitBatch
```
Payload is basically a Cloud Event:
```json
{
  "msgs": [
    {
      "id": "3426d090-1b8a-4a09-ac9c-41f2de24d5ac",
      "type": "example.type",
      "source": "example/uri",
      "spec_version": "1.0",
      "attributes": {
        "subject": {
          "ce_string": "Obi-Wan Kenobi"
        },
        "time": {
          "ce_timestamp": "1985-04-12T23:20:50.52Z"
        }
      },
      "text_data": "I felt a great disturbance in the force"
    },
    {
      "id": "3426d090-1b8a-4a09-ac9c-41f2de24d5ad",
      "type": "example.type",
      "source": "example/uri",
      "spec_version": "1.0",
      "attributes": {
        "subject": {
          "ce_string": "Yoda"
        },
        "time": {
          "ce_timestamp": "1985-05-11T12:02:05.25Z"
        }
      },
      "text_data": "Try not. Do or do not. There is no try."
    },
    {
      "id": "3426d090-1b8a-4a09-ac9c-41f2de24d5ae",
      "type": "example.type",
      "source": "example/uri",
      "spec_version": "1.0",
      "attributes": {
        "subject": {
          "ce_string": "Qui-Gon Jinn"
        },
        "time": {
          "ce_timestamp": "1985-06-08T14:31:41.16Z"
        }
      },
      "text_data": "The ability to speak does not make you intelligent."
    }
  ]
}
```
