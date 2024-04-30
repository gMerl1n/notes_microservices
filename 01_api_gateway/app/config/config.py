import yaml
import os


config_abs_path = "/".join(os.path.abspath(__file__).split('/')[0:-1])

class ConfigLoader:

    def __init__(self) -> None:
        self.path = config_abs_path

    def get_config(self):
        with open(self.path + "/config.yml") as stream:
            try:
              config = yaml.safe_load(stream)
            except yaml.YAMLError as exc:
              raise exc
        return config
    
    def get_login_url(self):
       config = self.get_config()
       return config["auth_service"]["login_url"]

    def get_register_url(self):
       config = self.get_config()
       return config["auth_service"]["register_url"]
    

config = ConfigLoader()