import logging
import os

class CustomLoggerBuilder:
    def __init__(self):
        self.logger = logging.getLogger()
        self.stream_handler = logging.StreamHandler()
        self.file_handler = None

        for handler in logging.root.handlers[:]:
            logging.root.removeHandler(handler)

    def getLogger(self):
        return self.logger
    
    def with_formatter(self,format: str):
        self.formatter = logging.Formatter(format)
        return self

    def with_console_output(self):
        self.stream_handler.setFormatter(self.formatter)
        self.logger.addHandler(self.stream_handler)
        return self

    def with_file_output(self, log_dir: str,file_name: str):
        if not os.path.exists(log_dir):
            os.makedirs(log_dir)
        self.file_handler = logging.FileHandler(os.path.join(log_dir,file_name))
        self.file_handler.setFormatter(self.formatter)
        self.logger.addHandler(self.file_handler)
        return self

    def with_log_level(self, level):
        self.logger.setLevel(level)
        return self

    def build(self):
        return self.logger
