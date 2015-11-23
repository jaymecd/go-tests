default:

myapp:
	go build -i -o myapp \
		-ldflags "-s \
			-X main.Version=1.0.2 \
			-X main.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` \
			-X main.GitHash=`git rev-parse HEAD`" \
		myapp.go
	./myapp
