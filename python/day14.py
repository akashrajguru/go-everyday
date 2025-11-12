import threading
import time


# Function to run thread 
def say(message, count):
    for i in range(count):
        print(message)



#  main execution
say("hello", 3)

# Create thread object
t1 = threading.Thread(target=say, args=("World", 5))

t1.start()

# Create and start an anonymous (lambda) thread.
t2 = threading.Thread(target=lambda msg: print(msg), args=("I am a new thread...",))
t2.start()

print("Main is waiting....")
t1.join()
t2.join()

print("Main is done.")