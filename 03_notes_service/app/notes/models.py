import uuid
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy import Column, String, Integer, Text, TIMESTAMP, ForeignKey
from app.settings.base import Base
from datetime import datetime


class Note(Base):

    __tablename__ = "notes"

    note_uuid = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    category_id = Column(Integer, ForeignKey("categories.category_id"), nullable=True)
    user_uuid = Column(UUID(as_uuid=True), default=uuid.uuid4)
    title = Column(String, nullable=False)
    body = Column(Text)
    update_at = Column(TIMESTAMP, nullable=True)
    created_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)



class Category(Base):

    __tablename__ = "categories"

    category_id =Column(Integer, primary_key=True, index=True, unique=True)
    category_name = Column(String, nullable=False)
    update_at = Column(TIMESTAMP, nullable=True)
    created_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)