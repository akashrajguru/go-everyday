dynamic_numbers = []
print(f"Enpty list: {dynamic_numbers}")


#  add to list
dynamic_numbers.append(10)
dynamic_numbers.append(20)
dynamic_numbers.append(30)
print(f"List : {dynamic_numbers}")
print(f"Item at index {dynamic_numbers[1]}")
print(f"Length of list: {len(dynamic_numbers)}")

mix_list = [1, "Hello", True]
print(f"Mixed list {mix_list}")

vowels = ["a", "e", "i", "o", "u"]
print("\nLooping:")
for val in vowels:
    print(f"value: {val}")

# If you need the index, you use the 'enumerate()' function.
for index, val in enumerate(vowels):
    print(f"Index: {index} Value: {val}")