import Binary_Tree as bt

#Function that reads in lines from the file as a list and strips them of the \n char
def open_file(file_name):
    with open(file_name) as input:
        lines = input.readlines()
        print(lines)
    return lines

def split(curr_sum):
    new_sum = []

def explode_left(val):
    root = tree.root
    to_find = tree.root

    while root.left is None or root.left == to_find:
        to_find = root
        root = root.parent

        if root is None: return

    root = root.l

    while root.r is not None:
        root = root.r
    
    root.val += val

def explode_right(val):
    root = tree.root
    to_find = root.copy()


    while root.right is None or root.right == to_find:
        to_find = root
        root = root.parent

        if root is None: return

    root = root.r

    while root.l is not None:
        root = root.l

    root.val += 1

def explode(node, depth):
    new = None
    
    if depth == 4 and node.val is None:

def find_sum(lines):
    global tree
    tree = None

    #for every line of the file add it to the binary tree
    for line in lines:
        snail_num = eval(line) #inputs the line as an array
        print(snail_num)

        #checks if any tree exists
        if tree is None:
            tree = bt.Binary_Tree()
            tree.create_tree(snail_num) #creates the brand new tree
            continue
    
        new_num = bt.Binary_Tree()
        new_num.create_tree(snail_num) #creates new snail number tree
        tree.add_new_tree(tree.root, new_num.root) #adds new snail number tree the the right and old tree to the left

        del new_num #clear the memory allocated to new_num     

def main():
    lines = open_file('18/input.txt')

    find_sum(lines)

if __name__ == '__main__':
    main()