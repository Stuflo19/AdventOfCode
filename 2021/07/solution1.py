file_name = "07/input.txt"

#function to open file and read in line as an array of nums
def open_file():
    with open(file_name, 'r') as input:
        lines = input.readlines()

        for line in lines:
            line = line.strip().split(",")

            line = list(map(int, line)) #turns the string array into an int array

            return line.copy()

#function to find the least amount of fuel used
def least_fuel(input, point):
    fuel_used = 0

    #for every number in the input
    for num in input:
        while num < point:
            fuel_used += 1 #increase fuel used by the cost of the move
            num += 1 
        while num > point:
            fuel_used += 1 #increas fuel used by the cost of the move
            num -= 1

    return fuel_used #return the number of fuel this move used

#function to find the maximum number from the input
def max_num(input):
    max = 0

    for num in input:
        if num > max:
            max = num
    
    return num


def main():
    fuel = 0 #holds the fuel used by the current iteration
    min = 0 #holds the minimum fuel used

    input = open_file() #sets the input equal to the int array

    nums = max_num(input) #sets the number that we have to range between

    #loops for all numbers from 0 - max number in the array
    for i in range(nums):
        fuel = least_fuel(input, i) #finds current iterations fuel used
        
        #if it is the first iteration set min to fuel
        if i == 0:
            min = fuel

        #set min fuel
        if fuel < min:
            min = fuel
    
    print(min)

if __name__ == "__main__":
    main()
            



    