from abc import ABC, abstractmethod
from datetime import datetime
from sqlalchemy.ext.asyncio import AsyncSession
from repository.category_repository import ICategoryRepository
from repository.notes_repository import INoteRepository
from domain.domain import CategoryEntity


class ICategoryService(ABC):

    @abstractmethod
    async def create_category(self, async_session: AsyncSession, category_name: str, user_id: int) -> int:
        pass

    @abstractmethod
    async def get_category_by_id(self, async_session: AsyncSession, category_id: int, user_id: int) -> CategoryEntity:
        pass

    @abstractmethod
    async def get_categories(self, async_session: AsyncSession, user_id: int) -> list[CategoryEntity] | None:
        pass

    @abstractmethod
    async def remove_category_by_id(self, async_session: AsyncSession, category_id: int, user_id: int) -> dict | None:
        pass

class CategoryService(ICategoryService):

    def __init__(self, category_repo: ICategoryRepository, notes_repo: INoteRepository) -> None:
        self.__category_repo = category_repo
        self.__notes_repo = notes_repo

    async def create_category(self, async_session: AsyncSession, category_name: str, user_id: int) -> int:

        new_category = CategoryEntity(
            category_name=category_name,
            user_id=user_id,
            update_at=int(datetime.now().timestamp()),
            created_at=int(datetime.now().timestamp())
        )

        category_id = await self.__category_repo.create_category(async_session=async_session, category=new_category)
        return category_id

    async def get_category_by_id(self, async_session: AsyncSession, category_id: int, user_id: int) -> CategoryEntity:

        category = await self.__category_repo.get_category_by_id(async_session=async_session,
                                                                 category_id=category_id,
                                                                 user_id=user_id)
        return category

    async def get_categories(self, async_session: AsyncSession, user_id: int) -> list[CategoryEntity] | None:
        categories = await self.__category_repo.get_categories(async_session=async_session, user_id=user_id)
        return categories

    async def remove_category_by_id(self, async_session: AsyncSession, category_id: int, user_id: int) -> dict | None:


        removed_notes_ids = await self.__notes_repo.remove_all_notes(async_session=async_session, user_id=user_id)

        removed_category_id = await self.__category_repo.remove_category_by_id(async_session=async_session,
                                                                               category_id=category_id,
                                                                               user_id=user_id)
        if removed_category_id is None:
            return

        return {"removed_notes_ids": removed_notes_ids, "removed_category_id": removed_category_id}
