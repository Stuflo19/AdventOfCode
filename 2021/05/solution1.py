import numpy as np

file_name = "05/input.txt"
x1 = 0
y1 = 0
x2 = 0
y2 = 0
temp = 0
visitedpoints = np.zeros((1000, 1000))
crossed = 0

with open(file_name, 'r') as input:
    lines = input.readlines()
    
    for line in lines:
        line = line.strip().split(" -> ")
        temp = line[0].split(',')
        x1 = int(temp[0])
        y1 = int(temp[1])

        temp = line[1].split(',')
        x2 = int(temp[0])
        y2 = int(temp[1])

        if(y1 == y2 and x1 < x2):
            visitedpoints[x2][y2] += 1
            for x in range(x1,x2):
                visitedpoints[x][y1] += 1
        elif(y1 == y2 and x1 > x2):
            visitedpoints[x1][y1] += 1
            for x in range(x2,x1):
                visitedpoints[x][y1] += 1
        elif(x1 == x2 and y1 < y2):
            visitedpoints[x2][y2] += 1
            for y in range(y1,y2):
                visitedpoints[x1][y] += 1
        elif(x1 == x2 and y1 > y2):
            visitedpoints[x1][y1] += 1
            for y in range(y2,y1):
                visitedpoints[x1][y] += 1
        
    for row in range(len(visitedpoints)):
        for col in range(len(visitedpoints[row])):
            #print(visitedpoints[row][col], end = " ")
            if visitedpoints[row][col] > 1:
                crossed += 1
        #print()
    
    print(crossed)




        





