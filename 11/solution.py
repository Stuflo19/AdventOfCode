import numpy as np

def open_file():
    file_name = "11/input.txt"

    with open(file_name, 'r') as input:
        lines = input.readlines()

        return lines.copy()

#Takes in all lines from file and adds them to 2D numpy array
def create_grid(lines):
    global squids
    created = False

    for line in lines:
        line = list(map(int, line.strip()))

        if not created:
            created = True
            squids = np.array([line])
        else:
            squids = np.concatenate((squids, [line]))

#goes through all squids to find if any have a charge greater than 9 and if so calculates the impact of their flash
def check_squids(step):
    total = 0
    flashes_caused = 0
    global flashed
    flashed = np.zeros((len(squids), len(squids[0]))) #Array used to figure out which squids have already flashed and cannot gain charge this turn

    for row in range(len(squids)):
        for col in range( len(squids[row]) ):
            if flashed[row][col] == 0:
                squids[row][col] += 1

            if(squids[row][col] > 9):
                flashes_caused = chain_reaction(row, col) #Calculates the flashes caused by squid that is flashes
                #Task 2, Finds occurences when all squids flash at the same time
                if flashes_caused == ( len(squids) * len(squids[0]) ):
                    print("\nBig flash at: " + str(step + 1))

                total += flashes_caused #Finds the total number of flashes from each step
    
    return total

#function that finds how many flashes are caused by a single squid flash
def chain_reaction(row,col):
    total = 1
    squids[row][col] = 0 #reset this squids charge
    flashed[row][col] = 1 #set it as a squid that has flashed this turn

    #checks if squid below has a charge greater than 9
    if row != 0 and flashed[row-1][col] == 0:
        squids[row-1][col] += 1

        if squids[row-1][col] > 9:
            total += chain_reaction(row-1, col)

    #checks if squid above has charge greater than 9
    if row != ( len(squids) - 1) and flashed[row+1][col] == 0:
        squids[row+1][col] += 1

        if squids[row+1][col] > 9:
            total += chain_reaction(row+1, col)

    #checks if squid on the left has charge greater than 9
    if col != 0 and flashed[row][col-1] == 0:
        squids[row][col-1] += 1

        if squids[row][col-1] > 9:
            total += chain_reaction(row, col-1)

    #checks if squid on the right has charge greater than 9
    if col != ( len(squids[row]) - 1 ) and flashed[row][col+1] == 0:
        squids[row][col+1] += 1

        if squids[row][col+1] > 9:
            total += chain_reaction(row, col+1)

    #Edge cases
    #checks if bottom right squid has charge greater than 9
    if ( col != ( len(squids[row]) - 1 ) and row != 0 ) and flashed[row-1][col+1] == 0:
        squids[row-1][col+1] += 1

        if squids[row-1][col+1] > 9:
            total += chain_reaction(row-1, col+1)
    
    #checks if top right squid has charge greater than 9
    if ( col != ( len(squids[row]) - 1 ) and row != (len(squids) - 1) ) and flashed[row+1][col+1] == 0:
        squids[row+1][col+1] += 1

        if squids[row+1][col+1] > 9:
            total += chain_reaction(row+1, col+1)
    
    #checks if top left squid has charge greater than 9
    if ( col != 0 and row != (len(squids) - 1) ) and flashed[row+1][col-1] == 0:
        squids[row+1][col-1] += 1

        if squids[row+1][col-1] > 9:
            total += chain_reaction(row+1, col-1)
    
    #checks if bottom left squid has a charge greater than 9
    if ( col != 0 and row != 0 ) and flashed[row-1][col-1] == 0:
        squids[row-1][col-1] += 1

        if squids[row-1][col-1] > 9:
            total += chain_reaction(row-1, col-1)
    
    return total #returns how many squids have flashed

def main():
    total = 0

    lines = open_file() #reads in the lines from the input file

    create_grid(lines) #uses the line to create a global 2D squid array

    #loops in the range 0, 100 (task 2 uses range 1, x)
    for step in range(3):
        print(squids)
        #calculates the total from each step and adds it to the current total
        total += check_squids(step)

    print(total)

if __name__ == "__main__":
    main()