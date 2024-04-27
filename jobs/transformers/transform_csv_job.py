from datetime import datetime
import hashlib
import pytz
from jobs.job_interface import ETLJob
    
class TransformCsvJob(ETLJob):
    def execute(self,data=None):
        self.set_data_context(data)
        self.next()
