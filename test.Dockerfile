FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY *.go ./
COPY .env ./
RUN go mod tidy

# Build
RUN go build -o /aldebaran-build

EXPOSE 8080


# Run
CMD [ "/aldebaran-build" ]
