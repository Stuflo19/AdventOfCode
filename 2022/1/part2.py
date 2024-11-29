with open('input.txt', 'r') as f:
    lines = f.readlines()
    max = []
    curr = 0

    for line in lines:
        if line.strip():
            curr += int(line)
            print(line)
        else:
            print("\n\nempty line")
            if not max or max[0] < curr:
                max.insert(0, curr) 
            curr = 0
    
    print("\n\nMax: ", max)

    print("\nTotal: ", max[0] + max[1] + max[2])
        