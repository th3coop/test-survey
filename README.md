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

### Sample log output
For your reference, this is what I get in **log.log** when I run `python test.py`:
```
-- Executing 'go run main.go' --

Enter text: Coool kids
Coool kids

[1;92m? [0m[1;99mPick one[0m  [36m[Use arrows to move, space to select, type to filter][0m
[1;36m> red[0m
[1;99m  blue[0m
Error: Incorrect function.
```
