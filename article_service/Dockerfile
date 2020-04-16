FROM alpine:latest

RUN mkdir /app
RUN mkdir /app/conf

WORKDIR /app

ADD ./cmd/article_service /app

CMD ./article_service