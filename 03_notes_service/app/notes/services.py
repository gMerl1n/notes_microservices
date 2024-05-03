from datetime import datetime
from .models import Note, Category
from . import schema
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import update, and_, select


async def create_note(note: schema.Note, async_session: AsyncSession):

    note_to_add = note.dict(exclude_none=True)
    
    # if note_to_add["category_id"] is not None:
    #     category = await get_category_by_id(note.category_id, async_session=async_session)
    #     note_to_add["category_id"] = category

    # print(note_to_add)

    async with async_session() as session:
        new_note = Note(**note_to_add, created_at=datetime.now())
        session.add(new_note)
        await session.commit()
        await session.refresh(new_note)


    return new_note.note_uuid


async def create_category(category: schema.Category, async_session: AsyncSession):

    async with async_session() as session:
        new_category = Category(category_name=category.category_name, created_at=datetime.now())
        session.add(new_category)
        await session.commit()
        await session.refresh(new_category)

    return new_category.category_id


async def get_category_by_id(category_id: int, async_session: AsyncSession):

    async with async_session() as session:
        query = select(Category).where(Category.category_id==category_id)
        res = await session.execute(query)
        category = res.fetchone()
        if category is not None:
             return category[0]

