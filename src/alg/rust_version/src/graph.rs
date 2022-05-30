#![allow(dead_code)]
pub struct Solution {}
impl Solution {
    fn dfs(graph: &[Vec<i32>], node: i32, target: i32, path: &mut Vec<i32>, mut paths: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        path.push(node);

        if node == target {
            // We have reached the target. Append path to output
            paths.push(path.clone());
        } else {
            // Branch out to all children
            for child in &graph[node as usize] {
                paths = Self::dfs(graph, *child, target, path, paths);
            }
        }
        
        // backtrack
        path.pop();
        paths
    }
    pub fn all_paths_source_target(graph: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        Self::dfs(&graph, 0, (graph.len() - 1) as i32, &mut vec![], vec![])
    }
}