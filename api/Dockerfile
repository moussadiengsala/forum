
FROM golang:1.20-alpine3.18
LABEL name="FORUM PROJECT"
LABEL description="zone01 Dakar"
LABEL authors="mousdieng; abdouazindiaye; dgaye; pka; mnom"
RUN mkdir /app
RUN apk update && apk add bash && apk add tree
RUN apk add --no-cache gcc musl-dev
ADD . /app
WORKDIR /app
RUN go build -o forum
CMD ["./forum"]
