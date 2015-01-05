# This is a simple POC using scipy (faster than networkx)

import numpy as np
from scipy.sparse import csgraph
from scipy.sparse import csr_matrix

G_dense = np.array([[0, 3, 1, 1],
                    [3, 0, 0, 1],
                    [1, 0, 0, 0],
                    [1, 1, 0, 0]])

G_sparse = csr_matrix(G_dense)
distances, predecessors = csgraph.shortest_path(G_sparse, return_predecessors=True)

path = []
i = 1
while i != 0:
    path.append(i)
    i = predecessors[0, i]
path.append(0)

i1 = 0;
i2 = 1;

print "Shortest path from {} to {} = {}".format(i1, i2, path[::-1])
print "This path takes {} steps and costs {}".format(len(path), distances[i1, i2])
