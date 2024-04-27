from abc import ABC
from config.config_wrapper import ConfigWrapper
config_parser = ConfigWrapper()


class Factory(ABC):
    _instance = None 
    
    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
        return cls._instance
    
    def __init__(self):
        return self

    def get_config(self, step: str):
        
        if step is None:
            raise ValueError("Mode must be specified")
        default_config_section = f"default_{step}_options"
        default_config = config_parser.get_section(default_config_section)
        config = config_parser.get_section(step)

        merged_config = default_config.copy()
        merged_config.update(config)

        return merged_config

    def create(self):
        pass
