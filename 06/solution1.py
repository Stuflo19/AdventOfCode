file_name = "06\input.txt"
counter = 0
NUM_DAYS = 80

with open(file_name, 'r') as input:
    lines = input.readlines()

    for line in lines:
        line = line.strip().split(',')

    line = list(map(int, line))

    while(counter < NUM_DAYS):
        for num in range(len(line)):
            if line[num] > 0:
                line[num] -= 1
            elif line[num] == 0:
                line[num] = 6
                line.append(8)
        
        counter += 1

print(len(line))
