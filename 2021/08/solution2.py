
def open_file():
    file_name = "08/input.txt"

    with open(file_name) as input:
        lines = input.readlines()

        return lines.copy()

def find_values(lines, total):
    for line in lines:
        past_pipe = False
        output = {}
        puzzle = {}
        after_pipe = []
        line = line.strip().split(" ")

        for chars in line:

            if len(chars) == 3 or len(chars) == 4 or len(chars) == 2 or len(chars) == 7:
                output[len(chars)] = chars
            elif len(chars) in output:
                output[len(chars)].append(chars)
            else:
                output[len(chars)] = [chars]

            if chars == "|":
                past_pipe = True
                continue
            
            if past_pipe:
                after_pipe.append(chars)

        puzzle = create_figure(output)

        result = create_values(puzzle, after_pipe)

        print(result)

        total += int(result)
        # print("top: " + str( puzzle['top'] ))
        # print("top left: " + str( puzzle['top_left'] ))
        # print("top right: " + str( puzzle['top_right'] ))
        # print("middle: " + str( puzzle['middle'] ))
        # print("bottom left: " + str( puzzle['bottom_left'] ))
        # print("bottom right: " + str( puzzle['bottom_right'] ))
        # print("bottom: " + str( puzzle['bottom'] ))
    return total

def create_figure(output):
    puzzle = {}

    for char in list(output[3]):
        #print(char  + " : " + output[2])
        if char not in output[2]:
            puzzle['top'] = char
        
    for nums in output[6]:
        not_in = 0
        matched = 0

        for match in list(output[4]):
            if nums.find(match) == -1:
                not_in += 1 

        for match in list(output[2]):
            #print(match + " : " + nums)
            if nums.find(match) >= 0:
                matched += 1

        #print(matched)
        # print(output[2] + " : " + nums)
        if not_in == 1 and matched == 2:
            #print("matched")
            for char in list(output[4]):
                if char in nums and char not in output[3]:
                    puzzle['top_left'] = char
                if char not in nums:
                    puzzle['middle'] = char
    
    for nums in output[5]:
        if puzzle['top'] in nums and puzzle['middle'] in nums and puzzle['top_left'] in nums:
            for char in list(nums):
                if char in output[2]:
                    puzzle['bottom_right'] = char

                    for leftover in list(output[2]):
                        if leftover != puzzle['bottom_right']:
                            puzzle['top_right'] = leftover
            
            for char in list(nums):
                if char not in puzzle['top'] and char not in puzzle['middle'] and char not in puzzle['top_left'] and char not in puzzle['top_right'] and char not in puzzle['bottom_right']:
                    puzzle['bottom'] = char
    
    for nums in output[6]:
        if puzzle['middle'] not in nums:
            for char in nums:
                if char not in puzzle['top'] and char not in puzzle['bottom'] and char not in puzzle['middle'] and char not in puzzle['top_left'] and char not in puzzle['top_right'] and char not in puzzle['bottom_right']:
                    puzzle['bottom_left'] = char

    return puzzle

def create_values(puzzle, after_pipe):
    result = ""

    for vals in after_pipe:
        if len(vals) == 2:
            result += '1'
        elif len(vals) == 3:
            result += '7'
        elif len(vals) == 4:
            result += '4'
        elif len(vals) == 5:
            if puzzle['top_right'] not in vals:
                result += '5'
            elif puzzle['bottom_right'] not in vals:
                result += '2'
            else:
                result += '3'
        elif len(vals) == 6:
            if puzzle['middle'] not in vals:
                result += '0'
            elif puzzle['top_right'] not in vals:
                result += '6'
            else:
                result += '9'
        elif len(vals) == 7:
            result += '8'

    return result

def main():
    total = 0

    lines = open_file()

    result = find_values(lines, total)

    print(result)


if __name__ == "__main__":
    main()

