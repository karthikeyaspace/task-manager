GO Task Manager

route
    - GET /tasks - get all tasks
    - POST /tasks - create a new task
    - GET /task/:id - get a task by id
    - PUT /task/:id - update a task by id
    - DELETE /task/:id - delete a task by id

model
    - Task
        - id: number
        - name: string
        - description: string
        - completed: boolean
    
service 
    - getTasks
    - createTask
    - getTask
    - updateTask
    - deleteTask


<!-- https://chatgpt.com/c/67245f29-42d8-800e-ac10-afafbf1bbfaf -->
