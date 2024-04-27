from dotenv import load_dotenv
try:
    load_dotenv()
    print("Environment variables loaded successfully")
except Exception as e:
    print(f"Error loading .env file: {e}")
    sys.exit(1)

import datetime
import sys
import logging
from assembler.assembler import Assembler
from config.config_wrapper import ConfigWrapper
from utils.logger import CustomLoggerBuilder


def start_job():
    try:
        logging.info('Started ETL Job...')
        config_parser = ConfigWrapper()
        jobs = config_parser.get_config('job_options', 'jobs', "").split(',')
        assembler = Assembler()
        for job in jobs:
            try:
                assembler.assemble(job_name=job)
                assembler.start()
                logging.info(f"ETL job executed for ${job}")
            except Exception as e:
                logging.exception(
                    f"An error occurred while processing job {job}: {e}")

    except Exception as e:
        logging.exception(f"Job {job} failed to execute due to", e)
        raise e


def main():
    current_date = datetime.datetime.now()
    logger = CustomLoggerBuilder() \
        .with_formatter('[%(asctime)s] [File - %(filename)s - Line: %(lineno)d] %(levelname)s: %(message)s\n')\
        .with_console_output()\
        .with_file_output('logs', f'{current_date.strftime("%d-%m-%Y")}-log.log')\
        .with_log_level(logging.NOTSET)\
        .build()
    try:
        logger.info("Starting Job...")
        start_job()
        sys.exit(0)
    except Exception as e:
        logger.exception(e)
        sys.exit(1)


if __name__ == "__main__":
    main()
