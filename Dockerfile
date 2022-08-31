## Build (About 500mb image)
FROM golang:1.18-buster AS build

WORKDIR /app

# COPY go.mod .
# COPY go.sum .
# COPY *.go ./
COPY ./ ./

RUN go mod tidy

RUN go build -o /aldebaran-build

## Deploy (Reduced to 20mb or so)
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /aldebaran-build /aldebaran-build

EXPOSE 8080
COPY .env ./

USER nonroot:nonroot

ENTRYPOINT ["/aldebaran-build"]
