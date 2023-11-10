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



# Maintainer
- [Iman Mohammadi](https://github.com/Imanm02)