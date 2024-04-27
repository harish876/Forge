from factory.factory_interface import Factory
from jobs.loaders.log_load_job import LogLoadJob
from jobs.loaders.lsq_load_job import LsqLoadJob
from jobs.loaders.talisma_load_job import TalismaLoadJob


class LoaderFactory(Factory):
    def __init__(self):
        super()

    def create(self, mode, **kwargs):
        merged_config = self.get_config(mode)

        match mode:
            case "load_log":
                return LogLoadJob(config = merged_config)
            case "load_lsq":
                return LsqLoadJob(config = merged_config)
            case "load_talisma":
                return TalismaLoadJob(config = merged_config)
            case _:
                raise ValueError("Invalid loader type")
