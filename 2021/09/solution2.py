import numpy as np

def open_file():
    file_name = "09/input.txt"

    with open(file_name, 'r') as input:
        lines = input.readlines()

        return lines.copy()


def create_grid(lines):
    points = None

    for line in lines:
        line = line.strip()

        if not points:
            points = [ list(line) ]
        else:
            points.append(list(line))
    
    return points.copy()


def check_grid(points):
    lows = {}
    basin_size = []

    #4 points to search, i
    for row in range( len(points) ):
        for col in range( len(points[row]) ):
            total = 0
            checked = np.zeros((len(points), len(points[row])))
            current_val = int(points[row][col])

            if row != 0 and current_val > int(points[row-1][col]):
                continue

            if row != ( len(points) - 1) and current_val > int(points[row+1][col]):
                continue

            if col != 0 and current_val > int(points[row][col-1]):
                continue

            if col != ( len(points[row]) - 1 ) and current_val > int(points[row][col+1]):
                continue

            if current_val+1 in lows:
                lows[current_val+1] += (current_val + 1)
            else:
                lows[current_val+1] = (current_val + 1)
            
            total = search_basin(points, row, col, checked)
            basin_size.append(total)

    return basin_size

def search_basin(points,row,col,checked):
    total = 1
    checked[row][col] = 1

    #take in value with 4 smaller around it, start at value on the left and check
    if row != 0 and checked[row-1][col] == 0 and int(points[row-1][col]) != 9:
        total += search_basin(points, row-1, col, checked)

    if row != ( len(points) - 1) and checked[row+1][col] == 0 and int(points[row+1][col]) != 9:
        total += search_basin(points, row+1, col, checked)

    if col != 0 and checked[row][col-1] == 0 and int(points[row][col-1]) != 9:
        total += search_basin(points, row, col-1, checked)

    if col != ( len(points[row]) - 1 ) and checked[row][col+1] == 0 and int(points[row][col+1]) != 9:
        total += search_basin(points, row, col+1, checked)
    
    return total

def find_max_basins(basins):
    vals = []
        
    for i in range(3):
        vals.append(np.amax(basins))

        basins.remove(np.amax(basins))
    
    return vals

def main():
    lines = open_file()

    points = create_grid(lines)

    basins = check_grid(points)

    vals = find_max_basins(basins)

    result = vals[0] * vals[1] * vals[2]

    print(result)


if __name__ == "__main__":
    main()
