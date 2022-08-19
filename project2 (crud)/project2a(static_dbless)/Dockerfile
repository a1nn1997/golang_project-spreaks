FROM golang:1.18

WORKDIR /
COPY . .
RUN go get -d github.com/gorilla/mux 
RUN go mod tidy 
RUN go build -o ./out/dist 


CMD ./out/dist