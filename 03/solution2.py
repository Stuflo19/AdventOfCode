file_name = '03\input.txt'
oxygen = set()
co2 = set()
Olinenums = []
Clinenums = []
co2rem = set()
oxyrem = set()
oxyFound = False
co2Found = False

with open(file_name, 'r') as input:
    lines = input.readlines()

    for i in range(12):
        Onum1 = 0
        Onum0 = 0
        Cnum1 = 0
        Cnum0 = 0
        Olinenums.clear()
        Clinenums.clear()
        oxyrem.clear()
        co2rem.clear()


        if i == 0:
            for line in lines:
                Olinenums.append(int(line[i]))
            
            for num in Olinenums:
                if num == 0:
                    Onum0 += 1
                else:
                    Onum1 +=1

            if(Onum1 > Onum0 or Onum1 == Onum0):
                for line in lines:
                    if line[i] == '1':
                        oxygen.add(line)
                    else:
                        co2.add(line)
            else:
                for line in lines:
                    if line[i] == '0':
                        oxygen.add(line)
                    else:
                        co2.add(line)

        else:
            for line in oxygen:
                Olinenums.append(int(line[i]))
            for line in co2:
                Clinenums.append(int(line[i]))

            for num in Olinenums:
                if num == 0:
                    Onum0 += 1
                else:
                    Onum1 +=1
            
            for num in Clinenums:
                if num == 0:
                    Cnum0 += 1
                else:
                    Cnum1 +=1

            if not oxyFound:
                if Onum1 > Onum0 or Onum1 == Onum0:
                    for search in oxygen:
                        if search[i] == '0':
                            oxyrem.add(search)
                else:
                    for search in oxygen:
                        if search[i] == '1':
                            oxyrem.add(search)

                oxygen -= oxyrem

            if not co2Found:
                if Cnum1 > Cnum0 or Cnum1 == Cnum0:
                    for search in co2:
                        if search[i] == '1':
                            co2rem.add(search)
                else:
                    for search in co2:
                        if search[i] == '0':
                            co2rem.add(search)

                co2 -= co2rem

            if len(oxygen) == 1:
                oxyFound = True
            if len(co2) == 1:
                co2Found = True

        if co2Found and oxyFound:
                break

print("co2: ")
print(co2)
print("\noxygen: ")
print(oxygen) 

# code on these 2 lines is from: https://stackoverflow.com/questions/17528374/python-convert-set-to-string-and-vice-versa
co2 = ', '.join(list(map(str, co2)))
oxygen = ', '.join(list(map(str, oxygen)))

co2 = int(co2, 2)
oxygen = int(oxygen, 2)

print(co2 * oxygen)