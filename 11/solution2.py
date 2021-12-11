import numpy as np

def open_file():
    file_name = "11/input.txt"

    with open(file_name, 'r') as input:
        squids = np.array([list(line.strip()) for line in (input.readlines())]).astype(int)
        
        return squids

#function that finds how many flashes are caused by a single squid flash
def chain_reaction(squids):

    if np.any(squids > 9):
        for row in len(squids):
            for col in len(squids[row]):
                effected_squids = squids[row-1:row+2,col-1:col+2]

    return chain_reaction(squids) #returns how many squids have flashed

def main():
    total = 0

    squids = open_file() #reads in the lines from the input file

    #loops in the range 0, 100 (task 2 uses range 1, x)
    for step in range(3):
        
        print(squids)

        squids = squids + 1

        #calculates the total from each step and adds it to the current total
        total += 1

    print(total)

if __name__ == "__main__":
    main()