score = 0
filename = "2/input.txt"

you = ""
opponent = ""

# A = Rock
# B = Paper
# C = Scissors

# X = Rock
# Y = Paper
# Z = Scissors

# A beat by Y
# B beat by Z
# C beat by X

# rock = 1 point
# paper = 2 point
# scissors = 3 point

with open(filename, 'r') as f:
    lines = f.readlines()
    for line in lines:
        line = line.strip().split(" ")

        print(score)
        
        if(line[0] == "A"): opponent = "rock"
        if(line[0] == "B"): opponent = "paper"
        if(line[0] == "C"): opponent = "scissors"

        if( line[1] == "X"): 
            score += 1
            you = "rock"
        elif( line[1] == "Y"): 
            score += 2
            you = "paper"
        elif( line[1] == "Z"): 
            score += 3
            you = "scissors"

        if(you == opponent):
            score += 3
        
        if(line[0] == "A" and line[1] == "Y"):
            score += 6
        elif(line[0] == "A"):
            continue
        
        if(line[0] == "B" and line[1] == "Z"):
            score += 6
        elif(line[0] == "B"):
            continue
        
        if(line[0] == "C" and line[1] == "X"):
            score += 6
        elif(line[0] == "C"):
            continue
    
print(score)
