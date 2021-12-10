def open_file():
    file_name = "10/input.txt"

    with open(file_name, "r") as input:
        lines = input.readlines()

        return lines.copy()

def search_syntax(lines):
    error = {')': 0, ']': 0, '}': 0, '>': 0} #holds the values for each error bracket found
    worth = {')': 3, ']': 57, '}': 1197, '>': 25137} #holds the values of each found error bracket
    expected = ["{}", '[]', '()', '<>'] #holds all the expected tag combinations

    for line in lines:
        line = list( line.strip() )
        opening = []

        for syntax in line:
            #checks for opening tags and adds them to a list
            if syntax == '{' or syntax == '[' or syntax == '(' or syntax == '<':
                opening.insert(0, syntax)
                continue
            
            #if there is no matching error tag then note error causing tag
            if (opening[0]+syntax) not in expected:
                error[syntax] += 1 * worth[syntax] 
                break
            #else correct tag so remove from list
            else:
                opening.pop(0)

    return error

def main():
    lines = open_file()

    error = search_syntax(lines)

    print(sum(error.values()))

if __name__ == '__main__':
    main()
        
