from abc import ABC, abstractmethod
import logging


class ETLJob(ABC):

    def __init__(self,**kwargs):
        """
            kwargs takes in a config object defined from the config file.
            Kwargs.get method returns the config value or a None 
        """
        logging.debug("Using default Implementation of Constructor")
        self.next_job: ETLJob | None = None

    def set_next_job(self, job):
        logging.debug("Using default Implementation of set_next_handler")
        self.next_job = job

    def next(self,data=None):
        if self.next_job:
            if data is not None:
                self.next_job.next(data)
            else:
                self.next_job.execute(self.__data)
        else:
            logging.info("End of current job chain")

    def set_data_context(self, data=None):
        self.__data = data
        
    @abstractmethod
    def execute(self,data=None):
        self.set_data_context(self.__data)
        pass
