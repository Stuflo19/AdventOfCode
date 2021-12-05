class bingoGrid:
    grid = [ [] ]

    def __init__(self, numbers):
        self.grid = numbers
        print(self.grid)
        self.bingo = False
        self.moves = 0
        self.sum = 0
        self.call = 0

    def check_bingo(self, move, call):
        for row in range(len(self.grid)):  # outer loop  
            for col in range(len(self.grid[row])):  # inner loop
                if self.grid[0][col] == 'F' and self.grid[1][col] == 'F' and self.grid[2][col] == 'F' and self.grid[3][col] == 'F' and self.grid[4][col] == 'F':
                    print("bingo")
                    self.bingo = True
                    self.moves = move
                    self.call = call
                    self.find_sum()
                if self.grid[row][0] == 'F' and self.grid[row][1] == 'F' and self.grid[row][2] == 'F' and self.grid[row][3] == 'F' and self.grid[row][4] == 'F':
                    print("bingo")
                    self.bingo = True
                    self.moves = move
                    self.call = call
                    self.find_sum()

    def find_sum(self):
        self.sum = 0
        for row in range(len(self.grid)):  # outer loop  
            for col in range(len(self.grid[row])):  # inner loop  
                if self.grid[row][col] != 'F':
                    self.sum += int(self.grid[row][col])

    def check_square(self, row, col, numb):
        if(numb in self.grid):
            for row in range(len(self.grid)):  # outer loop  
                for col in range(len(self.grid[row])):  # inner loop
                    if self.grid[row][col] == numb:
                        self.grid[row][col] = 'F'

class bingoSquares:
    def __init__(self, number):
        self.num = number
        self.called = False

file_name = "04\input.txt"
counter = 0
gridcounter = 0
numbers = []
nums = []
grid = [ [] ]
grids = []
sum = 0

min_moves = 1000
min_sum = 0
min_call = 0

with open(file_name, 'r') as input:
    lines = input.readlines()

    for line in lines:
        if (counter == 0):
            line = line.strip().split(",")
            for num in line:
                numbers.append(num)

        else:
            if( line == "\n" and counter > 1):
                for row in range(len(grid)):  # outer loop  
                    for col in range(len(grid[row])):  # inner loop
                        print(grid[row][col], end = " ")
                    print()

                grids.append(bingoGrid(grid))
                grid.clear()
                gridcounter = 0
            else:
                grid.append(line.strip().split())
                gridcounter += 1

        counter += 1

print(counter)
for num in numbers:
    for g in grids:
        for row in range(len(g.grid)):  # outer loop  
            for col in range(len(g.grid[row])):  # inner loop
                g.check_square(row, col, num)
    g.check_bingo(numbers.index(num), num)      
                
for g in grids:
    #print(g.grid)
    if g.moves < min_moves:
        min_moves = g.moves
        min_call = g.call
        min_sum = g.sum

print(min_moves)
print(min_call)
print(min_sum)
print( int(min_sum) * int(min_call) )
