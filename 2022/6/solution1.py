filename = "6/input.txt"

def solution(part, size):
    count = 0

    while True:
        errors = 0
        # print(line[count:count+4])
        currpos = line[count:count+size]
        for char in currpos:
            # print(char)
            # print(currpos.count(char))
            if currpos.count(char) > 1:
                errors += 1
                count += 1
                break
        
        if errors == 0:
            break
    
    print("=============================")
    print("Part: " + part)
    print(count+size)
    print("=============================")

with open(filename, 'r') as f:
    line = f.readline()

    solution(part="1", size=4)
    solution(part="2", size=14)

        
