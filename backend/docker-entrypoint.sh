echo "start" >> /build/output.log
sleep 10
go run user/rpc/user.go -f user/rpc/etc/user.yaml >> /build/output.log 2>&1 &
sleep 10
go run feed/rpc/feed.go -f feed/rpc/etc/feed.yaml >> /build/output.log 2>&1 &
sleep 10
go run favorite/rpc/favorite.go -f favorite/rpc/etc/favorite.yaml >> /build/output.log 2>&1 &
sleep 10
go run relation/rpc/relation.go -f relation/rpc/etc/relation.yaml >> /build/output.log 2>&1 &
sleep 10
go run comment/rpc/comment.go -f comment/rpc/etc/comment.yaml >> /build/output.log 2>&1 &
sleep 10
go run user/api/user.go -f user/api/etc/user.yaml >> /build/output.log 2>&1 &
sleep 10
go run feed/api/feed.go -f feed/api/etc/feed-api.yaml >> /build/output.log 2>&1 &
sleep 10
go run favorite/api/favorite.go -f favorite/api/etc/favorite-api.yaml >> /build/output.log 2>&1 &
sleep 10
go run relation/api/relation.go -f relation/api/etc/relation.yaml >> /build/output.log 2>&1 &
sleep 10
go run comment/api/comment.go -f comment/api/etc/comment-api.yaml >> /build/output.log 2>&1 &

tail -f /build/output.log
