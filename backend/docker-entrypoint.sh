echo "start" >> /build/output.log
cd user/rpc && go run user.go -f etc/user.yaml >> /build/output.log 2>&1 &
cd feed/rpc && go run feed.go -f etc/feed.yaml >> /build/output.log 2>&1 &
cd favorite/rpc && go run favorite.go -f etc/favorite.yaml >> /build/output.log 2>&1 &
cd relation/rpc && go run relation.go -f etc/relation.yaml >> /build/output.log 2>&1 &
cd comment/rpc && go run comment.go -f etc/comment.yaml >> /build/output.log 2>&1 &


cd user/api && go run user.go -f etc/user.yaml >> /build/output.log 2>&1 &
cd feed/api && go run feed.go -f etc/feed.yaml >> /build/output.log 2>&1 &
cd favorite/api && go run favorite.go -f etc/favorite.yaml >> /build/output.log 2>&1 &
cd relation/api && go run relation.go -f etc/relation.yaml >> /build/output.log 2>&1 &
cd comment/api && go run comment.go -f etc/comment.yaml >> /build/output.log 2>&1 &

tail -f output.log