class InvalidInputException(Exception):
    def __init__(self):
        super().__init__(f"Price(float) and Quantity(integer) must be numeric and positive.\n")

class BookNotFoundException(Exception):
    def __init__(self):
        super().__init__(f"The book requested is not found\n")

class AllFieldsRequiredException(Exception):
    def __init__(self):
        super().__init__(f"All fields are required.\n")

class OutOfStockException(Exception):
    def __init__(self,quantity):
        self.quantity = quantity
        super().__init__(f"Error: Only {self.quantity} copies available. Sale cannot be completed.\n")

class DuplicateBookException(Exception):
    def __init__(self):
        super().__init__("A book with the same title already exists.")

class DuplicateCustomerException(Exception):
    def __init__(self):
        super().__init__("A customer with the same email already exists.")


def validateInput(price, quantity):
    try:
        price = float(price)
        quantity = int(quantity)
        if price <= 0 or quantity <= 0:
            raise ValueError()
        return price, quantity
    except Exception as e:
        return None, None