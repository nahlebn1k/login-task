FROM golang
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /tasktest cmd/main.go
CMD [ "/tasktest" ]