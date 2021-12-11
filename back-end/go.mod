module github.com/eliabe71/vida-saudavel/back-end/backend

go 1.17

replace github.com/eliabe71/vida-saudavel/back-end/src/router => ./src

replace github.com/eliabe71/vida-saudavel/back-end/src/database/database => ./src/database

replace github.com/eliabe71/vida-saudavel/back-end/src/models => ./src/models

require (
	github.com/eliabe71/vida-saudavel/back-end/src/database/database v0.0.0-00010101000000-000000000000 // indirect
	github.com/eliabe71/vida-saudavel/back-end/src/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/eliabe71/vida-saudavel/back-end/src/router v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
)
