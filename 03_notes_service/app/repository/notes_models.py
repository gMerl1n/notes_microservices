from app.settings.base import Base

from datetime import datetime
from sqlalchemy import Column, String, Integer, Text, TIMESTAMP, ForeignKey


class Note(Base):

    __tablename__ = "notes"

    note_id = Column(Integer, primary_key=True, unique=True)
    category_id = Column(Integer, ForeignKey("categories.category_id"), nullable=True, default=None)
    user_id = Column(Integer, primary_key=True, unique=True)
    title = Column(String, nullable=False)
    body = Column(Text, nullable=True)
    update_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)
    created_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)



class Category(Base):

    __tablename__ = "categories"

    category_id = Column(Integer, primary_key=True, unique=True)
    category_name = Column(String, nullable=False)
    user_id = Column(Integer, primary_key=True, unique=True)
    update_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)
    created_at = Column(TIMESTAMP, nullable=False, default=datetime.timestamp)