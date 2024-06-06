from jobs.job_interface import ETLJob

# yeah! new file created. Lets check the factory
class ExtractDemoJob(ETLJob):
	def __init__(self, config):
		super().__init__()

	def execute(self, data=None):
		self.set_data_context('foobar')
		self.next()
