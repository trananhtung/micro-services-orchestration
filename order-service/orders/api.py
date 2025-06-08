from ninja import Router, Schema
from typing import List
from .models import Order
from django.shortcuts import get_object_or_404
from datetime import datetime

router = Router()

class OrderIn(Schema):
    customer_name: str
    status: str = "pending"

class OrderOut(Schema):
    id: int
    customer_name: str
    status: str
    created_at: datetime
    updated_at: datetime

@router.post("/orders", response=OrderOut)
def create_order(request, data: OrderIn):
    order = Order.objects.create(**data.dict())
    return order

@router.get("/orders", response=List[OrderOut])
def list_orders(request):
    return list(Order.objects.all())

@router.get("/orders/{order_id}", response=OrderOut)
def get_order(request, order_id: int):
    order = get_object_or_404(Order, id=order_id)
    return order

@router.put("/orders/{order_id}", response=OrderOut)
def update_order(request, order_id: int, data: OrderIn):
    order = get_object_or_404(Order, id=order_id)
    for attr, value in data.dict().items():
        setattr(order, attr, value)
    order.save()
    return order 