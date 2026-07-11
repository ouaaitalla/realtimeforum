#!/bin/bash

# Start Go backend
go run backend/main.go &
GO_PID=$!

echo "Go backend started with PID: $GO_PID"

# Start Node server
node frontend/server.js &
NODE_PID=$!

echo "Node server started with PID: $NODE_PID"

echo "========================="
echo "Go PID   : $GO_PID"
echo "Node PID : $NODE_PID"
echo "========================="

# Wait for both processes
wait $GO_PID $NODE_PID