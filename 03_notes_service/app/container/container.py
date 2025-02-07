import punq

from repository.notes_repository import INoteRepository, NoteRepository
from services.notes_services import INoteService, NoteService


class DIContainer:
    container = punq.Container()

    container.register(INoteRepository, NoteRepository)
    container.register(INoteService, NoteService)

    def get_notes_service(self):
        return self.container.register(INoteService)


container = DIContainer()
