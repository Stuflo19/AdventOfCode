filename = "4/input.txt"

# 2-8,3-7 
# 3>2
# 7<8

# 2-4,6-8
# 

with open(filename, 'r') as f:
    lines = f.readlines()
    total = 0

    for line in lines:
        line = line.split(',')
        elf1 = line[0].strip().split('-')
        elf2 = line[1].strip().split('-')

        # print("elf 1: ", elf1)
        # print("elf 2: ", elf2)

        # if int(elf1[0]) >= int(elf2[0]) and int(elf1[1]) <= int(elf2[1]):
        #     # print("=========================")
        #     # print("elf 1: ", elf1)
        #     # print("elf 2: ", elf2)
        #     # print("=========================")
        #     total += 1

        # elf1[0] less than elf2[0] and elf1[0] is greater than elf2[1] so that elf2 has overlap?
        if int(elf1[0]) <= int(elf2[0]) and int(elf1[1]) >= int(elf2[0]):
            print("=========================")
            print("elf 1: ", elf1)
            print("elf 2: ", elf2)
            print("=========================")
            total += 1
        elif int(elf2[0]) <= int(elf1[0]) and int(elf2[1]) >= int(elf1[0]):
            print("=========================")
            print("elf 1: ", elf1)
            print("elf 2: ", elf2)
            print("=========================")
            total += 1

    print("============================")
    print("total: ", total)
    print("============================")