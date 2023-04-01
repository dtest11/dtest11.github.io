```go


// DFS algorithm in Java
  // Graph creation
  Graph(int vertices) {
    adjLists = new LinkedList[vertices];
    visited = new boolean[vertices];
    for (int i = 0; i < vertices; i++)
      adjLists[i] = new LinkedList<Integer>();
  }


  // Add edges
  void addEdge(int src, int dest) {
    adjLists[src].add(dest);
  }


  // DFS algorithm
  void DFS(int vertex) {
    visited[vertex] = true;
    System.out.print(vertex + " ");


    Iterator<Integer> ite = adjLists[vertex].listIterator();
    while (ite.hasNext()) {
      int adj = ite.next();
      if (!visited[adj])
        DFS(adj);
    }
  }
}
```