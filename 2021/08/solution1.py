
def open_file():
    file_name = "08/input.txt"

    with open(file_name) as input:
        lines = input.readlines()

        return lines.copy()

def find_values(lines):
    output = {}
    past_pipe = False

    for line in lines:
        line = line.strip().split(" ")
        print(line)

        for chars in line:
            if chars == "|":
                past_pipe = True
                continue
            
            if past_pipe:
                print(len(chars))
                if len(chars) == 3 or len(chars) == 4 or len(chars) == 2 or len(chars) == 7:
                    if len(chars) in output:
                        output[len(chars)] += 1
                    else:
                        output[len(chars)] = 1
    
        past_pipe = False
        print(output)
    return sum(output.values())


def main():
    lines = open_file()

    result = find_values(lines)

    print(result)


if __name__ == "__main__":
    main()

