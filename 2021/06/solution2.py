#attempting to better understand more efficient code by basing it on the code from https://gist.github.com/JoaoNunoAbreu/574025462876ca88acdd9c09f22ab41c

file_name = "06\input.txt"
NUM_DAYS = 256 #number of days to calculate population for

#funtion to read in file for input
def open_file():
    with open(file_name, 'r') as input:
        lines = input.readlines()

        for line in lines:
            line = line.strip().split(',')

        line = list(map(int, line))

        return line.copy()

#function used to create the first dictionary
def fill_dictionary(line):
    new_dict = {} #creating dictionary to store values

    #loops for every number in the line from file
    for num in line:
        #if the number is in the dictionary already
        if num in new_dict:
            new_dict[num] += 1 #add 1 to total of those numbers
        #if number is not already in the dictionary
        else:
            new_dict[num] = 1  #Create the new entry with 1

    return new_dict #return the created dictionary

#function to create a new dictionary for every day
def new_dictionary(old_dict):
    new_dict = {}

    #for every number in the old dictionary
    for num in old_dict:
        #if statements to check if the number is 0 then create the same amount of new 6's and 8's as the old dictionary had 0's
        if num == 0 and 6 not in new_dict and 8 not in new_dict:
            new_dict[6] = old_dict[num]
            new_dict[8] = old_dict[num]
        elif num == 0 and 6 not in new_dict and 8 in new_dict:
            new_dict[6] = old_dict[num]
            new_dict[8] += old_dict[num]
        elif num == 0 and 6 in new_dict and 8 not in new_dict:
            new_dict[6] += old_dict[num]
            new_dict[8] = old_dict[num]
        elif num == 0 and 6 in new_dict and 8 in new_dict:
            new_dict[6] += old_dict[num]
            new_dict[8] += old_dict[num]
        #else statement adds the values from the old dictionary down a level e.g. 123 1's -> 123 0's
        else:
            if num - 1 in new_dict:
                new_dict[num-1] += old_dict[num]
            else:
                new_dict[num-1] = old_dict[num]

    del old_dict # deletes the old dictionary from memory
    return new_dict #returns the newly created dictionary

def main():
    line = open_file() #opens and reads in line from file

    dict = fill_dictionary(line) #adds read in line to dictionary

    #loops for specified number of days
    for i in range(NUM_DAYS):
        dict = new_dictionary(dict) #creates a new dictionary every day based off the previous days dictionary
    print(dict) #prints the final dictionary
    print(sum(dict.values())) #prints the final population from the sum of all the dictionary entries

if __name__ =="__main__":
    main()
