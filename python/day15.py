import threading
import time

def worker(id):
    print(f"worker {id} starting....")
    time.sleep(1)
    print(f"Worker {id} done.")

# Main execution
# Lets store thread objects in list
threads = []
new_workers = 3

for i in range(1, new_workers + 1):
    t = threading.Thread(target=worker, args=(i,))
    t.start()
    threads.append(t)

print("Main is waiting for threads to finish....")

# Loop over list of threads and join() each other.
# This blocks the main thread, waiting for 't' to finish
# before moving to the next one.
for t in threads:
    t.join()

print("All threads finish. Main is done.")