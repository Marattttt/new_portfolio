FROM marattttt/new_portfolio:base
# Install curl to add NodeSource repository
RUN apt-get update && apt-get install -y curl

# Add NodeSource repository and install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
	apt-get install -y nodejs

WORKDIR /app/src

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/server ./cmd/server/main.go

WORKDIR /app/runtime

USER app_user 

CMD [ "/app/server" ]

