4. You can now access the API at `http://localhost:8080/books`.

## Endpoints Details

### GET `/books`

Retrieve all books stored in the database.

**Response:**

```json
[
    {
        "id": 1,
        "title": "Example Book 1",
        "author": "Author 1",
        "year": 2020
    },
    {
        "id": 2,
        "title": "Example Book 2",
        "author": "Author 2",
        "year": 2021
    }
]
## Dependencies

- Gin: HTTP web framework for Golang.
