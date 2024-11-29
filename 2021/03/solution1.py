file_name = '03\input.txt'
num1 = 0 #stores how many times the number 1 appeared in the current column
num0 = 0 #stores how many times the number 0 appeared in the current column
epsilon = "" #stores the string value of all the least popular numbers from each column
gamma = "" #stores the string value of all the most popular numbers from each column
linenums = [] #stores all the values of each of the rows in the current column

#opens the file
with open(file_name, 'r') as input:
    lines = input.readlines() #stores all the lines in an array

    #loop for each of the columns
    for i in range(12):
        num1 = 0 #resets num1 back to 0 for each iteration
        num0 = 0 #resets num0 back to 0 for each iteration
        linenums.clear() #clears the linenums array for each iteration

        #loops for each line from the file
        for line in lines:
            linenums.append(int(line[i])) #adds the numbers of each column to an array
            
        #loops for each of the values in the current column and checks if it is a 1 or 0
        for num in linenums:
            if num == 0:
                num0 += 1
            else:
                num1 +=1

        #if 1 is more popular
        if num1 > num0:
            epsilon += "0"
            gamma += "1"

        #else 0 is more popular
        else:
            epsilon += "1"
            gamma += "0"

#print the values created
print("Gamma: ")
print(gamma)
print("\nEpsilon: ")
print(epsilon) 

#convert the binary to integers
gamma = int(gamma, 2)
epsilon = int(epsilon, 2)

#print the power consumption
print(gamma * epsilon)