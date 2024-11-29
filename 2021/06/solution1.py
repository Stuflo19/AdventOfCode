file_name = "06\input.txt"
counter = 0
NUM_DAYS = 80

#opens file
with open(file_name, 'r') as input:
    lines = input.readlines() #reads lines from file

    for line in lines:
        line = line.strip().split(',') #strips \n from end of line and splits the line by ','

    line = list(map(int, line)) #turns the string array into an int array

    #loops while loop counter is < NUM_DAYS
    while(counter < NUM_DAYS):
        #loops for the size of the line list
        for num in range(len(line)):
            #if the value at the position of line is > 0 then subtract 1 from it
            if line[num] > 0:
                line[num] -= 1
            #else if the value is 0 then add a new 6 and 8
            elif line[num] == 0:
                line[num] = 6
                line.append(8)
        
        counter += 1

print(len(line)) #print final population
