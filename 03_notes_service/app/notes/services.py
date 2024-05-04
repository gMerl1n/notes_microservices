from datetime import datetime
from .models import Note, Category
from . import schema
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import update, and_, select


async def create_note(note: schema.Note, async_session: AsyncSession):

    """
    insert a new note into db
    """

    category_id = await get_category_by_name(note.category_name, async_session=async_session)

    async with async_session() as session:
        new_note = Note(category_id=category_id,
                        user_uuid=note.user_uuid,
                        title=note.title,
                        body=note.body,
                        created_at=datetime.now())
        session.add(new_note)
        await session.commit()
        await session.refresh(new_note)


    return new_note.note_uuid


async def create_category(category: schema.Category, async_session: AsyncSession):

    """
    insert a new category into db
    """

    async with async_session() as session:
        new_category = Category(category_name=category.category_name, created_at=datetime.now())
        session.add(new_category)
        await session.commit()
        await session.refresh(new_category)

    return new_category.category_id


async def get_category_by_name(category_name: str, async_session: AsyncSession):

    """
    find category into db by name
    """

    async with async_session() as session:
        query = select(Category).where(Category.category_name == category_name)
        res = await session.execute(query)
        category = res.scalar()
        if category is not None:
             return category.category_id


async def get_list_notes(async_session: AsyncSession):

    """
    extract all notes from db
    """

    notes = []

    async with async_session() as session:
        query = select(Note)
        res = await session.execute(query)
        if res is not None:
            for note in res.scalars():
                notes.append({
                    "note_uuid": str(note.note_uuid), #str() is necessary to serialize by protobuf
                    "title": note.title,
                    "body": note.body,
                    "update_at": datetime.timestamp(note.update_at) if note.update_at is not None else None, 
                    "created_at": datetime.timestamp(note.created_at),
                })
                #notes.append(note.__dict__)

            return notes
        

async def get_note_by_id(uuid, async_session: AsyncSession):
    async with async_session() as session:
        query = select(Note).where(Note.note_uuid == uuid)
        res = await session.execute(query)
        note = res.scalar()
        print(note.note_uuid)
        if note is not None:
            return note



# // {
# //     "category_name": "dating"
# // }

# // {
# //     "user_uuid": "133e4567-e89b-12d3-a456-426614174000",
# //     "title": "title1",
# //     "body": "body1"
# // }