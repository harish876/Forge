from jobs.job_interface import ETLJob

class LoaderJsonJob(ETLJob):
	def __init__(self, config):
		super().__init__()

	def execute(self, data=None):
		self.next()
