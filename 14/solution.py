def open_file(file_name):
    with open(file_name, 'r') as input:
        lines = [list(line.strip().split(' -> ')) for line in (input.readlines())]
        return lines

#part 1
    # def get_polymer(lines):
    #     count = {}
    #     rules = {}
    #     steps = 0
    #     polymer = ''

    #     for line in lines:
    #         if len(line) == 1 and polymer == '':
    #             polymer = list(line[0])
    #         elif len(line) != 1:
    #             rules[line[0]] = line[1]

    #     while steps < 10:
    #         next_char = 0
    #         for char in polymer[:-1]:
    #             next_char += 1
    #             current_combo = char + polymer[next_char]
    #             if current_combo in rules:
    #                 polymer.insert(next_char, rules[current_combo])
    #                 next_char += 1

    #         steps += 1
        
    #     for char in polymer:
    #         if char not in count:
    #             count[char] = 1
    #         else:
    #             count[char] += 1

    #     max = 0
    #     min = count['N']
    #     for char in count:
    #         if count[char] < min:
    #             min = count[char]
    #         if count[char] > max:
    #             max = count[char]

    #     print("Min: " + str(min))
    #     print("Max: " + str(max))
    #     print("Sum: " + str(max - min))

def new_dictionary(pairs, rules):
    new_dict = {}

    for pair in pairs:
        if pair in rules:
            new_pair = [pair[0] + rules[pair], rules[pair] + pair[1]]
            if new_pair[0] not in new_dict:
                new_dict[new_pair[0]] = pairs[pair]
            else:
                new_dict[new_pair[0]] += pairs[pair]
            
            if new_pair[1] not in new_dict:
                new_dict[new_pair[1]] = pairs[pair]
            else:
                new_dict[new_pair[1]] += pairs[pair] 
    
    del pairs
    return new_dict

def find_appearances(pairs):
    dictlen = len(pairs)
    dictpos = 0
    individual_chars = {}

    for pair in pairs:
        if pair[0] not in individual_chars:
            individual_chars[pair[0]] = pairs[pair]
        else:
            individual_chars[pair[0]] += pairs[pair]
    
    return individual_chars

#Part 2, using a dictionary to count occurences
def get_polymer(lines):
    pairs = {}
    rules = {}
    individual_chars = {}
    steps = 0

    for line in lines:
        if len(line) == 1 and line != '':
            polymer = list(line[0])
            next_char = 1
            for pair in polymer[:-1]:
                pair = pair + polymer[next_char]
                if pair not in pairs:
                    pairs[pair] = 1
                else:
                    pairs[pair] += 1
                
                next_char += 1

        elif len(line) != 1:
            rules[line[0]] = line[1]
    
    print(pairs)

    while steps < 40:
        pairs = new_dictionary(pairs, rules)
        steps += 1

    individual_chars = find_appearances(pairs)

    maximum = max(individual_chars.values())
    minimum = min(individual_chars.values())

    print("Growth: " + str(maximum - minimum))
    # print(pairs)
    # print(rules)
    # print(individual_chars)
        

def main():
    lines = open_file("14/input.txt")
    get_polymer(lines)

if __name__ == '__main__':
    main()
