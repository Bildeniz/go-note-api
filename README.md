Simple Note API project with GoLang

runing on 0.0.0.0:8080
# API:
### * GET "/api/notes"
returning a list with 200
````json
[
    {
        "id": 1,
        "created_at": "2025-01-17T17:10:11.6608327+03:00",
        "title": "New Note",
        "content": "Hello World"
    }
]
````
### * POST "/api/notes"
POST parameter 
```JSON
{
"Title": "<minimum length: 1, maximum length: 255, required>",
"Content": "<minimum length: 1, required>"
}
```

returning a object with  201
````JSON
{
    "ID": 5,
    "CreatedAt": "2025-01-28T17:54:42.2985952+03:00",
    "UpdatedAt": "2025-01-28T17:54:42.2985952+03:00",
    "DeletedAt": null,
    "Title": "Hello World",
    "Content": "This is a test project"
}
````
