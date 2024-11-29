filename = "4/input.txt"

with open(filename, 'r') as f:
    lines = f.readlines()
    total = 0

    for line in lines:
        line = line.split(',')
        elf1 = line[0].strip().split('-')
        elf2 = line[1].strip().split('-')

        print("elf 1: ", elf1)
        # print("elf 2: ", elf2)

        if int(elf1[0]) >= int(elf2[0]) and int(elf1[1]) <= int(elf2[1]):
            # print("=========================")
            # print("elf 1: ", elf1)
            # print("elf 2: ", elf2)
            total += 1
            # print("=========================")
        elif int(elf1[0]) <= int(elf2[0]) and int(elf1[1]) >= int(elf2[1]):
            # print("=========================")
            # print("elf 1: ", elf1)
            # print("elf 2: ", elf2)
            total += 1
            # print("=========================")

    print("============================")
    print("total: ", total)
    print("============================")