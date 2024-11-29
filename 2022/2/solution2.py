score = 0
filename = "2/input.txt"

draw = 0
win = 0
lose = 0

# A = Rock
# B = Paper
# C = Scissors

# X = Lose
# Y = Draw
# Z = Win

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

        # print(score)
        
        if(line[0] == "A"): 
            draw = 4
            win = 8
            lose = 3
        if(line[0] == "B"): 
            draw = 5
            win = 9
            lose = 1
        if(line[0] == "C"): 
            draw = 6
            win = 7
            lose = 2

        if line[1] == "X":
            # lose
            score += lose
        if line[1] == "Y":
            # Draw
            score += draw
        if line[1] == "Z":
            # Win
            score += win
    
print(score)
