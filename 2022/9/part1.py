# Test points visited
# (0,0)
# (1,0)
# (2,0)
# (3,0)
# (4,1)
# (4,2)
# (4,3)
# (3,4)
# (2,4)
# (3,3)
# (4,3)
# (3,2)
# (2,2)
# (1,2)
tail_location = {'0,0': True}

def add_to_dict(pos1, pos2):
    if str(pos1) + ',' + str(pos2) in tail_location:
        return
    else:
        tail_location[str(pos1) + ',' + str(pos2)] = True

def check_tail(head, tail):
    if head[0]-2 == tail[0]:
        tail[0]+=1
        if head[1] > tail[1]:
            tail[1] += 1
        elif head[1] < tail[1]:
            tail[1] -= 1
        add_to_dict(tail[0], tail[1])
    elif head[0]+2 == tail[0]:
        tail[0]-=1
        if head[1] > tail[1]:
            tail[1] += 1
        elif head[1] < tail[1]:
            tail[1] -= 1
        add_to_dict(tail[0], tail[1])
    
    if head[1]-2 == tail[1]:
        tail[1]+=1
        if head[0] > tail[0]:
            tail[0] += 1
        elif head[0] < tail[0]:
            tail[0] -= 1
        add_to_dict(tail[0], tail[1])
    elif head[1]+2 == tail[1]:
        tail[1]-=1
        if head[0] > tail[0]:
            tail[0] += 1
        elif head[0] < tail[0]:
            tail[0] -= 1
        add_to_dict(tail[0], tail[1])

with open('9/input.txt', 'r') as f:
    tail = [0,0]
    head = [0,0]
    
    lines = f.readlines()
    for line in lines:
        line = line.strip().split(" ")
        print(line[0],line[1])
        add_to_dict(tail[0], tail[1])
        
        if line[0] == 'R':
            for i in range(0, int(line[1])):
                head[0] += 1          
                check_tail(head, tail)
        elif line[0] == 'L':
            for i in range(0, int(line[1])):
                head[0] -= 1           
                check_tail(head, tail)
        elif line[0] == 'U':
            for i in range(0, int(line[1])):
                head[1] += 1              
                check_tail(head, tail)
        elif line[0] == 'D':
            for i in range(0, int(line[1])):
                head[1] -= 1              
                check_tail(head, tail)
        
    # print(tail_location)
    print(len(tail_location))