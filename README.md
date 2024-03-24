# Simple API for Books using Golang

This is a simple API for managing books, implemented using Golang. The API provides endpoints for performing CRUD (Create, Read, Update, Delete) operations on books.

## Endpoints

- `/books`: 
  - **GET**: Retrieve all books.
  - **POST**: Add a new book.
  - **PUT**: Update an existing book.
  - **DELETE**: Delete a book.

## Usage

1. Clone the repository:

git clone https://github.com/lokeshb003/Simple-API-using-Golang

2. Navigate to the project directory:

   cd Simple-API-using-Golang

3. Build and run the application:

   go build
   ./Simple-API-using-Golang

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
        "name": "Author 1"
    },
    {
        "id": 2,
        "title": "Example Book 2",
        "name": "Author 2"
    }
]
## Dependencies

- Gin: HTTP web framework for Golang.
