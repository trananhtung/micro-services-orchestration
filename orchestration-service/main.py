from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import httpx

app = FastAPI()

ORDER_SERVICE_URL = "http://localhost:8000/api/orders"
INVENTORY_SERVICE_URL = "http://localhost:3000/inventory/update-stock"
SHIPPING_SERVICE_URL = "http://localhost:8080/shipments"

# Map product_id to productName (example)
PRODUCT_ID_TO_NAME = {
    1: "Product A",
    2: "Product B",
    3: "Product C"
}

class OrchestrateRequest(BaseModel):
    customer_name: str
    product_id: int
    quantity: int

class OrchestrateResponse(BaseModel):
    order_id: int
    shipment_id: int
    message: str

@app.post("/orchestrate", response_model=OrchestrateResponse)
async def orchestrate_order(req: OrchestrateRequest):
    async with httpx.AsyncClient() as client:
        # 1. Create order
        try:
            order_payload = {"customer_name": req.customer_name}
            order_resp = await client.post(ORDER_SERVICE_URL, json=order_payload, timeout=5)
            order_resp.raise_for_status()
            order = order_resp.json()
            order_id = order["id"]
        except Exception as e:
            raise HTTPException(status_code=502, detail=f"Order service error: {str(e)}")

        # 2. Inventory deduction (update stock)
        try:
            product_name = PRODUCT_ID_TO_NAME.get(req.product_id)
            if not product_name:
                raise HTTPException(status_code=400, detail=f"Unknown product_id: {req.product_id}")
            inventory_payload = {
                "productName": product_name,
                "quantity": req.quantity
            }
            inv_resp = await client.post(INVENTORY_SERVICE_URL, json=inventory_payload, timeout=5)
            inv_resp.raise_for_status()
        except HTTPException as e:
            raise e
        except Exception as e:
            raise HTTPException(status_code=502, detail=f"Inventory service error: {str(e)}")

        # 3. Create shipment
        try:
            shipment_payload = {"order_id": order_id}
            ship_resp = await client.post(SHIPPING_SERVICE_URL, json=shipment_payload, timeout=5)
            ship_resp.raise_for_status()
            shipment = ship_resp.json()
            shipment_id = shipment["id"]
        except Exception as e:
            raise HTTPException(status_code=502, detail=f"Shipping service error: {str(e)}")

        return OrchestrateResponse(
            order_id=order_id,
            shipment_id=shipment_id,
            message="Order, inventory, and shipment processed successfully."
        ) 