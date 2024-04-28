from jobs.job_interface import ETLJob

class TransformJsonJob(ETLJob):
    def __init__(self, config):
        super()

    def execute(self, data=None):
        print(data)
        self.set_data_context(data)
        self.next()
