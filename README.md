## Phase 1
To run main.go with hot reload
```
cd cmd/hash
```
```
reflex -r '\.go' -s -- sh -c "go run main.go"
```
By using reflex, whenever the `.go` file is change, it will run main.go again.
<hr>

## Phase 2
To run with docker
```
cd cmd/hash
```
```
docker compose up
```