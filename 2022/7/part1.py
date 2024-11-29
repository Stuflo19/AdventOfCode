
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
    
    print(sum(v for v in dict.values() if v <= 100000))
            

with open('7/input.txt', 'r') as f:
    data = f.readlines()

    create_structure(data)

    
