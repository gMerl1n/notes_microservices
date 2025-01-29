from fastapi import APIRouter
from app.handlers.notes_handlers import router_notes


routes = APIRouter()


routes.include_router(router=router_notes, prefix="/notes")