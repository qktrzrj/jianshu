build:
	export GOOS=linux
	export GOARCH=amd64
	go build -o main ./user_service/cmd/user_service/main.go
	mv main user_service/cmd/user_service/user_service
	go build -o main article_service/cmd/article_service/main.go
	mv main article_service/cmd/article_service/article_service
	go build -o main comment_service/cmd/comment_service/main.go
	mv main comment_service/cmd/comment_service/comment_service
	go build -o main zan_service/cmd/zan_service/main.go
	mv main zan_service/cmd/zan_service/zan_service
	go build -o main gateway/cmd/gateway/main.go
	mv main gateway/cmd/gateway/gateway
run:
	docker-compose up --build