from abc import ABC, abstractmethod
from repository.notes_repository import INoteRepository
from sqlalchemy.ext.asyncio import AsyncSession


class INoteService(ABC):

    @abstractmethod
    def get_note_by_id(self, async_session: AsyncSession, note_id: int):
        pass

    @abstractmethod
    def save_note(self, async_session: AsyncSession, note: dict):
        pass


class NoteService(INoteService):

    def __init__(self, notes_repo: INoteRepository):
        self.__notes_repo = notes_repo

    def get_note_by_id(self, async_session: AsyncSession, note_id: int):
        note_by_id = self.__notes_repo.get_note_by_id(async_session=async_session,
                                                      note_id=note_id)
        if note_by_id is not None:
            return note_by_id

    def save_note(self, async_session: AsyncSession, note: dict):
        id_note = self.__notes_repo.save_note(async_session=async_session, note=note)
        return id_note