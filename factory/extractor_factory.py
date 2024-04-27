from factory.factory_interface import Factory
from jobs.extractors.extract_csv_job import ExtractCsvJob


class ExtractorFactory(Factory):
    def __init__(self):
        super()

    def create(self, mode, **kwargs):
        merged_config = self.get_config(mode)

        match mode:
            case "extract_csv":
                return ExtractCsvJob(config = merged_config)
            case _:
                raise ValueError("Invalid extract type")
