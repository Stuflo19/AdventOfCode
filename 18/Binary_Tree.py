import Node

class Binary_Tree:
    def create_tree(self, value):
        if type(value) == int:
            return Node(value, None, None)
        
        left_node = self.create_tree(value[0])
        right_node = self.create_tree(value[1])

        parent = Node(None, left_node, right_node)
        left_node.parent = parent
        right_node.parent = parent
    
    def add_snail_number(value):
        l = value[0]
        r = value[1]
