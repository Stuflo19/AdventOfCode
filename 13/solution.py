import numpy as np
np.set_printoptions(threshold=np.inf)

def open_file(file_name):
 with open(file_name, 'r') as input:
    lines = np.array([list(line.strip().split(',')) for line in (input.readlines())])

    return lines

def find_max(lines):
    max = []
    maxX = 0
    maxY = 0

    for line in lines:
        if line[0] == '':
            break
        if int(line[0]) > maxX:
            maxX = int(line[0]) + 1
        if int(line[1]) > maxY:
            maxY = int(line[1]) + 1
    
    print(maxX)
    print(maxY)
    max.append(maxX)
    max.append(maxY)

    return max

def create_paper(lines):
    max = find_max(lines)
    paper = np.zeros((max[1], max[0]))
    points_mapped = False

    for line in lines:
        if not points_mapped:
            if line[0] == '':
                points_mapped = True
                continue
            else:
                x = int(line[0])
                y = int(line[1])
                paper[y, x] = 1
        else:
            line = line[0].split(" ")
            line = line[2].split('=')
            move = int(line[1])

            if(line[0] == 'x'):
                for y in range(len(paper)):
                    for x in range(move,len(paper[y])):
                        if move == x:
                            paper[y][x] = 0
                        if paper[y][x] == 1:
                            paper[y][x] = 0
                            new_x = len(paper[0])-1 - x
                            paper[y, new_x] = 1

                paper = paper[0:len(paper[0]+30),0:move]

            elif line[0] == 'y':
                for y in range(move, len(paper)):
                    for x in range(len(paper[y])):
                        if move == y:
                            paper[y][x] = 0
                        if paper[y][x] == 1:
                            paper[y][x] = 0
                            new_y = 0 + (len(paper)-1 - y)
                            paper[new_y, x] = 1

                paper = paper[:move,:len(paper)+30] #Find out why I need +30 here
            print("Amount: " + str(np.count_nonzero(paper)))      
            
    return paper

def main():
    lines = open_file("13/input.txt")

    paper = create_paper(lines)

    # Print them out
    for y in range(0, len(paper)):
        line = ''
        for x in range(0, len(paper[y])):
            if paper[y][x] == 1:
                line += '#'
            else:
                line += ' '
        print(line)

if __name__ == "__main__":
    main()
