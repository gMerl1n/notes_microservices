from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import JSONResponse
from sqlalchemy.ext.asyncio import AsyncSession
from services.category_services import ICategoryService
from container.container import container
from settings.async_session import get_async_session
from handlers.schema_request import (
    CategoryCreateRequest,
    CategoryGetRequestById,
    CategoryRemoveRequestById,
    CategoriesGetRequest
)


router_categories = APIRouter()


@router_categories.post("/create_category")
async def create_category(category_create_request: CategoryCreateRequest,
                          async_session: AsyncSession = Depends(get_async_session),
                          category_service: ICategoryService = Depends(container.get_category_service)) -> JSONResponse:

    category_id = await category_service.create_category(async_session=async_session,
                                                         category_name=category_create_request.category_name,
                                                         user_id=category_create_request.user_id)

    if category_id is None:
        raise HTTPException(status_code=500, detail="Something went wrong")

    return JSONResponse(content=category_id, status_code=201)


@router_categories.post("/get_category")
async def get_category(category_get_request: CategoryGetRequestById,
                       async_session: AsyncSession = Depends(get_async_session),
                       category_service: ICategoryService = Depends(container.get_category_service)) -> JSONResponse:

    category = await category_service.get_category_by_id(async_session=async_session,
                                                         category_id=category_get_request.category_id,
                                                         user_id=category_get_request.user_id)

    if category is None:
        raise HTTPException(status_code=400, detail="Not found")

    return JSONResponse(content=category.to_dict(), status_code=200)


@router_categories.post("/get_all_categories")
async def get_categories(categories_get_request: CategoriesGetRequest,
                         async_session: AsyncSession = Depends(get_async_session),
                         category_service: ICategoryService = Depends(container.get_category_service)) -> JSONResponse:

    categories = await category_service.get_categories(async_session=async_session,
                                                       user_id=categories_get_request.user_id)
    if categories is None:
        raise HTTPException(status_code=400, detail="Not found")

    return JSONResponse(content=[c.to_dict() for c in categories])


@router_categories.delete("/remove_category_by_id")
async def remove_category_by_id(category_remove_request: CategoryRemoveRequestById,
                                async_session: AsyncSession = Depends(get_async_session),
                                category_service: ICategoryService = Depends(container.get_category_service)) -> JSONResponse:

    removed_category_and_note_id = await category_service.remove_category_by_id(async_session=async_session,
                                                                                category_id=category_remove_request.category_id,
                                                                                user_id=category_remove_request.user_id)
    if removed_category_and_note_id is None:
        raise HTTPException(status_code=400, detail="Not found")

    return JSONResponse(content=removed_category_and_note_id, status_code=200)
