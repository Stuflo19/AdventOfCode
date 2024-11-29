import math

def open_file():
    file_name = "10/input.txt"

    with open(file_name, "r") as input:
        lines = input.readlines()

        return lines.copy()

def remove_corrupted(lines):
    corrupted = []
    expected = ["{}", '[]', '()', '<>']

    for line in lines:
        full_line = line
        line = list( line.strip() )
        opening = []

        #checks for opening tags
        for syntax in line:
            if syntax == '{' or syntax == '[' or syntax == '(' or syntax == '<':
                opening.insert(0, syntax)
                continue
            
            #if tag found is not an opening tag check if it matches its corresponding opening tag
            #if not then the line must be corrupted
            if (opening[0]+syntax) not in expected:
                corrupted.append(full_line)
                break
            #else the tag is properly closed and can be removed from the list to be checked
            else:
                opening.pop(0)

    return corrupted #return all the lines found to be corrupt

def repair_missing(lines):
    totals = []
    fixes = [')', ']', '}', '>']
    expected = ["{}", '[]', '()', '<>']

    for line in lines:
        line = list( line.strip() )
        opening = []
        corrections = []
        total = 0

        #adds all openings to an array and removes them if a match is founf
        for syntax in line:
            if syntax == '{' or syntax == '[' or syntax == '(' or syntax == '<':
                opening.insert(0, syntax)
                continue
            
            opening.pop(0) #no check is done here as we are convinced there are no errors due to previous check

        #finds which closing bracket is required for each tag that is unclosed
        for opens in opening:
            for fix in fixes:
                if (opens+fix) in expected:
                    corrections.append(fix)
                    break
        
        #calculates the sum for the challenge
        for correction in corrections:
            total = total * 5
            total += (fixes.index(correction) + 1)
        
        totals.append(total) #adds this lines total

    totals.sort() #sorts the array in ascending order
    return totals

def main():
    lines = open_file()

    corrupted = remove_corrupted(lines)

    print(corrupted)

    filtered = [line for line in lines if line not in corrupted]

    solution = repair_missing(filtered)

    middle = math.floor(len(solution)/2)

    print(solution[middle])

if __name__ == '__main__':
    main()
        
