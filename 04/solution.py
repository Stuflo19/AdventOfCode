file_name = "04\input.txt"
counter = 0
numbers = []
nums = []
bingo_in = None
min = 1000
grid = [ [],[] ]
sum = 0
currentMin = False
lastCall = 0
bingo = False

with open(file_name, 'r') as input:
    lines = input.readlines()

    for line in lines:
        if (counter == 0):
            line = line.split(",")

            for num in line:
                numbers.append(num)

            counter += 1
        else:
            if line == "\n":
                for num in numbers:
                    for row in range(len(grid)):  # outer loop  
                        for col in range(len(grid[row])):  # inner loop  
                            if grid[row][col] == num and bingo is False:
                                print("im setting value to F")
                                grid[row][col] = 'F'
                                print(grid[row][col])
                            if grid[row][col] == 'F' and bingo is False:
                                if row+4 < len(grid):
                                    if grid[row+1][col] == 'F' and grid[row+2][col] == 'F' and grid[row+3][col] == 'F' and grid[row+4][col] == 'F':
                                        bingo_in = numbers.index(num)
                                        bingo = True
                                        if bingo_in < min:
                                            min = bingo_in
                                            lastCall = num
                                            currentMin = True


                                if col+4 < len(grid[row]):
                                    if grid[row][col+1] == 'F' and grid[row][col+2] == 'F' and grid[row][col+3] == 'F' and grid[row][col+4] == 'F':
                                        bingo_in = numbers.index(num)
                                        bingo = True
                                        if bingo_in < min:
                                            min = bingo_in
                                            lastCall = num
                                            currentMin = True
                if currentMin:
                    sum = 0
                    for row in range(len(grid)):  # outer loop  
                        for col in range(len(grid[row])):  # inner loop  
                            if grid[row][col] is not 'F':
                                sum += int(grid[row][col])
                counter += 1
                grid.clear()
                currentMin = False
                bingo = False
            else:
                grid.append(line.strip().split())
                print(grid)

print("\nMin: ")
print(min)
print("\nSum: ")
print(sum)
print("\nLast call: ")
print(lastCall)
print("\nFinal Answer: ")
print(int(lastCall)*int(sum))