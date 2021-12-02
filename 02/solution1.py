file_name = '02\directions.txt'
count = 0
direction = ""
amount = 0
horizontal = 0
depth = 0

with open(file_name, 'r') as directions:
    lines = directions.readlines()

    for line in lines:
        line = line.split()
        print(line[0])
        direction = line[0]
        amount = int(line[1])

        if direction == "forward":
            horizontal += amount
        elif direction == "down":
            depth += amount
        elif direction == "up":
            depth -= amount
            
            

print(depth * horizontal)