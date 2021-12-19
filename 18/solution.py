def open_file(file_name):
    with open(file_name) as input:
        lines = [list(line.strip()) for line in (input.readlines())]
        print(lines)
    return lines

def line_to_array(line):
    print("\n")
    print(line)
    quit()

def split(curr_sum):
    new_sum = []

def explode(curr_sum):
    new_sum = []

def add(curr_sum):
    new_sum = []

def find_sum(lines):
    curr_sum = []
    magnitude = 0

    for line in lines:
        print(line)
        quit()

        curr_sum.append(line)

    return magnitude

def main():
    lines = open_file('18/input.txt')

    line_to_array(lines)

    find_sum(lines)

if __name__ == '__main__':
    main()