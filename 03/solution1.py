file_name = '03\input.txt'
num1 = 0
num0 = 0
epsilon = ""
gamma = ""
linenums = []
savedlines = []


with open(file_name, 'r') as input:
    lines = input.readlines()

    for i in range(12):
        num1 = 0
        num0 = 0
        linenums.clear()

        for line in lines:
            linenums.append(int(line[i]))
            
        for num in linenums:
            if num == 0:
                num0 += 1
            else:
                num1 +=1

        if num1 > num0:
            epsilon += "0"
            gamma += "1"

        else:
            epsilon += "1"
            gamma += "0"

print("Gamma: ")
print(gamma)
print("\nEpsilon: ")
print(epsilon) 

gamma = int(gamma, 2)
epsilon = int(epsilon, 2)

print(gamma * epsilon)