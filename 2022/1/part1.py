with open('test.txt', 'r') as f:
    lines = f.readlines()
    max = 0
    curr = 0

    for line in lines:
        if line.strip():
            curr += int(line)
            print(line)
        else:
            print("\n\nempty line")
            if(max == 0 or max < curr):
                max = curr
            curr = 0
    
    print("\n\nMax: ", max)
        