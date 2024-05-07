from fastapi import APIRouter
from app.clients.auth.handlers import router_auth
from app.clients.notes_service.notes.handlers import router_notes
from app.clients.notes_service.category.handlers import router_category


routes = APIRouter()


routes.include_router(router=router_auth, prefix="/auth", tags=["auth"])
routes.include_router(router=router_notes, prefix="/notes", tags=["notes"])
routes.include_router(router=router_category, prefix="/categories", tags=["categories"])