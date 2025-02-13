import logging
from abc import ABC, abstractmethod
from datetime import datetime
from repository.notes_repository import INoteRepository
from repository.category_repository import ICategoryRepository
from sqlalchemy.ext.asyncio import AsyncSession
from domain.domain import NoteEntity

logging.basicConfig(
    format='%(asctime)s - %(message)s',
    datefmt='%d-%b-%y %H:%M:%S',
    level=logging.INFO
)


class INoteService(ABC):

    @abstractmethod
    async def get_note_by_id(self, async_session: AsyncSession, note_id: int, user_id: int) -> NoteEntity | None:
        pass

    @abstractmethod
    async def get_all_notes(self, async_session: AsyncSession, user_id: int) -> list[NoteEntity] | None:
        raise NotImplemented

    @abstractmethod
    async def save_note(self,
                        async_session: AsyncSession,
                        title: str,
                        body: str,
                        category_name: str,
                        user_id: int) -> int:
        raise NotImplemented

    @abstractmethod
    async def remove_note_by_id(self, async_session: AsyncSession, note_id: int, user_id: int) -> int | None:
        raise NotImplemented

    @abstractmethod
    async def remove_all_notes(self, async_session: AsyncSession, user_id: int) -> list[int] | None:
        raise NotImplemented


class NoteService(INoteService):

    def __init__(self, notes_repo: INoteRepository, categories_repo: ICategoryRepository) -> None:
        self.__notes_repo = notes_repo
        self.__categories_repo = categories_repo

    async def get_note_by_id(self, async_session: AsyncSession, note_id: int, user_id: int) -> NoteEntity | None:
        note = await self.__notes_repo.get_note_by_id(async_session=async_session,
                                                      note_id=note_id,
                                                      user_id=user_id)
        return note

    async def get_all_notes(self, async_session: AsyncSession, user_id: int) -> list[NoteEntity] | None:
        notes = await self.__notes_repo.get_all_notes(async_session=async_session, user_id=user_id)
        return notes

    async def save_note(self,
                        async_session: AsyncSession,
                        title: str,
                        body: str,
                        category_name: str,
                        user_id: int) -> int | None:
        category_id = await self.__categories_repo.get_category_id_by_name(async_session=async_session,
                                                                           category_name=category_name)

        if category_id is None:
            logging.warning(f"User ID: {user_id} "
                            f"Category with such a name does not exist: {category_name}. "
                            f"Impossible to create a new note")
            return

        new_note = NoteEntity(
            category_id=category_id,
            user_id=user_id,
            title=title,
            body=body,
            update_at=int(datetime.now().timestamp()),
            created_at=int(datetime.now().timestamp()),
        )

        id_note = await self.__notes_repo.save_note(async_session=async_session, note=new_note)
        return id_note

    async def remove_note_by_id(self, async_session: AsyncSession, note_id: int, user_id: int) -> int | None:
        removed_note_id = await self.__notes_repo.remove_note_by_id(async_session=async_session,
                                                                    note_id=note_id,
                                                                    user_id=user_id)
        return removed_note_id

    async def remove_all_notes(self, async_session: AsyncSession, user_id: int) -> list[int] | None:
        removed_notes_ids = await self.__notes_repo.remove_all_notes(async_session=async_session, user_id=user_id)
        return removed_notes_ids
