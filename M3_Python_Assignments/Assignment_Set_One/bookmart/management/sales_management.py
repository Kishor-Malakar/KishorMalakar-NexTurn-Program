from ..models.transaction import Transaction
from ..errors.error import *
from .book_management import books
from .customer_management import add_customer, customers

sales = []

def sell_book(customer_name, email, phone, book_title, quantity):
    book = next((b for b in books if b.title.lower() == book_title.lower()), None)
    temp, quantity = validateInput(1,quantity)
    if quantity is None:
        raise InvalidInputException()
    if not book:
        raise BookNotFoundException()
    if book.quantity < quantity:
        raise OutOfStockException(book.quantity)
    
    customer = next((c for c in customers if c.email == email), None)
    if not customer:
        add_customer(customer_name, email, phone)
    
    book.quantity -= quantity
    sales.append(Transaction(customer_name, email, phone, book_title, quantity))
    return f"Sale successful! Remaining quantity: {book.quantity}"

def view_sales():
    if len(sales)==0:
        print("No Sales Available!\n")
    return [sale.view_transaction_details() for sale in sales]
