from config.config_wrapper import ConfigWrapper
import logging
from factory.extractor_factory import ExtractorFactory
from factory.transformer_factory import TransformerFactory
from factory.loader_factory import LoaderFactory
from factory.reporter_factory import ReporterFactory
from jobs.job_interface import ETLJob

class Assembler():
    def __init__(self):
        self.config_parser: ConfigWrapper = ConfigWrapper()
        self.extractor = ExtractorFactory()
        self.transformer = TransformerFactory()
        self.loader = LoaderFactory()
        self.reporter = ReporterFactory()

    def assemble(self, job_name):
        self.handler: ETLJob = self.yield_job_chain(job_name=job_name)

    def yield_job_chain(self, job_name: str):

        job_steps = self.config_parser.get_config(job_name, "steps", "").split(",")
        handler_head = None
        handler_tail = handler_head

        for step in job_steps:
            step_type = self.config_parser.get_config(step,"type")
            match step_type:
                case "extractor":
                    if handler_head is None:
                        handler_head = self.extractor.create(step)
                        handler_tail = handler_head
                    else:
                        extract_job = self.extractor.create(step)
                        handler_tail.set_next_job(extract_job)
                        handler_tail = extract_job

                case "transformer":
                    transform_job = self.transformer.create(step)
                    handler_tail.set_next_job(transform_job)
                    handler_tail = transform_job

                case "loader":
                    load_job = self.loader.create(step)
                    handler_tail.set_next_job(load_job)
                    handler_tail = load_job

                case "reporter":
                    report_job = self.reporter.create(step)
                    handler_tail.set_next_job(report_job)
                    handler_tail = report_job

        return handler_head

    def start(self):
        try:
            self.handler.execute()
        except Exception as e:
            logging.error(e)
            raise e
