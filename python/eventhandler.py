# import time module, Observer, FileSystemEventHandler
import time
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
import logging
import S3Utils
from datetime import datetime
import os

class OnMyWatch:
    # Set the directory on watch

    def __init__(self):
        self.observer = Observer()

    def run(self):
        watchDirectory = os.getenv("WATCH_DIR") if os.getenv("WATCH_DIR") is not None else "/coredumps"

        event_handler = Handler()
        self.observer.schedule(event_handler, watchDirectory, recursive = True)
        self.observer.start()
        print("begin to watch file, dir : " + watchDirectory)
        try:
            while True:
                time.sleep(5)
        except:
            self.observer.stop()
            print("Observer Stopped")

        self.observer.join()


class Handler(FileSystemEventHandler):

    @staticmethod
    def on_any_event(event):
        bucket_name = "sn-dump-collector"
        pod_name = os.getenv("POD_NAME") if os.getenv("POD_NAME") is not None else "POD_NAME"
        namespace_name = os.getenv("NAMESPACE_NAME") if os.getenv("NAMESPACE_NAME") is not None else "POD_NAME"

        key_name = datetime.today().strftime('%Y-%m-%d') + "_" + namespace_name + "_" + pod_name

        if event.is_directory:
            return None

        elif event.event_type == 'created':
            # Event is created, you can process it now
            print("Watchdog received created event - % s." % event.src_path)
            S3Utils.write_to_s3(bucket=bucket_name, keyname=key_name, file_path=event.src_path)
        elif event.event_type == 'modified':
            # Event is modified, you can process it now
            print("Watchdog received modified event - % s." % event.src_path)
            S3Utils.write_to_s3(bucket=bucket_name, keyname=key_name, file_path=event.src_path)


if __name__ == '__main__':
    watch = OnMyWatch()
    watch.run()
