FROM golang

ENV GO111MODULE=off

COPY ./source/utils /go/src/github.com/SamuelBFavarin/recipe/source/utils
COPY ./source/api /go/src/github.com/SamuelBFavarin/recipe/source/api
COPY ./source/config /go/src/github.com/SamuelBFavarin/recipe/source/config
COPY ./source/entity /go/src/github.com/SamuelBFavarin/recipe/source/entity


WORKDIR /go/src/app
COPY ./source/main.go .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]