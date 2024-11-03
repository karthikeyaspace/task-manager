### GO Task Manager

#### routes
    - GET /tasks - get all tasks
    - POST /create-task - create a new task
    - PUT /update-task - update a task by id
    - DELETE /delete-task - delete a task by id

#### model
    - Task
        - taskid: string
        - name: string
        - description: string
        - completed: boolean
        - priority: int
    
#### service 
    - getTasks
    - createTask
    - updateTask
    - deleteTask

#### various server request errors
    - http.StatusBadRequest
    - http.StatusNotFound
    - http.StatusInternalServerError


start postres in docker
```code 
docker run --name task -e POSTGRES_USER=task -e POSTGRES_PASSWORD=task -e POSTGRES_DB=task -p 5432:5432 -d postgres 
```
