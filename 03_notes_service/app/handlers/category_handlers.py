from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import JSONResponse
from services.category_services import ICategoryService
from handlers.schema_request import CategoryCreateRequest, CategoryGetRequestById, CategoryRemoveRequestById
from container.container import container
from sqlalchemy.ext.asyncio import AsyncSession
from settings.async_session import get_async_session

router_categories = APIRouter()


@router_categories.post("/create_category")
async def create_category(category_create_request: CategoryCreateRequest,
                          async_session: AsyncSession = Depends(get_async_session),
                          category_service: ICategoryService = Depends(container.get_category_service)):

    category_id = await category_service.create_category(async_session=async_session,
                                                         category_name=category_create_request.category_name,
                                                         user_id=category_create_request.user_id)

    if category_id is None:
        raise HTTPException(status_code=500, detail="Something went wrong")

    return JSONResponse(content=category_id, status_code=201)


@router_categories.get("/get_category")
async def get_category(category_get_request: CategoryGetRequestById,
                       async_session: AsyncSession = Depends(get_async_session),
                       category_service: ICategoryService = Depends(container.get_category_service)):

    category = await category_service.get_category_by_id(async_session=async_session,
                                                         category_id=category_get_request.category_id,
                                                         user_id=category_get_request.user_id)

    if category is None:
        raise HTTPException(status_code=400, detail="Not found")

    return category


@router_categories.get("/get_all_categories")
async def get_categories(user_id: int,
                         async_session: AsyncSession = Depends(get_async_session),
                         category_service: ICategoryService = Depends(container.get_category_service)):

    categories = await category_service.get_categories(async_session=async_session, user_id=user_id)
    if categories is None:
        raise HTTPException(status_code=400, detail="Not found")

    return categories


@router_categories.delete("/remove_category_by_id")
async def remove_category_by_id(category_remove_request: CategoryRemoveRequestById,
                                async_session: AsyncSession = Depends(get_async_session),
                                category_service: ICategoryService = Depends(container.get_category_service)):

    removed_category_id = await category_service.remove_category_by_id(async_session=async_session,
                                                                       category_id=category_remove_request.category_id,
                                                                       user_id=category_remove_request.user_id)
    if removed_category_id is None:
        raise HTTPException(status_code=400, detail="Not found")

    return removed_category_id
