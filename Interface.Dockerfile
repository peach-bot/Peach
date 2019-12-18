# Use python 3.7 as base image
FROM python:3.7

# Expose port 5000 to the world
EXPOSE 5000

# Update all package sources
RUN apt update

# Change the working directory
WORKDIR /app

# Copy all files
COPY . .

# Install all python requirements
RUN pip3 install -r requirements_interface.txt

# Set the entry point
CMD [ "python", "./peach_interface/main.py" ]
