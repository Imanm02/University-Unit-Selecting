import requests
import os
import time
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

# Set HTTP headers for requests
site_headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:91.0) Gecko/20100101 Firefox/91.0',
    'Accept': 'application/json',
    'Accept-Language': 'en-US,en;q=0.5',
    'Referer': 'https://my.edu.sharif.edu/courses/marked',
    'Content-Type': 'application/json',
    'Authorization': os.getenv('AUTHORIZATION_TOKEN'),  # Load authorization token from environment
    'Origin': 'https://my.edu.sharif.edu',
    'Connection': 'keep-alive',
    'Sec-Fetch-Dest': 'empty',
    'Sec-Fetch-Mode': 'cors',
    'Sec-Fetch-Site': 'same-origin',
    'TE': 'trailers'
}

# List of courses to register for (course code, units)
courses = [('40254-1', '3'), ('40124-3', '3')]

def register_course(course_code, units):
    # Define data payload for course registration
    site_data = {
        'action': 'add',
        'course': course_code,
        'units': units
    }
    # Make a POST request to the registration API endpoint
    response = requests.post('https://my.edu.sharif.edu/api/reg', headers=site_headers, json=site_data)
    # Return the JSON response
    return response.json()

# Continuously attempt to register for courses until list is empty
while courses:
    for course in courses.copy():
        try:
            # Attempt to register for the course
            response = register_course(course[0], course[1])
            # Check if registration was successful
            if response['jobs'][0]['result'] == 'OK':
                print(f"{course[0]} registered successfully.")
                courses.remove(course)
            else:
                print(f"Couldn't register {course[0]}. ERROR: {response['jobs'][0]['result']}")
        except Exception as e:
            # Print error message if an exception occurs
            print(f"An error occurred: {e}")
    # Wait for 5 seconds before trying again
    time.sleep(5)