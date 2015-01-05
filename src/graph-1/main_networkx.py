import sys
import os
import networkx as nx
from optparse import OptionParser

# parse command line arguments
usage = "usage: %prog [options] Source Target"
parser = OptionParser(usage)
parser.add_option("--all",
              action="store_true", dest="print_all",
              help="Print the set of minimal distance path from S to T")

parser.add_option("-f", "--filename",
                  default="graph.csv", dest="filename",
                  metavar="FILE", help="Read the graph from FILE in csv format")

(opts, args) = parser.parse_args()
if len(args) != 2:
    parser.print_help()
    sys.exit(1)

if os.path.exists(opts.filename) == False:
    print("File {} does not exists".format(opts.filename))
    sys.exit(1)

try:
    source = int(args[0])
    target = int(args[1])
except ValueError:
    parser.error("Source and Target must be of type int")

G = nx.read_edgelist(opts.filename, nodetype=int, delimiter=",", create_using=nx.DiGraph(), data=(('weight',float),))

if G.has_node(source) == False or G.has_node(target) == False:
    parser.error("Source and Target must be nodes inside the Graph")

if nx.has_path(G, source, target) ==  False:
    print("No path found from {} to {}".format(source, target))
    sys.exit(1)

path = nx.shortest_path(G, source = source, target = target, weight = "weight")
lenght = nx.shortest_path_length(G, source = source, target = target, weight = "weight")

print "Shortest path from {} to {} = {}".format(source, target, path)
print "This path takes \033[92m{}\033[0m steps and costs \033[92m{}\033[0m".format(len(path), lenght)


if opts.print_all:
    print("The list of all path with weight {} :".format(lenght))
    print([p for p in nx.all_shortest_paths(G, source = source, target = target, weight = "weight")])
