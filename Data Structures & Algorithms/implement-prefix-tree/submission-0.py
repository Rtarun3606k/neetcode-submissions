class TriNode:
    def __init__(self):
        self.children = {}
        self.endOfWord = False



class PrefixTree:

    def __init__(self):
        self.root = TriNode()
        
        

    def insert(self, word: str) -> None:
        current = self.root
        for i in word:
            if i not in current.children:
                current.children[i] = TriNode()
            current = current.children[i]
        current.endOfWord = True

    def search(self, word: str) -> bool:
        current = self.root
        for i in word:
            if i not in current.children:
                return False
            current = current.children[i]
        return current.endOfWord

    def startsWith(self, prefix: str) -> bool:
        current = self.root
        for i in prefix:
            if i not in current.children:
                return False
            else:
                return True
        
        