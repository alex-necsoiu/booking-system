solc:
	solc --bin --abi --overwrite contracts/bookingService.sol -o build
abigen:
	abigen --pkg=bookingsystem --bin=build/BookRoom.bin --abi=build/BookRoom.abi  --out=gen/bookingSystem.go
testnet:
	go run ./deploy/$(ls -1 *.go | grep -v _tests.go)
	go run ./client/$(ls -1 *.go | grep -v _test.go)
test-client:
	gotest -v ./client/*_test.go