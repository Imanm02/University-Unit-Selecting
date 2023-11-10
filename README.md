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

# Maintainer
- [Iman Mohammadi](https://github.com/Imanm02)