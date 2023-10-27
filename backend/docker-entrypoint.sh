echo "start" >> /build/output.log
go run user/rpc/user.go -f user/rpc/etc/user.yaml >> /build/output.log 2>&1 &
go run feed/rpc/feed.go -f feed/rpc/etc/feed.yaml >> /build/output.log 2>&1 &
go run favorite/rpc/favorite.go -f favorite/rpc/etc/favorite.yaml >> /build/output.log 2>&1 &
go run relation/rpc/relation.go -f relation/rpc/etc/relation.yaml >> /build/output.log 2>&1 &
go run comment/rpc/comment.go -f comment/rpc/etc/comment.yaml >> /build/output.log 2>&1 &


cd user/api && go run user.go -f etc/user.yaml >> /build/output.log 2>&1 &
cd feed/api && go run feed.go -f etc/feed.yaml >> /build/output.log 2>&1 &
cd favorite/api && go run favorite.go -f etc/favorite.yaml >> /build/output.log 2>&1 &
cd relation/api && go run relation.go -f etc/relation.yaml >> /build/output.log 2>&1 &
cd comment/api && go run comment.go -f etc/comment.yaml >> /build/output.log 2>&1 &

tail -f /build/output.log