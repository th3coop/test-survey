## Instructions
- be using *Python3* and have *virtualenv* installed
- create a virtualenv 
`virtualenv env`
- activate the virtual env
`env\Scripts\activate.bat`
- install *pexpect*
`pip install -r requirements.txt`
- run the test file
`python test.py`
- checkout the log file that should be created, **log.log**


## Results
When you run the test file you will get and exception reported
from `pexpect` that basically says "We can't find what you're looking for".
It appears that executing *line 10* in **test.py** is causing the 
survey library to crash.
If you comment out everything after *line 10* in `main()`, lines 11 to 13, you will get the same output in **log.log**.