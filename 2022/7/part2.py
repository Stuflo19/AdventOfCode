
def create_structure(data):
    pwd = []
    dict = {"['/']": 0}

    for line in data:
        line = line.strip()
        if line.startswith('$'):
            if line == '$ ls':
                print("ls")
            elif line.endswith('..'):
                from_dir = pwd[:]
                pwd.pop()
                dict[str(pwd)] += dict[str(from_dir)]
            else:
                pwd.append(line.split()[2])
                dict[str(pwd)] = 0    
        elif line.startswith('dir'):
            new_path = pwd[:]
            new_path.append(line.split()[1])
        else:
            dict[str(pwd)] += int(line.split()[0])
    
    while len(pwd) > 1:
        from_dir = pwd[:]
        pwd.pop()
        dict[str(pwd)] += dict[str(from_dir)]
    print(dict)

    space_required = 70000000 - dict["['/']"]
    space_required = 30000000 - space_required
    print(space_required)
    min = 0
    for val in dict.values():
        if val >= space_required:
            if min == 0:
                min = val
            elif val < min:
                min = val
    
    print(min)
                

with open('7/input.txt', 'r') as f:
    data = f.readlines()

    create_structure(data)

    
