from abc import ABC, abstractmethod
from datetime import datetime
from sqlalchemy.ext.asyncio import AsyncSession
from repository.category_repository import ICategoryRepository
from domain.domain import CategoryEntity


class ICategoryService(ABC):

    @abstractmethod
    async def create_category(self, async_session: AsyncSession, category: dict) -> int:
        pass

    @abstractmethod
    async def get_category_by_id(self, async_session: AsyncSession, category_id: int) -> CategoryEntity:
        pass


class CategoryService(ICategoryService):

    def __init__(self, category_repo: ICategoryRepository) -> None:
        self.__category_repo = category_repo

    async def create_category(self, async_session: AsyncSession, category: dict) -> int:

        new_category = CategoryEntity(
            category_name=category["category_name"],
            user_id=category["user_id"],
            update_at=int(datetime.now().timestamp()),
            created_at=int(datetime.now().timestamp())
        )

        category_id = await self.__category_repo.create_category(async_session=async_session, category=new_category)
        return category_id

    async def get_category_by_id(self, async_session: AsyncSession, category_id: int) -> CategoryEntity:
        category = await self.__category_repo.get_category_by_id(async_session=async_session, category_id=category_id)
        return category
