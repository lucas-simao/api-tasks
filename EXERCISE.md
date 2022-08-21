# API developer practical exercise

### Requirements

#### You are developing a software to account for maintenance tasks performed during a working day. This application has two types of users (Manager, Technician).
#### The technician performs tasks and is only able to see, create or update his own performed tasks.
#### The manager can see tasks from all the technicians, delete them, and should be notified when some tech performs a task.
#### A task has a summary (max: 2500 characters) and a date when it was performed, the summary from the task can contain personal information.

### Development

Features:
- Create API endpoint to save a new task
- Create API endpoint to list tasks
- Notify manager of each task performed by the tech (This notification can be just a print saying “The tech X
performed the task Y on date Z”)
- This notification should not block any http request
Tech Requirements:
- Use either Go or Node to develop this HTTP API
- Create a local development environment using docker containing this service and a MySQL database
- Use MySQL database to persist data from the application
- Features should have unit tests to ensure they are working properly