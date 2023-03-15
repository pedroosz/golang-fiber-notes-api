## Notes API
#### Technologies
- Golang
- Gorm
- Gofiber

#### Getting started
##### Dependencies
You can clone this repo and install all dependencies by using `go get github.com/pedroosz/golang-fiber-notes-api/`


#### Endpoints
- `GET` `/api/notes` - Get all notes
- `GET` `/api/notes/:id` - Get a specific note
- `POST` `/api/notes` - Create a note
- `DELETE` `/api/notes/:id` - Delete a note

#### Response examples

### GET /api/notes
##### Expected response
> **(Status code 200)**
```json
[
    {
        "ID": 1,
        "CreatedAt": "2023-03-15T15:16:42.006272132Z",
        "UpdatedAt": "2023-03-15T15:16:42.006272132Z",
        "DeletedAt": null,
        "title": "test",
        "content": "test"
    }
]
```

### GET /api/notes/:id
##### Expected response
> **(Status code 200)**
```json
{
    "ID": 1,
    "CreatedAt": "2023-03-15T15:16:42.006272132Z",
    "UpdatedAt": "2023-03-15T15:16:42.006272132Z",
    "DeletedAt": null,
    "title": "test",
    "content": "test"
}
```
##### Possible responses
> **(Status code 404)**
```json
{
    "message": "Not found",
}
```

### POST - /api/notes
##### Expected body
```json
{
    "title": "test",
    "content": "test"
}
```
##### Expected response
> **(Status code 200)**
```json
{
    "ID": 1,
    "CreatedAt": "2023-03-15T15:16:42.006272132Z",
    "UpdatedAt": "2023-03-15T15:16:42.006272132Z",
    "DeletedAt": null,
    "title": "test",
    "content": "test"
}
```
##### Possible responses
> **(Status code 503)**

### DELETE - /api/notes/:id
##### Expected response
> **(Status code 200)**
```json
{
    "message": "Note deleted."
}
```
##### Possible responses
> **(Status code 404)**
```json
{
    "message": "Not found",
}
```
