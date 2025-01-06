from flask import Blueprint, request, jsonify
from . import db
from app.models import Book

def error_response(status_code, error, message):
    return jsonify({"error": error, "message": message}), status_code

book_routes = Blueprint('book_routes', __name__)

genre_list = ['Fiction', 'Non-Fiction', 'Science Fiction', 'Mystery']

@book_routes.route('/books', methods=['POST'])
def create_book():
    data = request.json
    if 'title' not in data or 'author' not in data or 'published_year' not in data or 'genre' not in data:
        return error_response(400, "Bad Request", "Missing required fields")
    if data['genre'] not in genre_list:
        return error_response(400, "Bad Request", "Invalid genre")
    if data['published_year'] < 0:
        return error_response(400, "Bad Request", "Invalid published year")
    new_book = Book(title=data['title'], author=data['author'], published_year=data['published_year'], genre=data['genre'])
    db.session.add(new_book)
    db.session.commit()
    return jsonify({"message": "Book created", "book": {"id": new_book.id, "title": new_book.title, "author": new_book.author, "published_year": new_book.published_year, "genre": new_book.genre}}), 201

@book_routes.route('/books', methods=['GET'])
def get_books():
    genre = request.args.get('genre')
    author = request.args.get('author')
    books = Book.query.all()
    books_list = [{"id": book.id, "title": book.title, "author": book.author, "published_year": book.published_year, "genre": book.genre} for book in books]
    
    filtered_books = books_list

    if genre:
        filtered_books = [book for book in filtered_books if book['genre'].lower() == genre.lower()]

    if author:
        filtered_books = [book for book in filtered_books if book['author'].lower() == author.lower()]

    if len(filtered_books)==0:
        return error_response(400, "Bad Request", "No books available")
    
    return jsonify(filtered_books)

@book_routes.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    book = Book.query.get_or_404(book_id)
    if book is None:
        return error_response(404, "Not Found", "Book not found")
    return jsonify({"id": book.id, "title": book.title, "author": book.author, "published_year": book.published_year, "genre": book.genre})
@book_routes.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    book = Book.query.get_or_404(book_id)
    if book is None:
        return error_response(404, "Not Found", "Book not found")

    data = request.json

    if 'title' in data:
        book.title = data.get('title')

    if 'author' in data:
        book.author = data.get('author')

    if 'published_year' in data:
        if data['published_year'] < 0:  
            return error_response(400, "Bad Request", "Invalid published year")
        book.published_year = data.get('published_year')

    if 'genre' in data:
        if data['genre'] not in genre_list:  
            return error_response(400, "Bad Request", "Invalid genre")
        book.genre = data.get('genre')

    db.session.commit()

    return jsonify({"message": "Book updated","book": {"id": book.id,"title": book.title,"author": book.author,"published_year": book.published_year,"genre": book.genre,}})


@book_routes.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    book = Book.query.get_or_404(book_id)
    if book is None:
        return error_response(404, "Not Found", "Book not found")
    db.session.delete(book)
    db.session.commit()
    return jsonify({"message": "Book deleted"})
