FROM golang:1.21 as builder

LABEL maintainer="Nam HQ <namhq.1989@gmail.com>"

RUN mkdir -p /go/src/github.com/namhq1989/maids-bots
WORKDIR /go/src/github.com/namhq1989/maids-bots

COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy data to working dir
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o ./cmd/main ./cmd/main.go

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add tzdata zip ca-certificates

WORKDIR /go/src/github.com/namhq1989/maids-bots

COPY --from=builder /go/src/github.com/namhq1989/maids-bots .

# Command to run the executable
CMD ["./cmd/main"]