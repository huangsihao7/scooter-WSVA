download:
	GOPROXY=https://goproxy.cn go mod tidy
doc:
	mkdir -p ../scripts/swagger/api
	GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@v1.6.0
	GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@v0.1.0
	# user
	goctl api plugin -plugin goctl-swagger="swagger -filename ../scripts/swagger/api/user.json -host 172.22.121.53:7070" -api user/api/user.api -dir .
	# feed
	goctl api plugin -plugin goctl-swagger="swagger -filename ../scripts/swagger/api/feed.json -host 172.22.121.53:7070" -api feed/api/feed.api -dir .
	# favorite
	goctl api plugin -plugin goctl-swagger="swagger -filename ../scripts/swagger/api/favorite.json -host 172.22.121.53:7070" -api favorite/api/favorite.api -dir .
	# relation
	goctl api plugin -plugin goctl-swagger="swagger -filename ../scripts/swagger/api/relation.json -host 172.22.121.53:7070" -api relation/api/relation.api -dir .
	#live
	goctl api plugin -plugin goctl-swagger="swagger -filename ../scripts/swagger/api/live.json -host 172.22.121.53:7070" -api live/api/live.api -dir .
	# comment
	goctl api plugin -plugin goctl-swagger="swagger -filename ../scripts/swagger/api/comment.json -host 172.22.121.53:7070" -api comment/api/comment.api -dir .
build:
	mkdir -p output
	# user
	cd user/rpc && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/user_rpc" user.go
	cd user/api && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/user_api" user.go
	# feed
	cd feed/rpc && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/feed_rpc" feed.go
	cd feed/api && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/feed_api" feed.go
	# favorite
	cd favorite/rpc && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/favorite_rpc" favorite.go
	cd favorite/api && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/favorite_api" favorite.go
	# relation
	cd relation/rpc && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/relation_rpc" relation.go
	cd relation/api && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/relation_api" relation.go
	# live
	cd live/rpc && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/live_rpc" live.go
	cd live/api && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/live_api" live.go
	# comment
	cd comment/rpc && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/comment_rpc" comment.go
	cd comment/api && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/comment_api" comment.go
	# MQ
	cd mq && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../output/mq_api" mq.go
	# label
	cd label/rpc && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/label_rpc" label.go