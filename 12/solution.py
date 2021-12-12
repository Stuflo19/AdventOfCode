class Cave:
    caves = {}
    visited = {}

    def __init__(self, name, path):
        self.name = name
        self.paths = [path]
        Cave.caves[name] = self
        if name.islower():
            self.small = True
        else:
            self.small = False


    def add_path(self, path):
        self.paths.append(path)

    def get_paths(self):
        return self.paths

def open_file():
    file_name = "12/input.txt"

    with open(file_name, 'r') as input:
        lines = [list(line.strip().split('-')) for line in (input.readlines())]
        return lines

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

def navigate_caves(paths):
    total = 0
    path = []
    current_cave = Cave.caves['start']
    while len(path) != 0:
        for i in range(len(current_cave.paths)):
            print("visited " + current_cave.paths[i])

            if current_cave.paths[i] == 'end':
                total += 1
                break
            elif current_cave.paths[i] == 'start':
                continue  
            elif current_cave.small and current_cave.name in path:
                break
            else:
                print(current_cave.paths[i])
                path.append(current_cave.paths[i])

        paths = path.pop
        print(paths)

    return total

def aoc12(filename):
    graph = {
        "start": [],
        "end": []}
    
    with open(filename, 'r') as f:
        for line in f.readlines():
            a, b = line.strip().split('-')
            if a not in graph.keys():
                graph[a] = []
            graph[a].append(b)
            if b not in graph:
                graph[b] = []
            graph[b].append(a)

    node = 'start'
    path = [node] # path[0] added to track one revisit to small cave
    count = 0
    paths = [path]
    print("Graph node: " + str(graph[node]))
    while len(paths) != 0:
        for n in graph[node]:
            if n == 'start':
                continue
            if n == 'end':
                count += 1
                continue
            if n in path and n.islower(): # 'and path[0] == 1' added for day 12B 
                continue
            paths.append(path + [n])
        path = paths.pop()
        print(path)
        node = path[-1]
    
    print(count)

def main():
    total = 0
    path = []

    aoc12("12/input.txt")

    lines = open_file()

    create_caves(lines)

    # for cave in Cave.caves:
    #     print("Name: " + Cave.caves[cave].name)
    #     print(Cave.caves[cave].paths)

    total = navigate_caves(path)

    print(total)


if __name__ == "__main__":
    main()