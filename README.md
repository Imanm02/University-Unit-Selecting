# University-Unit-Selecting
This repository contains a Python script designed to automate taking courses in the Sharif University of Technology selecting units process.

## Prerequisites
- Python 3.x
- `requests` library

## Code Walkthrough

This script is designed to automate the process of course registration at Sharif University of Technology. Below is a detailed walkthrough of the script components and functionalities.

### Environment Setup

Before the script starts, it loads necessary environment variables:

```python
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()
```

`load_dotenv()` reads the `.env` file, which should be located in the same directory as the script, and loads the variables into the environment. This is where the script gets sensitive information, such as the `Authorization` token.

### Setting Up Headers

The script sets up the headers for the HTTP request:

```python
site_headers = {
    'User-Agent': '...',
    'Accept': 'application/json',
    ...
    'Authorization': os.getenv('AUTHORIZATION_TOKEN'),
    ...
}
```

`site_headers` includes the `Authorization` header, which uses the `AUTHORIZATION_TOKEN` from the environment variables.

### Courses to Register

The script defines the courses to register for:

```python
courses = [('40254-1', '3'), ('40124-3', '3')]
```

This list of tuples holds the course IDs and the corresponding units.

### Registration Function

The `register_course` function handles the API request to register for a course:

```python
def register_course(course_code, units):
    ...
    response = requests.post('https://my.edu.sharif.edu/api/reg', headers=site_headers, json=site_data)
    return response.json()
```

It constructs the data payload with the course details and sends a POST request to the university's course registration API.

### Main Registration Loop

The script runs a loop to attempt course registration:

```python
while courses:
    for course in courses.copy():
        ...
        response = register_course(course[0], course[1])
        ...
```

It iterates over a copy of the `courses` list, trying to register for each course. After each attempt, the script sleeps for a short period to prevent rapid-fire requests.

### Handling Responses

The script checks the API response after each registration attempt:

```python
if response['jobs'][0]['result'] == 'OK':
    print(f"{course[0]} registered successfully.")
    courses.remove(course)
else:
    print(f"Couldn't register {course[0]}. ERROR: {response['jobs'][0]['result']}")
```

If a course is successfully registered, it's removed from the list. Otherwise, an error message is displayed.

### Error Handling

In case of exceptions, the script captures and prints the error:

```python
except Exception as e:
    print(f"An error occurred: {e}")
```

This ensures that any issues during the request process are logged.

### Request Throttling

To manage the load on the server, the script waits for 5 seconds between registration attempts:

```python
time.sleep(5)
```

This delay helps to avoid overwhelming the server or hitting rate limits.

This automation script streamlines the course registration process. However, users should comply with Sharif University's policies regarding automated interactions with their systems.

## Maintainer
- [Iman Mohammadi](https://github.com/Imanm02)