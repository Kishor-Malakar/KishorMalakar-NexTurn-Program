from ..models.customer import Customer
from ..errors.error import AllFieldsRequiredException,DuplicateCustomerException

customers = []

def add_customer(name, email, phone):
    if not name or not email or not phone:
        raise AllFieldsRequiredException()
    
    for customer in customers:
        if customer.email.lower() == email.lower():
            raise DuplicateCustomerException()
        
    customers.append(Customer(name, email, phone))
    return "Customer added successfully!"

def view_customers():
    if len(customers)==0:
        print("No Customers Available!\n")
    return [customer.view_customer_details() for customer in customers]
