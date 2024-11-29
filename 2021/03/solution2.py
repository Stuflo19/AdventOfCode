file_name = '03\input.txt' #stores the file name to be read from
oxygen = set() #stores a set of all the oxygen values to be refined
co2 = set() #stores a set of all the co2 values to be refined
Olinenums = [] #stores all of the current columns numbers in oxygen
Clinenums = [] #stores all of the current columns number in co2
co2rem = set() #set to store all the values that have been found that are to be removed from the main set co2
oxyrem = set() #set to store all the values that have been found that are to be removed from the main set oxygen
oxyFound = False #stores whether the final value for oxygen has been found
co2Found = False #stores whether the final value for co2 has been found

with open(file_name, 'r') as input:
    lines = input.readlines()

    #loop for all 12 numbers of each line
    for i in range(12):
        Onum1 = 0
        Onum0 = 0
        Cnum1 = 0
        Cnum0 = 0
        Olinenums.clear()
        Clinenums.clear()
        oxyrem.clear()
        co2rem.clear()

        #if it is the first loop through then we need to populate the arrays for the first time
        if i == 0:
            #store all the values in the current column of numbers
            for line in lines:
                Olinenums.append(int(line[i]))
            
            #for each row in the column count how many times each number appears to find popular number
            for num in Olinenums:
                if num == 0:
                    Onum0 += 1
                else:
                    Onum1 +=1

            #if 1 is most popular or both are equally popular
            if(Onum1 > Onum0 or Onum1 == Onum0):
                #add every line beginning with 1 to oxygen and every line beginning with 0 to co2
                for line in lines:
                    if line[i] == '1':
                        oxygen.add(line)
                    else:
                        co2.add(line)
            #else 0 is the most popular
            else:
                #add every line beggining with 0 to oxygen and every line beginning with 1 to co2
                for line in lines:
                    if line[i] == '0':
                        oxygen.add(line)
                    else:
                        co2.add(line)

        #else this is not the first loop and we will have to be removing instead of adding
        else:
            #get the numbers from the current column from both oxygen and co2
            for line in oxygen:
                Olinenums.append(int(line[i]))
            for line in co2:
                Clinenums.append(int(line[i]))

            #find which number is more common in the current oxygen column
            for num in Olinenums:
                if num == 0:
                    Onum0 += 1
                else:
                    Onum1 +=1
            
            #find which number is more common in the current co2 column
            for num in Clinenums:
                if num == 0:
                    Cnum0 += 1
                else:
                    Cnum1 +=1

            #if there is still more than 1 value inside of oxygen
            if not oxyFound:
                
                #if 1 is most popular or 1 and 0 are equally popular
                if Onum1 > Onum0 or Onum1 == Onum0:
                    #find any rows with 0 in the current column and add it to the list to be removed from oxygen
                    for search in oxygen:
                        if search[i] == '0':
                            oxyrem.add(search)
                #else 0 us the most popular
                else:
                     #find any rows with 1 in the current column and add it to the list to be removed from oxygen
                    for search in oxygen:
                        if search[i] == '1':
                            oxyrem.add(search)

                oxygen -= oxyrem #remove the values found from the main set

            #if there is still more than 1 value inside co2
            if not co2Found:
                #if the most common value is 1 or 1 and 0 are equally common
                if Cnum1 > Cnum0 or Cnum1 == Cnum0:
                    #find any rows with 1 in the current column and add it to the list to be removed from co2
                    for search in co2:
                        if search[i] == '1':
                            co2rem.add(search)
                #else 0 is the most common
                else:
                    #find any rows with 0 in the current column and add it to the list to be removed from co2
                    for search in co2:
                        if search[i] == '0':
                            co2rem.add(search)

                co2 -= co2rem #remove the values found from the main set

            #if statements to check if there is only 1 value left within the set and the final value has been found
            if len(oxygen) == 1:
                oxyFound = True
            if len(co2) == 1:
                co2Found = True

        #if both final values have been found it exits the loop and stops attempting to find them
        if co2Found and oxyFound:
                break

#prints the final values found
print("co2: ")
print(co2)
print("\noxygen: ")
print(oxygen) 

#converts the values of the sets to a string
# code on these 2 lines is from: https://stackoverflow.com/questions/17528374/python-convert-set-to-string-and-vice-versa
co2 = ', '.join(list(map(str, co2)))
oxygen = ', '.join(list(map(str, oxygen)))

#converts the values to integers from a string of binary digits
co2 = int(co2, 2)
oxygen = int(oxygen, 2)

#prints the life support rating
print(co2 * oxygen)