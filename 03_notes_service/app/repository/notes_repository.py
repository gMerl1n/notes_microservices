import logging
from abc import ABC, abstractmethod
from sqlalchemy import select, and_, update, delete
from sqlalchemy.ext.asyncio import AsyncSession
from repository.models import Note
from domain.domain import NoteEntity

logging.basicConfig(
    format='%(asctime)s - %(message)s',
    datefmt='%d-%b-%y %H:%M:%S',
    level=logging.INFO
)


class INoteRepository(ABC):

    @abstractmethod
    async def get_note_by_id(self, async_session: AsyncSession, note_id: int, user_id: int) -> NoteEntity | None:
        raise NotImplemented

    @abstractmethod
    async def get_all_notes(self, async_session: AsyncSession, user_id: int) -> list[NoteEntity] | None:
        raise NotImplemented

    @abstractmethod
    async def save_note(self, async_session: AsyncSession, note: NoteEntity) -> int:
        raise NotImplemented

    @abstractmethod
    async def remove_note_by_id(self, async_session: AsyncSession, note_id: int, user_id: int) -> int | None:
        raise NotImplemented

    @abstractmethod
    async def remove_all_notes(self, async_session: AsyncSession, user_id: int) -> list[int] | None:
        raise NotImplemented


class NoteRepository(INoteRepository):

    async def get_note_by_id(self, async_session: AsyncSession, note_id: int, user_id: int) -> NoteEntity | None:
        query = select(Note).where(and_(Note.id == note_id, Note.user_id == user_id))
        note = await async_session.execute(query)
        if note is None:
            logging.warning(f"Note with user id {user_id} or note id {note_id} do not exist")
            return

        note_scalar = note.scalar()
        if note_scalar is None:
            logging.warning(f"Failed to get scalar data from query note. "
                            f"Probably, note with user id {user_id} or note id {note_id} do not exist")
            return

        return NoteEntity(
            id=note_scalar.id,
            category_id=note_scalar.category_id,
            user_id=note_scalar.user_id,
            title=note_scalar.title,
            body=note_scalar.body,
            update_at=int(note_scalar.update_at.timestamp()),
            created_at=int(note_scalar.update_at.timestamp()),
        )

    async def get_all_notes(self, async_session: AsyncSession, user_id: int) -> list[NoteEntity] | None:

        result: list[NoteEntity] = []

        query = select(Note).where(Note.user_id == user_id)
        notes = await async_session.execute(query)

        if notes is None:
            logging.warning(f"Notes with user id {user_id} do not exist")
            return

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

    async def remove_note_by_id(self, async_session: AsyncSession, note_id: int, user_id: int) -> int | None:
        query = delete(Note).where(and_(Note.id == note_id, Note.user_id == user_id)).returning(Note.id)
        removed_note_id = await async_session.execute(query)
        if removed_note_id is None:
            logging.warning(f"Note with user id {user_id} or note id {note_id} do not exist. Impossible to remove")
            return

        await async_session.commit()
        removed_note_id = removed_note_id.scalar()

        return removed_note_id

    async def remove_all_notes(self, async_session: AsyncSession, user_id: int) -> list[int] | None:
        query = delete(Note).where(Note.user_id == user_id).returning(Note.id)
        removed_note_ids = await async_session.execute(query)
        if removed_note_ids is None:
            logging.warning(f"Notes with user id {user_id} do not exist. Impossible to remove")
            return

        await async_session.commit()

        return removed_note_ids.scalars().all()
