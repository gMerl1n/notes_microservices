from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import JSONResponse
from services.category_services import ICategoryService
from handlers.schema_request import CategoryCreateRequest
from container.container import container
from sqlalchemy.ext.asyncio import AsyncSession
from settings.async_session import get_async_session

router_categories = APIRouter()


@router_categories.post("/")
async def create_category(category: CategoryCreateRequest,
                          async_session: AsyncSession = Depends(get_async_session),
                          category_service: ICategoryService = Depends(container.get_category_service)):

    category_id = await category_service.create_category(async_session=async_session,
                                                         category=category.model_dump())

    if category_id is None:
        raise HTTPException(status_code=500, detail="Something wrong")

    return JSONResponse(content=category_id, status_code=201)


@router_categories.get("/")
async def get_category(category_id: int):
    pass
