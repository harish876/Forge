from jobs.job_interface import ETLJob

class TransformJsonJob(ETLJob):
	def __init__(self, config):
		super().__init__()

	def execute(self, data=None):
		self.set_data_context('foobar')
		self.next()
