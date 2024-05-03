from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateNoteRequest(_message.Message):
    __slots__ = ("userUUID", "title", "body")
    USERUUID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    BODY_FIELD_NUMBER: _ClassVar[int]
    userUUID: str
    title: str
    body: str
    def __init__(self, userUUID: _Optional[str] = ..., title: _Optional[str] = ..., body: _Optional[str] = ...) -> None: ...

class CreateNoteResponse(_message.Message):
    __slots__ = ("noteUUID",)
    NOTEUUID_FIELD_NUMBER: _ClassVar[int]
    noteUUID: str
    def __init__(self, noteUUID: _Optional[str] = ...) -> None: ...

class GetNoteRequest(_message.Message):
    __slots__ = ("noteUUID",)
    NOTEUUID_FIELD_NUMBER: _ClassVar[int]
    noteUUID: str
    def __init__(self, noteUUID: _Optional[str] = ...) -> None: ...

class GetNoteResponse(_message.Message):
    __slots__ = ("noteUUID", "title", "body", "update_at", "created_at")
    NOTEUUID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    BODY_FIELD_NUMBER: _ClassVar[int]
    UPDATE_AT_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    noteUUID: str
    title: str
    body: str
    update_at: float
    created_at: float
    def __init__(self, noteUUID: _Optional[str] = ..., title: _Optional[str] = ..., body: _Optional[str] = ..., update_at: _Optional[float] = ..., created_at: _Optional[float] = ...) -> None: ...

class GetListNotesRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetListNotesResponse(_message.Message):
    __slots__ = ("notes",)
    NOTES_FIELD_NUMBER: _ClassVar[int]
    notes: _containers.RepeatedCompositeFieldContainer[GetNoteResponse]
    def __init__(self, notes: _Optional[_Iterable[_Union[GetNoteResponse, _Mapping]]] = ...) -> None: ...
