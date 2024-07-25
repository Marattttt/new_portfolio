FROM marattttt/new_portfolio:base

WORKDIR /app/src

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/server ./cmd/server/main.go

WORKDIR /app/runtime

CMD [ "/app/server" ]

