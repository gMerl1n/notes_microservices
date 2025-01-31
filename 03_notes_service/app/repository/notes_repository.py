from abc import ABC, abstractmethod
from sqlalchemy import select, and_, update, delete
from repository.notes_models import Note, Category
from sqlalchemy.ext.asyncio import AsyncSession



class INoteRepository(ABC):

    @abstractmethod
    async def get_note_by_id(self, async_session: AsyncSession, note_id: int) -> Note | None:
        raise NotImplemented

    @abstractmethod
    async def get_all_notes(self, async_session: AsyncSession, user_id: int) -> list[Note]:
        raise NotImplemented

    @abstractmethod
    async def save_note(self, async_session: AsyncSession, note: dict) -> int:
        raise NotImplemented

    @abstractmethod
    async def remove_note_by_id(self, async_session: AsyncSession, note_id: int) -> int:
        raise NotImplemented

    @abstractmethod
    async def remove_all_notes(self, async_session: AsyncSession, user_id: int) -> list[int]:
        raise NotImplemented


class NoteRepository(INoteRepository):

    async def get_note_by_id(self, async_session: AsyncSession, note_id: int) -> Note | None:
        query = select(Note).where(Note.note_id == note_id)
        note = await async_session.execute(query)
        if note is not None:
            return note.fetchone()

    async def get_all_notes(self, async_session: AsyncSession, user_id: int):
        pass

    async def save_note(self, async_session: AsyncSession, note: dict) -> int:
        new_note = Note(**note)
        async_session.add(new_note)
        await async_session.commit()
        await async_session.refresh(new_note)
        return new_note.note_id

    async def remove_note_by_id(self, async_session: AsyncSession, note_id: int) -> int:
        pass

    async def remove_all_notes(self, async_session: AsyncSession, user_id: int) -> list[int]:
        pass