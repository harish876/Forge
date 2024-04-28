from jobs.job_interface import ETLJob

class LoadJsonJob(ETLJob):
	def __init__(self, config):
		super()

	def execute(self, data=None):
		print(data)
