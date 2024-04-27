# Use the official Python 3.11 image
FROM python:3.11-buster

# Set the working directory in the container
WORKDIR /app

# Copy the requirements file into the container at /app
COPY requirements.txt /app

# Install dependencies
RUN pip install -r requirements.txt

# Copy the current directory contents into the container at /app
COPY . /app

# Run main.py when the container launches
CMD ["python", "main.py"]