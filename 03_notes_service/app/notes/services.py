from datetime import datetime
from .models import Note, Category
from . import schema
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import update, and_, func, select


async def verify_permission_note(note_uuid: str, user_uuid: str, async_session: AsyncSession):

    """
    check user`s permission to get note
    """

    async with async_session() as session:

        query = select(Note.user_uuid).where(Note.note_uuid == note_uuid)
        query_result = await session.execute(query)
        user_uuid_from_db = query_result.scalar()

    if user_uuid == str(user_uuid_from_db):
        return True


async def verify_permission_category(category_id: int, user_uuid: str, async_session: AsyncSession):

    """
    check user`s permission to get category and notes related to this category
    """
    
    async with async_session() as session:

        query = select(Category.user_uuid).where(Category.category_id == category_id)
        query_result = await session.execute(query)
        user_uuid_from_db = query_result.scalar()

    print(user_uuid)
    print(user_uuid_from_db)
    if user_uuid == str(user_uuid_from_db):
        return True


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


async def create_category(category: schema.CategoryCreate, async_session: AsyncSession):

    """
    insert a new category into db
    """

    async with async_session() as session:
        
        new_category = Category(user_uuid=category.user_uuid, category_name=category.category_name, created_at=datetime.now())

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


async def get_list_notes(user_uuid: str, async_session: AsyncSession):

    """
    extract all notes from db
    """

    notes = []


    async with async_session() as session:
        #left outer join
        query = select(Note, Category.category_name).join(Category, Note.category_id == Category.category_id, isouter=True).where(Note.user_uuid == user_uuid)
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
    

async def update_note(note_uuid: str, params_to_update: schema.NoteUpdate, async_session: AsyncSession):

    
    data_to_update = params_to_update.dict(exclude_none=True)

    print(data_to_update)
    
    
    async with async_session() as session:

        if data_to_update["category_name"] is not None:
            query = select(Category.category_id).where(Category.category_name == data_to_update["category_name"])
            category_id = await session.execute(query)
            data_to_update["category_id"] = category_id.scalar()

        data_to_update.pop("category_name")

        query_to_update = update(Note).where(Note.note_uuid == note_uuid).values(**data_to_update).returning(Note.note_uuid)
        note_uuid = await session.execute(query_to_update)
        await session.commit()

    note_uuid_updated = note_uuid.scalar()

    return str(note_uuid_updated) #str() is necessary to serialize by protobuf