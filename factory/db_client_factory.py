from factory.factory_interface import Factory
from db.mssql_service import MssqlClient

class DatabaseClientFactory(Factory):
    def __init__(self):
        super()

    def create(self, mode, **kwargs):
        """
        Creates a database client based on the specified mode.

        Args:
            1.mode (str): The mode specifying the type of database client to create.\n
            2.**kwargs: Additional keyword arguments.\n
            3.valid_modes: ["mssql"]

        Returns:
            DatabaseClient: An instance of the appropriate database client based on the mode.

        Raises:
            ValueError: If an invalid database type is provided.
        """

        merged_config = self.get_config(mode)

        match mode:
            case "mssql":
                mssql_client = MssqlClient(config= merged_config)
                return mssql_client.get_client("mssql+pyodbc")
            case _:
                raise ValueError("Invalid database type")
