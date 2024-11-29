filename = "3/input.txt"
from math import trunc
import string

with open(filename, 'r') as f:
    lines = f.readlines()
    total = 0

    for line in lines:
        # Stripts and finds half of the line length
        line=line.strip()
        half = trunc(len(line)/2)
        
        # Splits line into two parts
        part1 = set(line[0:half])
        part2 = set(line[half:half*2])
        # print(part1, ":", part2)

        # compare two sets
        match = part1 & part2
        char = match.pop()
        if(char.isupper()):
            # print("Value found:", (string.ascii_lowercase.find(char.lower())+1) + 26)
            total += (string.ascii_lowercase.find(char.lower())+1) + 26
        else:
            # print("Value found:", string.ascii_lowercase.find(char)+1)
            total += string.ascii_lowercase.find(char)+1
        
        # print("Moving total: ", total)
    
    print(total)

