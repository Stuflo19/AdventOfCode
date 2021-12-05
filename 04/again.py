file_name = "04\input.txt"
counter = 0
numbers = []
nums = []
bingo_in = 0
grid = [ [[]] ]
sum = 0
lastCall = 0
gridcounter = 0
gridline = 0

with open(file_name, 'r') as input:
    lines = input.readlines()

    for line in lines:
        if counter == 0:
            line.strip().split(",")

            for num in line:
                numbers.append(num)
        elif line == "\n":
            print(grid[gridcounter])
            gridcounter += 1
            gridline = 0
        else:
            grid.insert(0, line.strip().split())
            gridline += 1
        
        counter += 1

print("\nMin: ")
print(min)
print("\nSum: ")
print(sum)
print("\nLast call: ")
print(lastCall)
print("\nFinal Answer: ")
print(int(lastCall)*int(sum))