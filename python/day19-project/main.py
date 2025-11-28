from calculator import simple_math

def main():
    result = simple_math.add(10, 5)
    print(f"10 + 5 = {result}")

    private_func_result = simple_math._subtract(10, 5)
    print(f"10 - 5 = {private_func_result} (Warning: Accessed internal function)")

if __name__ == "__main__":
    main()
