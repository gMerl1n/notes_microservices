import punq

from repository.notes_repository import INoteRepository, NoteRepository
from repository.category_repository import ICategoryRepository, CategoryRepository
from services.notes_services import INoteService, NoteService
from services.category_services import ICategoryService, CategoryService


class DIContainer:
    container = punq.Container()

    container.register(INoteRepository, NoteRepository)
    container.register(INoteService, NoteService)

    container.register(ICategoryRepository, CategoryRepository)
    container.register(ICategoryService, CategoryService)

    def get_notes_service(self) -> INoteService:
        return self.container.resolve(INoteService)

    def get_category_service(self) -> ICategoryService:
        return self.container.resolve(ICategoryService)


container = DIContainer()
