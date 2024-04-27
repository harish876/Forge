from datetime import datetime
import hashlib
import logging
import pytz
from jobs.job_interface import ETLJob
    
class CsvTransformJob(ETLJob):
    def execute(self,data=None):
        data['name_hash'] = hashlib.sha256(data["name"].encode()).hexdigest()
        data['created_at'] = datetime.now(pytz.timezone('Asia/Kolkata'))

        self.set_data_context(data)
        self.next()
