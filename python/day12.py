# Like traditional oop define a class
class Person:
    # This constructor method run wen a new instance of the class is created.
    # here 'self' is equivalent of the Go 'receiver' (p or *p)
    def __init__(self, name, age):
        self.name = name
        self.age = age

    # Getter method just to read data
    def greet(self):
        print(f"Hi, my name is {self.name} and I am {self.age}")

    def set_age(self, new_age):
        self.age = new_age
# -------------Main code execution ----------------
person1 = Person("Alice", 30)
person1.greet()
person1.set_age(31)
person1.greet()