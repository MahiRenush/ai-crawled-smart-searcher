#Runs react server 
run-ui:
	cd frontend && npm start

# Requesting you to first, run this by make store-data which would upsert data in Bolt DB and then run-server.
store-data:
	cd server; go run ../data-handler/write.go

# But before any writes to store data in "event.db" file of Bolt DB, update the static information there manually after every run.

# Runs backend server
run-server:
	cd server; go run main.go
