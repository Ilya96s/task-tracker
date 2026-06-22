1) создавать списки
    1) создать пользователя
    2) создать список
2) создавать задачи
    1) создать пользователя
    2) создать список
    3) создать задачу
3) помечать задачи как выполненные
4) удалять задачи
5) редактировать задачу
6) удалить список
7) редактировать список
8) получить список задач
9) поделиться списком (на чтение/запись)
    1) создать пользователя А
    2) создать пользователя Б
    3) пользователь А создает список
    4) пользователь А делится списком с пользователем Б
10) создать список для шаблонов задач
11) приоритет задач
12) переместить задачу из одного списка в другой
13) регистрация пользователя
14) логирование пользователя


HTTP API

resource: list, task, user, (template пока не горит)

- list.create
- POST /v1/lists {"name": "test"} -> {"list" {id":"aa-bb", "name": "test"}}

- list.get
- GET /v1/lists/{list_id} {"id": "aa-bb"} -> {"name": "test"}

- list.update
- PATCH /v1/list/{list_id} {"name": "test2"} -> {"id":"aa-bb", "name": "test2"}

- list.delete
- POST /v1/lists/{list_id}



- task.create
- #2 POST /v1/lists/{list_id}/tasks {"name": "do"} -> {"task" {id":"bb-cc", "name": "do", "done": false}}

- task.get
- GET /v1/lists/{list_id}/tasks/{task_id} {"id": "aa-bb"} -> {"name": "test", "done": false}

- task.update
- PATCH /v1/lists/{list_id}/tasks/{task_id} {"name": "test2"} -> {"id":"aa-bb", "name": "test2", "done": false}

- task.delete
- POST /v1/lists/{list_id}/tasks/{task_id}


DONE TASK
- task.update
- PATCH /v1/lists/{list_id}/tasks/{task_id} {"done": true} -> {"id":"aa-bb", "name": "test2", "done": true}
- task.done
- POST /v1/lists/{list_id}/tasks/{task_id}/done -> {"id":"aa-bb", "name": "test2", "done": true}

CHANGE PRIORITY AND MARK DONE
- task.update_and_done
- POST /v1/lists/{list_id}/tasks/{task_id}/update_and_done -> {"id":"aa-bb", "name": "test2", "done": true}

GET TASKS
- GET /v1/lists/{list_id}?expand=tasks -> {"list" : {"id": "aa-bb", "name" : "test", "tasks": [...]}}
- GET /v1/lists/{list_id} -> {"lists": {"id": "aa-bb", "name" : "test"}, "tasks": [...]}
- GET /v1/lists/{list_id}/tasks -> [...]

GET LISTS
- GET /v1/lists -> {"lists" :[...]}


1) НАСТРОИТЬ ЧТЕНИЕ КОНФИГУРАЦИИ (ПОКА НЕ ВАЖНО)- done
2) НУЖНО ПОДКЛЮЧИТЬ БАЗУ - done
3) РЕАЛИЗОВАТЬ АПИ - done
  4) ПОДКЛЮЧИТЬ МИГРАЦИИ (ПОКА НЕ ВАЖНО)
5) НАСТРОИТЬ ДОКР КОМПОУЗ (БД И СЕРВЕР) (ПОКА НЕ ВАЖНО)


