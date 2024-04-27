from factory.factory_interface import Factory
from jobs.loaders.load_csv_job import LoadCsvJob


class LoaderFactory(Factory):
    def __init__(self):
        super()

    def create(self, mode, **kwargs):
        merged_config = self.get_config(mode)

        match mode:
            case "load_csv":
                return LoadCsvJob(config = merged_config)
            case _:
                raise ValueError("Invalid loader type")
