# TL;DR

1. Generic ETL Pipeline 
2. Follows chain of responsinbility design pattern
3. Config driven to switch between pandas framework.

# Getting Started

1. Command to generate a requirements.txt ``` pip freeze > requirements.txt```
2. Command to install all requirements  ```pip install -r requirements.txt```
3. Get env variables for accessing mongo db and mssql db data

# Todo

1. Correct problems related to installation. Pymongo is not getting installed properly.
2. Add Asyncio support or make IO operations concurrent.
3. Migrate Web Engage Job to this pipeline.
4. Add Scheduled jobs for report analysis and cleanup.
5. Onboarding Docs, Usage and Examples



