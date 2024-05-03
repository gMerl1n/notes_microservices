from .models import Note, Category
from . import schema
from sqlalchemy.ext.asyncio import AsyncSession


async def create_note(note: schema.Note, async_session: AsyncSession):

    new_note = Note(**note.dict())
    async_session.add(new_note)
    await async_session.commit()

    return new_note.note_uuid
