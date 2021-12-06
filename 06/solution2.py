#attempting to better understand more efficient code by basing it on the code from https://gist.github.com/JoaoNunoAbreu/574025462876ca88acdd9c09f22ab41c

file_name = "06\input.txt"
NUM_DAYS = 256

def open_file():
    with open(file_name, 'r') as input:
        lines = input.readlines()

        for line in lines:
            line = line.strip().split(',')

        line = list(map(int, line))

    return line

def fill_dictionary(line):
    new_dict = {}

    for num in line:
        if num in new_dict:
            new_dict[num] += 1
        else:
            new_dict[num] = 1

    return new_dict

def new_dictionary(old_dict):
    new_dict = {}

    for num in old_dict:
        if num == 0 and 6 not in new_dict and 8 not in new_dict:
            new_dict[6] = old_dict[num]
            new_dict[8] = old_dict[num]
        elif num == 0 and 6 not in new_dict and 8 in new_dict:
            new_dict[6] = old_dict[num]
            new_dict[8] += old_dict[num]
        elif num == 0 and 6 in new_dict and 8 not in new_dict:
            new_dict[6] += old_dict[num]
            new_dict[8] = old_dict[num]
        elif num == 0 and 6 in new_dict and 8 in new_dict:
            new_dict[6] += old_dict[num]
            new_dict[8] += old_dict[num]
        else:
            if num - 1 in new_dict:
                new_dict[num-1] += old_dict[num]
            else:
                new_dict[num-1] = old_dict[num]

    return new_dict

def main():
    with open(file_name, 'r') as input:
        lines = input.readlines()

        for line in lines:
            line = line.strip().split(',')

        line = list(map(int, line))

    dict = fill_dictionary(line)

    for i in range(NUM_DAYS):
        dict = new_dictionary(dict)
    print(dict)
    print(sum(dict.values()))

if __name__ =="__main__":
    main()
