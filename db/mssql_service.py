import logging
import os
from sqlalchemy import Connection, create_engine
from sqlalchemy.pool import QueuePool
from sqlalchemy.engine import URL
import pymssql


class MssqlClient():
    def __init__(self, config) -> Connection:
        """
        Args:
            config(dict:[ str: str | bool | None ]):\n 
                1. driver: str (example: "ODBC Driver 13 for SQL Server")
                2. trust server certificate (snake cased): str ("yes" or "no")
                3. stream results (snake cased): bool 

        Raises:
            e: Error is raised if ODBC driver is incompatible or the database connection is not established.

        Returns:
            Connection: Returns a Connection Pool Object for a MSSQL connection
        """

        self.conn = None
        self.driver = config.get('driver', 'ODBC Driver 13 for SQL Server')
        self.trust_server_certificate = config.get('trust_server_certificate', "no")
        self.stream_results = bool(config.get('stream_results', False))
        
        self.SERVER = os.getenv("MSSQL_SERVER")
        self.DATABASE = os.getenv("MSSQL_DATABASE")
        self.USERNAME = os.getenv("MSSQL_USERNAME")
        self.PASSWORD = os.getenv("MSSQL_PASSWORD")
        self.PORT = os.getenv("MSSQL_PORT")


    def get_client(self,driver:str):
        """
        Args:
            driver:str - "mssql+pyodbc" / "mssql+pymssql"
        """
        if self.conn is not None:
            return self.conn
        
        query = {}
        if driver == "mssql+pyodbc":
            query["driver"] = self.driver
            query["TrustServerCertificate"] = self.trust_server_certificate
        
        try:
            connection_url = URL.create(
                driver,
                username=self.USERNAME,
                password=self.PASSWORD,
                host=self.SERVER,
                port=self.PORT,
                database=self.DATABASE,
                query=query
            )

            engine = create_engine(connection_url, poolclass=QueuePool)
            conn = engine.connect().execution_options(stream_results=self.stream_results)
            logging.info("Connected to Database Sucessfully")
            self.conn = conn
            return self.conn

        except Exception as e:
            logging.error(f"Error connecting to MSSQL Database - {e}")
            raise e
    
    def get_clientv1(self):
        if self.conn is not None:
            return self.conn
        
        conn = pymssql.connect(
            server=self.SERVER, 
            user=self.USERNAME, 
            password=self.PASSWORD, 
            database=self.DATABASE,
            port=self.PORT
        )
        self.conn = conn
        return conn