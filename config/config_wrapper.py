import configparser
import os
import logging

class ConfigWrapper(configparser.ConfigParser):
    _instance = None  # Class-level variable to hold the instance

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
        return cls._instance
    
    def __init__(self):        
        super().__init__(os.environ, interpolation=configparser.ExtendedInterpolation())
        dir_name = os.path.dirname(__file__)

        base_ini_file = os.path.join(dir_name, "settings.ini")
        if not os.path.exists(base_ini_file):
            raise Exception(f"The base INI file '{base_ini_file}' was not found in this directory")
        
        logging.info(f"The base INI file '{base_ini_file}' was found in this directory")
        self.read(base_ini_file)

        env = os.environ.get("ENV", None)
        logging.info(f"The environment variable 'ENV' has the value '{env}'")
        if env is None:
            raise ValueError("The environment variable: 'environment' has not been set") 
        logging.info(f"The current environment is '{env}'")

        environment_specific_ini_file = os.path.join(dir_name, f"settings.{env}.ini")
        if os.path.exists(environment_specific_ini_file):
            logging.info(f"The environment specific INI file '{environment_specific_ini_file}' was found in this directory")
            self.read(environment_specific_ini_file)

    def get_config(self, section, option, default = None):
        """
        Attempts to get a configuration value for the given section and option.
        Returns default if the section or option does not exist and the default arg
        is passed to the method.
        """
        try:
            return self.get(section, option)
        
        except (configparser.NoSectionError, configparser.NoOptionError):
            logging.error(f"Could not get configuration for the section {section} and for option {option}\n. Returning Default Config ${default}")
            return default
        
        except Exception as e:
            logging.error(f"Unexpected error ${e}")
            return default
        
    
    def get_section(self, section):
        """
        Returns all the options in the given section as a dictionary.
        """        
        if self.has_section(section):
            config =  {option: self.get_config(section, option) for option in self.options(section)}
        else:
            config = {}
        
        return config