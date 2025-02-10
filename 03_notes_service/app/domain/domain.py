from dataclasses import dataclass, field
from typing import Optional


@dataclass
class NoteEntity:
    category_id: int
    user_id: int
    title: str
    body: str
    update_at: int
    created_at: int
    id: Optional[int] = field(default=None)

    def to_dict(self):
        return self.__dict__

    @classmethod
    def to_model(cls, dict_obj):
        return cls(**dict_obj)


@dataclass
class CategoryEntity:
    category_name: str
    user_id: int
    update_at: int
    created_at: int
    id: Optional[int] = field(default=None)

    def to_dict(self):
        return self.__dict__

    @classmethod
    def to_model(cls, dict_obj):
        return cls(**dict_obj)