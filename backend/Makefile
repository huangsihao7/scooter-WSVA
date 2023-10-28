download:
	GOPROXY=https://goproxy.cn go mod tidy

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
	# comment
	cd comment/rpc && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/comment_rpc" comment.go
	cd comment/api && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/comment_api" comment.go
	# MQ
	cd mq/api && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/mq_api" mq.go
	# label
	cd label/rpc && CGO_ENABLED=0 go build -ldflags "-s -w" -o "../../output/label_rpc" label.go