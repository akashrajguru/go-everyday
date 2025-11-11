import math

# No interface is defined
# Define Rectangle class
class Rectangle:
    def __init__(self, width, height):
        self.width = width
        self.height = height
    
    def area(self):
        return self.width * self.height
    

# Define Circle class
class Circle:
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return math.pi * self.radius * self.radius
    
# This function doesn't care about the type of 'shape'.
# It just assumes 'shape' has an '.area()' method.
# This is "Duck Typing".
def print_shape_area(shape):
    try:
        print(f"The are of this shape is :{shape.area()}")
    except AttributeError:
        print(f"Error: Object does not have an area method")


rect = Rectangle(10,5)
circ = Circle(7)

print_shape_area(rect)
print_shape_area(circ)

my_string = "hello"
print_shape_area(my_string)