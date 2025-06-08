# Orchestration Service

This Orchestration Service coordinates the Order Service, Inventory Service, and Shipping Service into a single workflow.

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
