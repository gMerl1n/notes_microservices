from app.settings.base import Base
from app.domain.domain import NoteEntity, CategoryEntity
from datetime import datetime
from sqlalchemy import Column, String, Integer, Text, TIMESTAMP, ForeignKey


class Note(Base):
    __tablename__ = "notes"

    id = Column(Integer, primary_key=True, unique=True)
    category_id = Column(Integer, ForeignKey("categories.id"), nullable=True, default=None)
    user_id = Column(Integer, nullable=False)
    title = Column(String, nullable=False)
    body = Column(Text, nullable=True)
    update_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)
    created_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)

    @classmethod
    def to_note_model(cls, obj: NoteEntity):
        return cls(
            category_id=obj.category_id,
            user_id=obj.user_id,
            title=obj.title,
            body=obj.body,
            update_at=datetime.fromtimestamp(obj.update_at),
            created_at=datetime.fromtimestamp(obj.created_at)
        )


class Category(Base):
    __tablename__ = "categories"

    id = Column(Integer, primary_key=True, unique=True)
    category_name = Column(String, nullable=False)
    user_id = Column(Integer, nullable=False)
    update_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)
    created_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)

    @classmethod
    def to_category_model(cls, obj: CategoryEntity):
        return cls(
            category_name=obj.category_name,
            user_id=obj.user_id,
            update_at=datetime.fromtimestamp(obj.update_at),
            created_at=datetime.fromtimestamp(obj.created_at)
        )
