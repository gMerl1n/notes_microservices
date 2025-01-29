from abc import ABC, abstractmethod
from sqlalchemy import select, and_, update, delete
from repository.notes_models import Note, Category
from sqlalchemy.ext.asyncio import AsyncSession


class INoteRepository(ABC):

    @abstractmethod
    async def get_note_by_id(self, async_session: AsyncSession, id_note: int) -> int:
        raise NotImplemented

    @abstractmethod
    async def save_note(self, async_session: AsyncSession, note: dict):
        raise NotImplemented


class NoteRepository(INoteRepository):

    async def get_note_by_id(self, async_session: AsyncSession, id_note: int):
        query = select(Note).where(Note.note_id == id_note)
        note = await async_session.execute(query)
        if note is not None:
            return note.fetchone()

    async def save_note(self, async_session: AsyncSession, note: dict) -> int:
        new_note = Note(**note)
        async_session.add(new_note)
        await async_session.commit()
        await async_session.refresh(new_note)
        return new_note.note_id