# Use python 3.7 as base image
FROM python:3.7

# Update all package sources
RUN apt update

# Change the working directory
WORKDIR /app

# Copy all files
COPY . .

# Install all python requirements
RUN pip3 install -r requirements_bot.txt

# Set the entry point
CMD [ "python", "./peach_bot/main.py" ]
