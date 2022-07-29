

def arrayRankTransform(arr):
    ranks = {v:i for i,v in enumerate(sorted(set(arr)),1 )}
    return [ranks[v] for v in arr]