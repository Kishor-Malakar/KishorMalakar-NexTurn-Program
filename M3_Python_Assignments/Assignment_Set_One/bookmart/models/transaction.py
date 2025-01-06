from .customer import Customer

class Transaction(Customer):
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def view_transaction_details(self):
        return f"Customer: {self.name}, Book: {self.book_title}, Quantity Sold: {self.quantity_sold}"
