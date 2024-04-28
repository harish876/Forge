from factory.factory_interface import Factory
from jobs.loaders.loader_json_job import LoaderJsonJob

class Factory(Factory):
	def __init__(self):
		super().__init__()

	def create(self, mode, **kwargs):
		merged_config = self.get_config(mode)

		match mode:
			case "loader_json_job":
				return LoaderJsonJob(config=merged_config)
			case _:
				raise ValueError("Invalid extract type")
