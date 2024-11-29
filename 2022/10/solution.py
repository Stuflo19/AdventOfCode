with open("10/input.txt", "r") as f:
    lines = f.readlines()
    
def part1():
    cycles = 0
    total = 0
    x_reg = 1
    desired_cycles = 20
    adding = False
    
    for line in lines:
        line = line.strip().split(' ')
        if(line[0] == 'addx'):
            cycles += 2
            if cycles >= desired_cycles: adding = True
            else: x_reg += int(line[1])
        elif(line[0] == 'noop'):
            cycles += 1
        
        if(cycles >= desired_cycles):
            total += x_reg * desired_cycles
            desired_cycles += 40
            if adding: x_reg += int(line[1])
            adding = False
            # print(x_reg)
    print("Part 1: ", total)
    
    
def part2():
    print("Part 2: ")
    cycles = 0
    current_screen = ['.' for _ in range(40)]
    x_reg = 1
    desired_cycles = 40
    adding = False
    check = lambda x: True if x == x_reg-1 else True if x == x_reg else True if x == x_reg+1 else False
    
    for line in lines:
        line = line.strip().split(' ')
        if(line[0] == 'addx'):
            if(check(cycles)):
                current_screen[cycles] = '#'
            if(check(cycles+1)):
                current_screen[cycles+1] = '#'
            
            cycles += 2
            if cycles > desired_cycles: adding = True
            else: x_reg += int(line[1])
        elif(line[0] == 'noop'):
            if(check(cycles)):
                current_screen[cycles] = '#'
            cycles += 1
            # print(current_screen)
        
        if(cycles >= desired_cycles):
            for pixel in current_screen:
                print(pixel, end='')
            print('')
            current_screen = ['.' for _ in range(40)]
            cycles = 0
            if adding: 
                if(check(cycles)):
                    current_screen[cycles] = '#'
                cycles += 1
                x_reg += int(line[1])
            adding = False
    
if __name__ == '__main__':
    part1()
    part2()