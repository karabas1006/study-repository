FROM golang:1.25-bookworm

WORKDIR /app

COPY . .

ENV GOPROXY=https://proxy.golang.org,direct
RUN go mod tidy

CMD ["make", "service-run"]