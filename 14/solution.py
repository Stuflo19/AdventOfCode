def open_file(file_name):
    with open(file_name, 'r') as input:
        lines = [list(line.strip().split(' -> ')) for line in (input.readlines())]
        return lines

def get_polymer(lines):
    count = {}
    rules = {}
    steps = 0
    polymer = ''

    for line in lines:
        if len(line) == 1 and polymer == '':
            polymer = list(line[0])
        elif len(line) != 1:
            rules[line[0]] = line[1]

    while steps < 10:
        next_char = 0
        for char in polymer[:-1]:
            next_char += 1
            current_combo = char + polymer[next_char]
            if current_combo in rules:
                polymer.insert(next_char, rules[current_combo])
                next_char += 1

        steps += 1
    
    for char in polymer:
        if char not in count:
            count[char] = 1
        else:
            count[char] += 1

    max = 0
    min = count['N']
    for char in count:
        if count[char] < min:
            min = count[char]
        if count[char] > max:
            max = count[char]

    print("Min: " + str(min))
    print("Max: " + str(max))
    print("Sum: " + str(max - min))

def main():
    lines = open_file("14/input.txt")
    get_polymer(lines)

if __name__ == '__main__':
    main()
