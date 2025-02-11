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
    async def get_category_by_id(self,
                                 async_session: AsyncSession,
                                 category_id: int,
                                 user_id: int) -> CategoryEntity | None:
        pass

    @abstractmethod
    async def get_category_id_by_name(self, async_session: AsyncSession, category_name: str) -> int | None:
        pass

    @abstractmethod
    async def get_categories(self, async_session: AsyncSession, user_id: int) -> list[CategoryEntity] | None:
        pass

    @abstractmethod
    async def remove_category_by_id(self, async_session: AsyncSession, category_id: int, user_id: int) -> int | None:
        pass

class CategoryRepository(ICategoryRepository):

    async def create_category(self, async_session: AsyncSession, category: CategoryEntity) -> int:

        new_category = Category.to_category_model(category)
        async_session.add(new_category)
        await async_session.commit()
        await async_session.refresh(new_category)
        return new_category.id

    async def get_category_by_id(self,
                                 async_session: AsyncSession,
                                 category_id: int,
                                 user_id: int) -> CategoryEntity | None:

        query = select(Category).where(and_(Category.id == category_id, Category.user_id == user_id))
        category_obj = await async_session.execute(query)
        if category_obj is None:
            return

        category_scalar = category_obj.scalar()
        if category_scalar is None:
            return

        return CategoryEntity(
            id=category_scalar.id,
            category_name=category_scalar.category_name,
            user_id=category_scalar.user_id,
            update_at=int(category_scalar.update_at.timestamp()),
            created_at=int(category_scalar.created_at.timestamp()),
        )

    async def get_category_id_by_name(self, async_session: AsyncSession, category_name: str) -> int | None:

        query = select(Category).where(Category.category_name == category_name)
        category = await async_session.execute(query)
        if category is None:
            return

        category_scalar = category.scalar()
        if category_scalar is None:
            return

        return category_scalar.id

    async def get_categories(self, async_session: AsyncSession, user_id: int) -> list[CategoryEntity] | None:

        result: list[CategoryEntity] = []

        query = select(Category).where(Category.user_id == user_id)
        categories = await async_session.execute(query)
        if categories is None:
            return

        for category in categories.scalars():
            result.append(
                CategoryEntity(
                    id=category.id,
                    category_name=category.category_name,
                    user_id=category.user_id,
                    update_at=int(category.update_at.timestamp()),
                    created_at=int(category.created_at.timestamp()),
                )
            )

        return result

    async def remove_category_by_id(self, async_session: AsyncSession, category_id: int, user_id: int) -> int | None:

        query = delete(Category).where(and_(Category.id == category_id, Category.user_id == user_id)).returning(Category.id)
        removed_category_id = await async_session.execute(query)
        if removed_category_id is None:
            return

        await async_session.commit()

        return removed_category_id.scalar()

