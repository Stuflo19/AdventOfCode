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

    #4 points to search, i
    for row in range( len(points) ):
        for col in range( len(points[row]) ):

            current_val = int(points[row][col])

            if row != 0 and current_val >= int(points[row-1][col]):
                continue

            if row != ( len(points) - 1) and current_val >= int(points[row+1][col]):
                continue

            if col != 0 and current_val >= int(points[row][col-1]):
                continue

            if col != ( len(points[row]) - 1 ) and current_val >= int(points[row][col+1]):
                continue

            if current_val+1 in lows:
                lows[current_val+1] += (current_val + 1)
            else:
                lows[current_val+1] = (current_val + 1)
    
    print(lows)
    return lows.copy()

def main():
    lines = open_file()

    points = create_grid(lines)

    lows = check_grid(points)

    print(sum( lows.values() ))


if __name__ == "__main__":
    main()
