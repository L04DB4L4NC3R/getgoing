### Todo-API 
todo api specification

<br>

| Task | URL | Method | Response code | Response |
|:----:|:---:|:------:|:-------------:|:--------:|
| Check API health | /ping | GET | 200 | pong |
| Create an entry | / | POST | 201 | json of the created object | 
| Read all entries | / | GET | 200 | json of all entries |
| Read entries by id | /?id=value | GET | 200 | json of entries using keys and values | 
| Delete entries | /:id | DELETE | 200 | status |
