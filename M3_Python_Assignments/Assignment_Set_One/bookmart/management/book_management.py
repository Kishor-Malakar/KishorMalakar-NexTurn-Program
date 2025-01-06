from ..models.book import Book
from ..errors.error import *

books = []

def add_book(title, author, price, quantity):
    price, quantity = validateInput(price, quantity)
    if price is None and quantity is None:
        raise InvalidInputException()
    
    for book in books:
        if book.title.lower() == title.lower():
            raise DuplicateBookException()
        
    books.append(Book(title, author, price, quantity))
    return "Book added successfully!"

def view_books():
    if len(books)==0:
        print("No Books Available!\n")
    return [book.view_book_details() for book in books]

def search_book(query):
    result = [book.view_book_details() for book in books if query.lower() in book.title.lower() or query.lower() in book.author.lower()]
    if result:
        return result
    else:
        raise BookNotFoundException()
