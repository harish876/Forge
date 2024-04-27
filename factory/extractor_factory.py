from factory.factory_interface import Factory
from jobs.extractors.csv_extract_job import CsvExtractJob
from jobs.extractors.json_file_extract import JsonExtractJob
from jobs.extractors.mssql_extract_job import MssqlExtractJob
from jobs.extractors import json_file_extract


class ExtractorFactory(Factory):
    def __init__(self):
        super()

    def create(self, mode, **kwargs):
        merged_config = self.get_config(mode)

        match mode:
            case "extract_csv":
                return CsvExtractJob(config = merged_config)
            case "extract_json":
                return JsonExtractJob(config = merged_config)
            case "extract_lsq_test_json":
                return JsonExtractJob(config = merged_config)
            case "extract_mssql":
                return MssqlExtractJob(config = merged_config)
            case _:
                raise ValueError("Invalid extract type")
