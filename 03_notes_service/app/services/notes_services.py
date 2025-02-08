from abc import ABC, abstractmethod
from datetime import datetime
from repository.notes_repository import INoteRepository
from repository.category_repository import ICategoryRepository
from sqlalchemy.ext.asyncio import AsyncSession
from domain.domain import NoteEntity


class INoteService(ABC):

    @abstractmethod
    async def get_note_by_id(self, async_session: AsyncSession, note_id: int) -> int | None:
        pass

    @abstractmethod
    async def save_note(self, async_session: AsyncSession, note: dict) -> int:
        pass


class NoteService(INoteService):

    def __init__(self, notes_repo: INoteRepository, categories_repo: ICategoryRepository) -> None:
        self.__notes_repo = notes_repo
        self.__categories_repo = categories_repo

    async def get_note_by_id(self, async_session: AsyncSession, note_id: int) -> NoteEntity | None:
        note = await self.__notes_repo.get_note_by_id(async_session=async_session,
                                                note_id=note_id)
        if note is not None:
            return NoteEntity(**note)

    async def save_note(self, async_session: AsyncSession, note: dict) -> int:

        category_id = await self.__categories_repo.get_category_id_by_name(async_session=async_session,
                                                                           category_name=note["category_name"])
        note.pop("category_name")
        note["category_id"] = category_id

        new_note = NoteEntity(
            category_id=category_id,
            user_id=note["user_id"],
            title=note["title"],
            body=note["body"],
            update_at=int(datetime.now().timestamp()),
            created_at=int(datetime.now().timestamp()),
        )

        id_note = await self.__notes_repo.save_note(async_session=async_session, note=new_note)
        return id_note
