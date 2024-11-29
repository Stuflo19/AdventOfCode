import numpy as np 

def CheckForLess(list, val, i, row):
    vertical = trees[: ,i]
    if i > 0:    
        if max(list[i+1:len(list)])<val or max(list[0:i])<val or max(vertical[row+1:len(vertical)])<val or max(vertical[0:row])<val:
            return True

with open("8/input.txt") as f:
    trees = np.array([list(map(int, line.strip())) for line in (f.readlines())]) 
    total = (len(trees)-2 + len(trees[:,1])) * 2
    center_trees = trees[1:len(trees)-1, 1:len(trees)-1]
    for row in range(0, len(center_trees)): 
        for i in range(0, len(center_trees[row])):
            if CheckForLess(trees[row+1], center_trees[row][i], i+1, row+1):
                total+=1
    
    print(total)