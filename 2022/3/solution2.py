filename = "3/input.txt"
from math import trunc
import string
from itertools import islice

with open(filename, 'r') as f:
    total = 0
    while True:
        lines = list(islice(f, 3))
        print(lines)
        if not lines:
            break
        
        part1 = set(lines[0].strip())
        part2 = set(lines[1].strip())
        part3 = set(lines[2].strip())

        checked =  part1 & part2 & part3
        char = checked.pop()
        # print(char)
        if(char.isupper()):
            # print("Value found:", (string.ascii_lowercase.find(char.lower())+1) + 26)
            total += (string.ascii_lowercase.find(char.lower())+1) + 26
        else:
            # print("Value found:", string.ascii_lowercase.find(char)+1)
            total += string.ascii_lowercase.find(char)+1
        
    print(total)

