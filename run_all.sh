#!/bin/bash

# Start Order Service (Django)
echo "Starting Order Service..."
cd order-service && source .venv/bin/activate && uvicorn order_service.asgi:application --host 0.0.0.0 --port 8000 &
ORDER_PID=$!
deactivate
cd ..

# Start Shipping Service (Go)
echo "Starting Shipping Service..."
cd shipping-service && go run main.go &
SHIPPING_PID=$!
cd ..

# Start Inventory Service (Node/NestJS)
echo "Starting Inventory Service..."
cd inventory-service && npm run start &
INVENTORY_PID=$!
cd ..

# Start Orchestration Service (FastAPI)
echo "Starting Orchestration Service..."
cd orchestration-service && source .venv/bin/activate && uvicorn main:app --reload --port 7000 &
ORCH_PID=$!
deactivate
cd ..

# Wait for all
wait $ORDER_PID $SHIPPING_PID $INVENTORY_PID $ORCH_PID 