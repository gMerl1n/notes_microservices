from fastapi import APIRouter, Depends, HTTPException
from grpc.aio import AioRpcError
from app.clients.notes_service.protos.genproto import notes_pb2
from app.clients.auth.service import get_user_uuid_from_token


router_category = APIRouter()


@router_category.post("/")
async def create_category(user = Depends(get_user_uuid_from_token)):
    pass


@router_category.get("/listcategories")
async def get_list_categories(user = Depends(get_user_uuid_from_token)):
    pass


@router_category.put("/")
async def update_category(user = Depends(get_user_uuid_from_token)):
    pass


@router_category.delete("/")
async def delete_category(user = Depends(get_user_uuid_from_token)):
    pass