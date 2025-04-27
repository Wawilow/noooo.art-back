FROM golang:alpine AS builder
LABEL mantainer="wawilow"

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
ENV CGO_ENABLED=0
RUN --mount=type=cache,target="/root/.cache/go-build" go build -ldflags="-w -s" -o build ./main.go

FROM golang:alpine AS runner
WORKDIR /app
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
COPY --from=builder /app/build .
USER appuser
EXPOSE 3000

ENTRYPOINT ["./build"]
