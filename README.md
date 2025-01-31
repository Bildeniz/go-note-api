# Notes Application API Documentation

This API provides a simple note management system. Users can create, list, update, and delete notes. The API is built using the `Gin` web framework and the `GORM` ORM library.

---

## Model: `Notes`

The model representing notes is as follows:

```go
type Notes struct {
    gorm.Model
    Title   string `binding:"required,min=1,max=255"` // Note title (required, 1-255 characters)
    Content string `binding:"required,min=1"`        // Note content (required, at least 1 character)
}
```

### Fields:
- **ID**: Automatically generated primary key.
- **CreatedAt**: The date the note was created.
- **UpdatedAt**: The date the note was last updated.
- **DeletedAt**: The date the note was deleted (used for soft delete).
- **Title**: Note title (required, 1-255 characters).
- **Content**: Note content (required, at least 1 character).

---

## Endpoints

### 1. **List All Notes**
Retrieves all notes.

- **URL**: `/notes`
- **Method**: `GET`
- **Response**:
    - `200 OK`: Notes retrieved successfully.
      ```json
      [
          {
              "ID": 1,
              "CreatedAt": "2023-10-01T12:00:00Z",
              "UpdatedAt": "2023-10-01T12:00:00Z",
              "DeletedAt": null,
              "Title": "First Note",
              "Content": "This is my first note."
          },
          {
              "ID": 2,
              "CreatedAt": "2023-10-02T12:00:00Z",
              "UpdatedAt": "2023-10-02T12:00:00Z",
              "DeletedAt": null,
              "Title": "Second Note",
              "Content": "This is my second note."
          }
      ]
      ```
    - `500 Internal Server Error`: An error occurred while retrieving the notes.
      ```json
      {
          "error": "No Notes found"
      }
      ```

---

### 2. **Create a New Note**
Creates a new note.

- **URL**: `/notes`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
      "Title": "New Note",
      "Content": "This is a new note."
  }
  ```
- **Response**:
    - `201 Created`: Note created successfully.
      ```json
      {
          "ID": 3,
          "CreatedAt": "2023-10-03T12:00:00Z",
          "UpdatedAt": "2023-10-03T12:00:00Z",
          "DeletedAt": null,
          "Title": "New Note",
          "Content": "This is a new note."
      }
      ```
    - `400 Bad Request`: Invalid request data.
      ```json
      {
          "error": "Title is required"
      }
      ```
    - `500 Internal Server Error`: Failed to create the note.
      ```json
      {
          "error": "Failed to create note"
      }
      ```

---

### 3. **Retrieve a Specific Note**
Retrieves a specific note by its ID.

- **URL**: `/notes/:id`
- **Method**: `GET`
- **Parameters**:
    - `id`: The ID of the note to retrieve.
- **Response**:
    - `200 OK`: Note retrieved successfully.
      ```json
      {
          "ID": 1,
          "CreatedAt": "2023-10-01T12:00:00Z",
          "UpdatedAt": "2023-10-01T12:00:00Z",
          "DeletedAt": null,
          "Title": "First Note",
          "Content": "This is my first note."
      }
      ```
    - `404 Not Found`: Note not found.
      ```json
      {
          "error": "Note not found"
      }
      ```

---

### 4. **Update a Note**
Updates an existing note by its ID.

- **URL**: `/notes/:id`
- **Method**: `PUT`
- **Parameters**:
    - `id`: The ID of the note to update.
- **Request Body**:
  ```json
  {
      "Title": "Updated Note",
      "Content": "This note has been updated."
  }
  ```
- **Response**:
    - `200 OK`: Note updated successfully.
      ```json
      {
          "ID": 1,
          "CreatedAt": "2023-10-01T12:00:00Z",
          "UpdatedAt": "2023-10-03T12:00:00Z",
          "DeletedAt": null,
          "Title": "Updated Note",
          "Content": "This note has been updated."
      }
      ```
    - `400 Bad Request`: Invalid request data.
      ```json
      {
          "error": "Bad request data"
      }
      ```
    - `404 Not Found`: Note not found.
      ```json
      {
          "error": "Note not found"
      }
      ```

---

### 5. **Delete a Note**
Deletes a note by its ID.

- **URL**: `/notes/:id`
- **Method**: `DELETE`
- **Parameters**:
    - `id`: The ID of the note to delete.
- **Response**:
    - `200 OK`: Note deleted successfully.
      ```json
      {
          "message": "Note deleted"
      }
      ```
    - `404 Not Found`: Note not found.
      ```json
      {
          "error": "Note not found"
      }
      ```
    - `500 Internal Server Error`: Failed to delete the note.
      ```json
      {
          "error": "Failed to delete note"
      }
      ```

---

## Example Usage

### 1. List All Notes
```bash
curl -X GET http://localhost:8080/notes
```

### 2. Create a New Note
```bash
curl -X POST -H "Content-Type: application/json" -d '{"Title": "New Note", "Content": "This is a new note."}' http://localhost:8080/notes
```

### 3. Retrieve a Specific Note
```bash
curl -X GET http://localhost:8080/notes/1
```

### 4. Update a Note
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"Title": "Updated Note", "Content": "This note has been updated."}' http://localhost:8080/notes/1
```

### 5. Delete a Note
```bash
curl -X DELETE http://localhost:8080/notes/1
```

---

This documentation provides a comprehensive guide to using the Notes API. You can use tools like `curl` or Postman to interact with the API and test its functionality.