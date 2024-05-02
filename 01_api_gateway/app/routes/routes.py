from fastapi import APIRouter
from app.clients.auth.handlers import router_auth


routes = APIRouter()


routes.include_router(router=router_auth, prefix="/auth")