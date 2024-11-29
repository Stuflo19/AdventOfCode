import numpy as np

def check_direction(list, tree):
    total = 0
    for v in list:
        if v < tree: total += 1
        else: 
            total += 1
            break
    return total

with open("8/input.txt", "r") as f:
    trees = np.array([list(map(int, line.strip())) for line in (f.readlines())])
    center_trees = trees[1:len(trees)-1, 1:len(trees)-1]
    max = 0
    for row in range(0, len(center_trees)):
        for value in range(0, len(center_trees[row])):
            tree = center_trees[row][value]
            vertical = trees[:,value+1]
            
            left = check_direction(reversed(trees[row+1][0:value+1]), tree)
            right = check_direction(trees[row+1][value+2:len(trees[row])], tree)
            up = check_direction(reversed(vertical[0:row+1]), tree)
            down = check_direction(vertical[row+2:len(vertical)], tree)
            
            # print(left * right * up * down)
            total = left * right * up * down
            if max == 0:
                max = total
            elif total > max:
                max = total
    print(max)