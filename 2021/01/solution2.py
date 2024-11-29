file_name = '01\depths.txt'
prevline = 9999
count = 0
num_increases = 0
to_sum = []
sum = 0
prevsum = 0

with open(file_name, 'r') as depths:
    lines = depths.readlines()

    for line in lines:
        sum = 0
        to_sum.insert(0, (int(line)))

        if count > 1:
            for i in range(3):
                sum += to_sum[i]
                print(sum)

            if sum > prevsum and prevsum != 0:
                num_increases += 1
            
            prevsum = sum
        
        count += 1

print(num_increases)