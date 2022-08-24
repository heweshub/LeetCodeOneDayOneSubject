
class Solution:
    def amountOfTime(self, root: Optional[TreeNode], start: int) -> int:
        parent = {}
        def dfs(node, pa):
            if node is None: return
            if node.val == start: self.start = node
            parent[node] = pa
            dfs(node.left, node)
            dfs(node.right, node)
        dfs(root, None)

        ans = -1
        vis = {None, self.start}
        q = [self.start]
        while q:
            ans += 1
            tmp = q
            q = []
            for node in tmp:
                for x in node.left,node.right,parent[node]:
                    if x not in vis:
                        vis.add(x)
                        q.append(x)
        return ans