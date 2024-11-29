class Cave:
    caves = {} #holds all the instances of Caves

    def __init__(self, name, path):
        self.name = name #holds the name of the cave
        self.paths = [path] #holds a list of the paths the cave has
        Cave.caves[name] = self #adds the cave to the cave dictionary

    #adds a new path to a cave
    def add_path(self, path):
        self.paths.append(path)

#opens the file and reads in the lines from it
def open_file():
    file_name = "12/input.txt"

    with open(file_name, 'r') as input:
        lines = [list(line.strip().split('-')) for line in (input.readlines())]
        return lines

#function used to read in the lines from the file and create Cave objects for each cave and its paths
def create_caves(lines):
    for line in lines:
        if line[0] not in Cave.caves:
            Cave(line[0], line[1])
        else:
            Cave.caves[line[0]].add_path(line[1])

        if line[1] not in Cave.caves:
            Cave(line[1], line[0])
        else:
            Cave.caves[line[1]].add_path(line[0])

#function based on https://www.reddit.com/user/Jools_Jops/ solution
#Figure out how to do this with true/false and also maybe try a recursive method
def navigate_caves():
    total = 0 #tracks total number of times end was reached
    path = [0, 'start'] #tracks the path we are currently on
    paths = [path] #holds all the paths that still need to be checked
    current_cave = Cave.caves['start'] #holds the current cave we are in

    #loops while the paths to run through are not empty
    while len(paths) != 0:
        #if the current cave is small and appears more than once set duplicate found        
        if current_cave.name.islower() and path.count(current_cave.name) > 1:
            path[0] = 1                  
        #loops for all the possible caves we can step to from the current cave     
        for cave in current_cave.paths:
            #if the next cave is start then skip it
            if cave == 'start':
                continue  
            #if the next cave is end then increment number of found paths
            if cave == 'end':
                total += 1
                continue
            #if the next cave is small and is already in path and appears twice already then skip
            if cave.islower() and cave in path and path[0] == 1:
                continue
            
            #adds the path to the paths to be continued
            paths.append(path + [cave])

        path = paths.pop() #the current path we are expanding is the last possible path from paths
        current_cave = Cave.caves[path[-1]] #Sets the next cave to be the last item in the path to be searched

    return total #returns number of caves

def main():
    total = 0

    lines = open_file()

    create_caves(lines)

    total = navigate_caves()

    print(total)

if __name__ == "__main__":
    main()