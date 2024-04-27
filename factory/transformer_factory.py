from factory.factory_interface import Factory
from jobs.transformers.csv_transform_job import CsvTransformJob
from jobs.transformers.lsq_transform_job import ProjectionTransform, ColumnSpecificTransform
from jobs.transformers import lsq_test_transform_job

class TransformerFactory(Factory):
    def __init__(self):
        super()
        
    def create(self, mode,**kwargs):
        
        merged_config = self.get_config(mode)
        
        match mode:
            case "transform_csv":
                return CsvTransformJob(config = merged_config)
            case "transform_lsq_projection":
                return ProjectionTransform(config = merged_config)
            case "transform_lsq_test_projection":
                return lsq_test_transform_job.ProjectionTransform(config = merged_config)
            case "transform_lsq_col_specific":
                return ColumnSpecificTransform(config = merged_config)
            case _:
                raise ValueError("Invalid transformer type")