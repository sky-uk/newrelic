FROM remind101/go:1.4-newrelic

COPY . /go/src/github.com/sky-uk/newrelic

WORKDIR /go/src/github.com/sky-uk/newrelic

RUN go-wrapper download -tags heroku ./...
RUN go-wrapper install -tags heroku ./...
