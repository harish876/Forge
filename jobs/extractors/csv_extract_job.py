import logging
import pandas as pd
from jobs.job_interface import ETLJob


class CsvExtractJob(ETLJob):

    def __init__(self, config):
        super()
        self.__filename =config.get("filename")
        self.delimiter = config.get("delimiter")

    def execute(self, data=None):
        try:
            if self.__filename is None:
                return

            data = pd.read_csv(self.__filename)
            for _, row in data.iterrows():
                self.set_data_context(row)
                self.next()

        except Exception as e:
            logging.error(e)
            raise e

        if data.empty:
            return
