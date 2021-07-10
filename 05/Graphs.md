### About graphs and nodes
A graph `G(V,E)` is a finete, nonempty set of vertices `V` (or nodes) and a set of edges `E`. There are two main types of graphs: `cyclic graphs` and `acyclic graphs`. A cyrclic graph is a graph where all or a number of its vertices are connected in a closed chain. In acyclic graphs, there are not any closed chains. A `directed graph` is a graph with edges that have a direction assosiated with them, while `directed acyclic graph` is directed graph with nocycles in it.

#### TIP
As a node may contain any kind of information, nodes are usually implemented using Go structures due to their versatility.