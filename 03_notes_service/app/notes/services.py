from datetime import datetime
from .models import Note, Category
from . import schema
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import update, and_, func, select


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
        #left outer join
        query = select(Note, Category.category_name).join(Category, Note.category_id == Category.category_id, isouter=True)
        query_result = await session.execute(query)

        if query_result is not None:

            for item in query_result:
                notes.append({
                    "note_uuid": str(item[0].note_uuid), #str() is necessary to serialize by protobuf
                    "category_name": item[1],
                    "title": item[0].title,
                    "body": item[0].body,
                    "update_at": datetime.timestamp(item[0].update_at) if item[0].update_at is not None else None, 
                    "created_at": datetime.timestamp(item[0].created_at),
                })

            return notes
        

async def get_note_by_id(uuid, async_session: AsyncSession):

    """
    extract a note by its uuid from db
    """

    async with async_session() as session:
        
        query = select(Note).where(Note.note_uuid == uuid)
        res = await session.execute(query)
        note = res.scalar()
        if note is not None:
            return note


async def get_notes_by_category(category_id: schema.NotesByCategory, async_session: AsyncSession):

    """
    extract a list of notes by category from db
    """

    notes = []

    async with async_session() as session:
        # inner join
        query_notes = select(Note, Category.category_name).select_from(Note).join(Category, Category.category_id == Note.category_id).where(Note.category_id == category_id.category_id)
        # count notes related to the selected category
        query_count_notes = select([func.count()]).select_from(Note).join(Category, Category.category_id == Note.category_id).where(Note.category_id == category_id.category_id)
        
        query_notes_result = await session.execute(query_notes)
        query_count_result = await session.execute(query_count_notes)

        count_notes = query_count_result.scalar() 

        for item in query_notes_result:
            notes.append({
                "note_uuid": str(item[0].note_uuid), #str() is necessary to serialize by protobuf
                "category": item[1],
                "title": item[0].title,
                "body": item[0].body,
                "update_at": datetime.timestamp(item[0].update_at) if item[0].update_at is not None else None, 
                "created_at": datetime.timestamp(item[0].created_at),
            })

        return notes, count_notes