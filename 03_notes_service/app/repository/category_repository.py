from abc import ABC, abstractmethod
from sqlalchemy import select, and_, update, delete
from repository.models import Category
from sqlalchemy.ext.asyncio import AsyncSession
from app.domain.domain import CategoryEntity


class ICategoryRepository(ABC):

    @abstractmethod
    async def create_category(self, async_session: AsyncSession, category: CategoryEntity) -> int:
        pass

    @abstractmethod
    async def get_category_by_id(self, async_session: AsyncSession, category_id: int):
        pass

    @abstractmethod
    async def get_category_id_by_name(self, async_session: AsyncSession, category_name: str) -> int | None:
        pass


class CategoryRepository(ICategoryRepository):

    async def create_category(self, async_session: AsyncSession, category: CategoryEntity) -> int:
        print(category)
        new_category = Category.to_category_model(category)
        async_session.add(new_category)
        await async_session.commit()
        await async_session.refresh(new_category)
        return new_category.id

    async def get_category_by_id(self, async_session: AsyncSession, category_id: int):
        query = select(Category).where(Category.id == category_id)
        category = await async_session.execute(query)
        if category is not None:
            return category.fetchone()

    async def get_category_id_by_name(self, async_session: AsyncSession, category_name: str) -> int | None:

        query = select(Category).where(Category.category_name == category_name)
        category = await async_session.execute(query)
        if category is not None:
            return category.scalar().id
