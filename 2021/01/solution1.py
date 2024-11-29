file_name = '01\depths.txt'
prevline = 0
count = 0

with open(file_name, 'r') as depths:
    lines = depths.readlines()

    for line in lines:
        if int(line) > int(prevline) and prevline != 0:
            count += 1
            
        prevline = line

print(count)