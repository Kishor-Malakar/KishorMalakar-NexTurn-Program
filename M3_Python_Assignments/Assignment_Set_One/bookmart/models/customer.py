class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def view_customer_details(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"