from factory.factory_interface import Factory
from jobs.reporters.lsq_reporter_job import LsqReporterJob

class ReporterFactory(Factory):
    def __init__(self):
        super()
        
    def create(self, mode,**kwargs):
        
        merged_config = self.get_config(mode)
        
        match mode:
            case "report_lsq":
                return LsqReporterJob(config = merged_config)
            case _:
                raise ValueError("Invalid Reporter type")