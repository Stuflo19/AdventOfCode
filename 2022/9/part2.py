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

def add_to_dict(tail_location, pos1, pos2):
    if str(pos1) + ',' + str(pos2) in tail_location:
        return
    else:
        tail_location[str(pos1) + ',' + str(pos2)] = True

def check_tail(head, tail, i, num_knots, tail_location):
    if head[0]-2 == tail[0]:
        tail[0]+=1
        if head[1] > tail[1]:
            tail[1] += 1
        elif head[1] < tail[1]:
            tail[1] -= 1
    elif head[0]+2 == tail[0]:
        tail[0]-=1
        if head[1] > tail[1]:
            tail[1] += 1
        elif head[1] < tail[1]:
            tail[1] -= 1
    
    if head[1]-2 == tail[1]:
        tail[1]+=1
        if head[0] > tail[0]:
            tail[0] += 1
        elif head[0] < tail[0]:
            tail[0] -= 1
    elif head[1]+2 == tail[1]:
        tail[1]-=1
        if head[0] > tail[0]:
            tail[0] += 1
        elif head[0] < tail[0]:
            tail[0] -= 1

    if i == num_knots-1:
        add_to_dict(tail_location, tail[0],tail[1])

def solution(part_no, num_knots):
    with open('9/input.txt', 'r') as f:
        tails = [[0,0] for _ in range(num_knots)]
        tail_location = {'0,0': True}
        
        lines = f.readlines()
        for line in lines:
            line = line.strip().split(" ")
            
            if line[0] == 'R':
                for i in range(0, int(line[1])):
                    tails[0][0] += 1          
                    for i in range(1,num_knots): 
                        check_tail(tails[i-1], tails[i], i, num_knots, tail_location)
            elif line[0] == 'L':
                for i in range(0, int(line[1])):
                    tails[0][0] -= 1           
                    for i in range(1,num_knots): 
                        check_tail(tails[i-1], tails[i], i, num_knots, tail_location)
            elif line[0] == 'U':
                for i in range(0, int(line[1])):
                    tails[0][1] += 1
                    for i in range(1,num_knots): 
                        check_tail(tails[i-1], tails[i], i, num_knots, tail_location)
            elif line[0] == 'D':
                for i in range(0, int(line[1])):
                    tails[0][1] -= 1              
                    for i in range(1,num_knots): 
                        check_tail(tails[i-1], tails[i], i, num_knots, tail_location)
            
            # add_to_dict(tails[num_knots-1][0], tails[num_knots-1][1])
            
        # print(tail_location)
        print(part_no, len(tail_location))

if __name__ == '__main__':
    solution("Part 1: ", 2)
    solution("Part 2: ", 10)