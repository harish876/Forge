import logging
import pandas as pd
import json
from jobs.job_interface import ETLJob

class LogLoadJob(ETLJob):
    def execute(self,data:pd.Series):
        if isinstance(data, pd.Series):
            for row in data:
                logging.info(json.dumps(row))
        elif isinstance(data, pd.DataFrame):
            print("data is a Pandas DataFrame",data)
        else:
            print("data is neither a Series nor a DataFrame")
