from fastapi import APIRouter
from app.handlers.notes_handlers import router_notes
from app.handlers.category_handlers import router_categories

routes = APIRouter()

routes.include_router(router=router_notes, prefix="/notes")
routes.include_router(router=router_categories, prefix="/categories")
