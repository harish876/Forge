from factory.factory_interface import Factory

class ReporterFactory(Factory):
    def __init__(self):
        super()
        
    def create(self, mode,**kwargs):

        merged_config = self.get_config(mode)
        
        match mode:
            case _:
                raise ValueError("Invalid Reporter type")