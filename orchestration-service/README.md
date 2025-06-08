# Orchestration Service

## Project Description
This project is a microservices orchestration layer built with FastAPI. It coordinates three core services—Order Service, Inventory Service, and Shipping Service—into a unified workflow. The orchestration service exposes a single API endpoint that, when called, will:
- Create a new order (via Order Service)
- Deduct or update inventory stock (via Inventory Service)
- Create a shipment (via Shipping Service)

This pattern is typical in e-commerce and logistics systems, where multiple backend services must work together to fulfill a business process. The orchestration service handles service-to-service communication, error handling, and response aggregation, providing a simple interface for clients or frontend applications.

It is suitable for:
- Demonstrating microservices orchestration patterns
- Prototyping or building real-world order fulfillment workflows
- Learning about FastAPI, service integration, and distributed system design

## Requirements
- Python 3.8+
- [uv](https://github.com/astral-sh/uv) (recommended for fast dependency management)

## Setup
Create a virtual environment (recommended):
```bash
python3 -m venv .venv
source .venv/bin/activate
```

Install dependencies with [uv](https://github.com/astral-sh/uv):
```bash
uv pip install fastapi uvicorn httpx
```

*Alternatively, you can use `pip install ...` if you don't have uv.*

## Run the server
```bash
uvicorn main:app --reload --port 7000
```

## Using the API
Send a POST request to:
```
POST http://localhost:7000/orchestrate
```
Example body:
```json
{
  "customer_name": "Alice",
  "product_id": 1,
  "quantity": 2
}
```

## Notes
- Make sure the dependent services (order, inventory, shipping) are running and accessible at the URLs configured in `main.py`.
- You can adjust the service URLs in `main.py` to fit your environment.

## API Documentation
Access Swagger UI at:
```
http://localhost:7000/docs
```
