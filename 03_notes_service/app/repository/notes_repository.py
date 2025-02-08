from abc import ABC, abstractmethod
from sqlalchemy import select, and_, update, delete
from sqlalchemy.ext.asyncio import AsyncSession
from repository.models import Note
from domain.domain import NoteEntity


class INoteRepository(ABC):

    @abstractmethod
    async def get_note_by_id(self, async_session: AsyncSession, note_id: int) -> Note | None:
        raise NotImplemented

    @abstractmethod
    async def get_all_notes(self, async_session: AsyncSession, user_id: int) -> list[NoteEntity] | None:
        raise NotImplemented

    @abstractmethod
    async def save_note(self, async_session: AsyncSession, note: NoteEntity) -> int:
        raise NotImplemented

    @abstractmethod
    async def remove_note_by_id(self, async_session: AsyncSession, note_id: int) -> int:
        raise NotImplemented

    @abstractmethod
    async def remove_all_notes(self, async_session: AsyncSession, user_id: int) -> list[int]:
        raise NotImplemented


class NoteRepository(INoteRepository):

    async def get_note_by_id(self, async_session: AsyncSession, note_id: int) -> Note | None:
        query = select(Note).where(Note.id == note_id)
        note = await async_session.execute(query)
        if note is not None:
            return note.fetchone()

    async def get_all_notes(self, async_session: AsyncSession, user_id: int) -> list[NoteEntity] | None:

        result: list[NoteEntity] = []

        query = select(Note).where(Note.user_id == user_id)
        notes = await async_session.execute(query)
        if notes is not None:

            for note in notes.scalars():
                result.append(
                    NoteEntity(
                        id=note.id,
                        category_id=note.category_id,
                        user_id=note.user_id,
                        title=note.title,
                        body=note.body,
                        update_at=int(note.update_at.timestamp()),
                        created_at=int(note.update_at.timestamp()),
                    )
                )

            return result

    async def save_note(self, async_session: AsyncSession, note: NoteEntity) -> int:
        new_note = Note.to_note_model(note)
        async_session.add(new_note)
        await async_session.commit()
        await async_session.refresh(new_note)
        return new_note.id

    async def remove_note_by_id(self, async_session: AsyncSession, note_id: int) -> int:
        # TODO
        pass

    async def remove_all_notes(self, async_session: AsyncSession, user_id: int) -> list[int]:
        # TODO
        pass
