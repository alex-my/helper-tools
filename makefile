mod:
	go mod download
	go mod tidy
	go mod verify
	go mod vendor
install i:
	go install -ldflags "-s -w" ./h.go 