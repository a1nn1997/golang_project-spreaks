FROM golang:1.18

WORKDIR /
COPY . .
RUN go get -d github.com/gorilla/mux \
    github.com/jinzhu/gorm \
    github.com/lib/pq 

CMD ["go","run","cmd/main/main.go"]