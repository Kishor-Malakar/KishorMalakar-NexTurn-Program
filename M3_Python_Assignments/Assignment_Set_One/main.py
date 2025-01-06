from bookmart import *

while True:
    print("\nWelcome to BookMart!")
    print("1. Book Management")
    print("2. Customer Management")
    print("3. Sales Management")
    print("4. Exit")
    choice = input("Enter your choice (1-4): ")

    if choice == "1":
        print("\n1.1. Add Book\n1.2. View Books\n1.3. Search Book")
        sub_choice = input("Enter your choice (1-3): ")
        if sub_choice == "1":
            title = input("Title: ")
            author = input("Author: ")
            price = input("Price: ")
            quantity = input("Quantity: ")
            try:
                print(add_book(title, author, price, quantity))
            except (InvalidInputException,DuplicateBookException) as e:
                print(f"Error: {e}")
        elif sub_choice == "2":
            print("\n".join(view_books()))
        elif sub_choice == "3":
            query = input("Enter title or author to search: ")
            try:
                print("\n".join(search_book(query)))
            except BookNotFoundException as e:
                print(f"Error: {e}")
    elif choice == "2":
        print("\n2.1. Add Customer\n2.2. View Customers")
        sub_choice = input("Enter your choice (1-2): ")
        if sub_choice == "1":
            name = input("Name: ")
            email = input("Email: ")
            phone = input("Phone: ")
            try:
                print(add_customer(name, email, phone))
            except (AllFieldsRequiredException,DuplicateCustomerException) as e:
                print(f"Error: {e}")
        elif sub_choice == "2":
            print("\n".join(view_customers()))
    elif choice == "3":
        print("\n3.1. Sell Book\n3.2. View Sales")
        sub_choice = input("Enter your choice (1-2): ")
        if sub_choice == "1":
            name = input("Customer Name: ")
            email = input("Customer Email: ")
            phone = input("Customer Phone: ")
            book_title = input("Book Title: ")
            quantity = input("Quantity: ")
            try:
                print(sell_book(name, email, phone, book_title, quantity))
            except (BookNotFoundException,OutOfStockException,InvalidInputException) as e:
                print(f"Error: {e}")
        elif sub_choice == "2":
            print("\n".join(view_sales()))
    elif choice == "4":
        print("Thank you for using BookMart!")
        break
    else:
        print("Invalid choice. Please try again.")
