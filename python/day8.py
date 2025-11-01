def say_hello():
    print("Hello from a function")

def add(x,y):
    sum = x+y
    return sum

# Python's way of handling errors is "exceptions".
# The "try...except" block is the standard way to handle
# code that might fail.
def divide(numerator, denominator):
    if denominator == 0:
        raise ValueError("Cannot devide by zero")
    
    return numerator / denominator

say_hello()

print(f"2 + 3 is :",add(2,3))

# We must "try" to run the code that might fail.
try:
    quotient = divide(10,2)
    print(f"10 / 2 = {quotient}")

    quotient2 = divide(9,0)
    print(f"9 / 0 = {quotient2}")
except ValueError as e:
    print(f"An error occurred: {e}")

# Python *can* return multiple values (as a tuple)
def get_name():
    return "John", "Deo"

first, last = get_name()
print(f"Name is :{first} {last}")