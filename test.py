from pexpect import popen_spawn as p
from pexpect import exceptions as e

def main():
    cmd = "go run main.go"
    child = p.PopenSpawn(cmd)
    child.logfile = Logger(cmd)
    child.expect("Enter text:")
    child.sendline("Coool kids")
    child.expect("Favourite color") # This appears to be the culprit line
    child.sendline("")
    child.expect("Response: red")
    print("Done")

class Logger:

    def __init__(self, cmd):
        self.logfile = open("log.log", "wb")
        self.logfile.write(("-- Executing '%s' --\n\n" % cmd).encode())
        self.logged = ""

    def write(self, s):
        self.logfile.write(s)
        self.logged += s.decode()

    def flush(self):
        self.logfile.flush()

if __name__ == "__main__":
    main()