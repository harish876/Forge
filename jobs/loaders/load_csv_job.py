import pandas as pd
from jobs.job_interface import ETLJob

class LoadCsvJob(ETLJob):
    def execute(self,data:pd.Series):
        print(data)
