FROM remind101/go:1.4-newrelic

COPY . /go/src/github.com/sky-uk/newrelic-go-agent

WORKDIR /go/src/github.com/sky-uk/newrelic-go-agent

RUN go-wrapper download -tags heroku ./...
RUN go-wrapper install -tags heroku ./...
