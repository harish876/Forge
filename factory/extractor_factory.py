from factory.factory_interface import Factory
from jobs.extractors.extract_json_job import ExtractJsonJob

class Factory(Factory):
	def __init__(self):
		super().__init__()

	def create(self, mode, **kwargs):
		merged_config = self.get_config(mode)

		match mode:
			case "extract_json_job":
				return ExtractJsonJob(config=merged_config)
			case _:
				raise ValueError("Invalid extract type")
