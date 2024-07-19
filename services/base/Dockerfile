FROM golang:1.22

WORKDIR /app/setup

COPY . .

RUN chmod +x ./setup.sh

RUN ./setup.sh

CMD [ "echo",  "Image has created an app_user with no password and restrictions" ]

