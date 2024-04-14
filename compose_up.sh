GOOS=linux GOARCH=arm64 go build -o main .

if docker container ls -a | grep -q go-test-app; then
   docker container rm -f app
fi

docker image rm go-test-app

docker compose up -d