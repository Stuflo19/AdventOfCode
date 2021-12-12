#Trying to learn about slices in python from https://www.reddit.com/user/zalazalaza/ solution to advent of code day 11
import numpy as np

def open_file():
    file_name = "11/input.txt"

    with open(file_name, 'r') as input:
        #pad adds 1 of the constant value around outside of array and asint means the whole array can be treated as 1 large integer
        squids = np.pad(np.array([list(line.strip()) for line in (input.readlines())]).astype(int), pad_width=1, constant_values=-9) 
        return squids

#function that finds how many flashes are caused by a single squid flash
def chain_reaction(squids):
    #checks if any value inside of the array is more than 9
    if np.any(squids > 9):
        #using range 1 to len-1 cuts out checking the padding
        for row in range(len(squids)-1):
            for col in range(1, len(squids[row])-1):
                if squids[row][col] > 9:
                    #create a 3x3 slice of all the values around and including the squid
                    effected_squids = squids[row-1:row+2,col-1:col+2]
                    effected_squids += 1 #increment effected squids by 1
                    effected_squids[effected_squids == 1] = 0 #whenever a squid is only at 1 it must have been turned to 0 this turn so reset to 0

                    #insert the effected slice back into the array
                    squids[row-1:row+2,col-1:col+2] = effected_squids
                    squids[row][col] = 0 #sets the squid that just glowed to have 0 charge

            squids[squids < 0] = -9 # makes sure if the padding was changed that it remains at -9 so it never gets counted as a squid

        return chain_reaction(squids) #call the function again recursively to see if there is any more squids to glow

    return squids #returns the updated array of squids

def main():
    total = 0

    squids = open_file() #reads in the lines from the input file

    #loops in the range 0, 100 (task 2 uses range 1, x)
    for step in range(1, 500):
        #increment grid by 1
        squids += 1

        #calls the chain reaction of squids glowing recursively
        squids = chain_reaction(squids)

        #get total from number of 0's from current iteration
        total += np.sum(squids == 0)

        #checks for total flash by checking the sum of 0's against the array size (-2 as there is padding on both sides)
        if np.all(squids[1:-1,1:-1] == 0):
            print("Total flash at: " + str(step)) #step +1 because of the way the challenge counts
            break
        if step == 100:
            print("Score at 100 steps: " + str(total))

if __name__ == "__main__":
    main()