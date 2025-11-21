import queue
import threading
import time

def producer(q):
    for i in range(1,6):
        q.put(i)
        time.sleep(0.2)

    # Simulate closing by sending a none value
    q.put(None)
    print("Producer finished")

    #  main execution
    # create a buffered queue
    q = queue.Queue(maxsize=2)
    q.put("Message 1")
    q.put("Message 2")

    print("Sent 2 message without a receiver yet")
    print(f"Received :{q.get()}")

# looping

num_q = queue.Queue()
t = threading.Thread(target=producer, args=(num_q,))
t.start()

print("waiting for numbers")

while True:
    val = num_q.get()
    if val is None:
        break
    print(f"Received number: {val}")

print("Done")