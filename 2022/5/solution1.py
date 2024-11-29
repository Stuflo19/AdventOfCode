from collections import OrderedDict

filename = "5/input.txt"
num_cranes = 3

with open(filename, 'r') as f:
    lines = f.readlines()
    crates = {}

    for line in lines:
        line = line.strip('\n').split(' ')
        if len(line) == 1 and line[0] == '':
            continue
        elif len(line) > 1:
            if line[1] == '1' and line[0] != 'move':
                continue
        if line[0] == 'move':
            print(line)
            # print(crates)
            for i in range(0, int(line[1])):
                to = int(line[5])-1
                fr = int(line[3])-1
                if len(crates[fr]) > 0:
                    crates[to].insert(0, crates[fr].pop(0))
            print("-------After--------")
            print("From:", crates[fr])
            print("To: ", crates[to])
            print(crates)

        elif line[0] != 'move':
            i = 0
            j = 0
            print(line)
            while i < len(line):
                if(line[i] != ''):
                    if not j in crates:
                        crates[j] = [line[i]]
                    else:
                        crates[j].append(line[i])
                    i += 1
                else:
                    i += 4
                
                j += 1

    for key in OrderedDict(sorted(crates.items())):
        print(key)
        print("char: ", crates[key][0])            