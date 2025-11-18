import threading
import queue
import time


def send_message(q):
    print("Thread: Thinking about what to send")
    time.sleep(2)

    q.put("Hello from the other side")
    print("Thread: message sent")

# main execution
message_queue = queue.Queue()

# start thread 
t = threading.Thread(target=send_message, args=(message_queue,))
t.start()

print("Main: Waiting for data...")
msg = message_queue.get()
print(f"Main: received message: {msg}")