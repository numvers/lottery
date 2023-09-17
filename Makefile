build:
	@ echo "building application in path /cmd/http/app"
	@ go build -C cmd/http -o app
	@ echo "build done"

test: 
	@ go test -coverprofile coverage.out -v ./...
	@ go test -json ./... > report.json
