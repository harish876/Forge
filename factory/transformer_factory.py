from factory.factory_interface import Factory
from jobs.transformers.transform_csv_job import TransformCsvJob

class TransformerFactory(Factory):
    def __init__(self):
        super()
        
    def create(self, mode,**kwargs):
        
        merged_config = self.get_config(mode)
        
        match mode:
            case "transform_csv":
                return TransformCsvJob(config = merged_config)
            case _:
                raise ValueError("Invalid transformer type")