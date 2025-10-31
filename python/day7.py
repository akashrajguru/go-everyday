# Python has two distinct loop keywords that is 'for' and 'while'

# Iteration loop
print("Iteration loop")
for i in range(5):
    print(f"i = {i}")

# While loop
print("\nWhile loop")
j = 0
while j < 3:
    print(f"j = {j}")
    j += 1 # j++ is not valid in python

# Infinite loop with break 
print("\nInfinite loop with break")
k = 0
while True: # Python infinite loop
    print(f"k = {k}")
    if k == 2:
        print("Braking loop")
        break
    k += 1