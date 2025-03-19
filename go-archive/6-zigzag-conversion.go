func convert(s string, numRows int) string {
    
    if numRows == 1 {
        // the solution relies on having some 
        // intermediate levels
        return s
    }

    /**
    practically a bfs problem.
    for each row find the peers on the same level
    traverse
    **/

    ans := []byte{}
    gap := 2 * numRows - 2
    // source nodes
    for src := 0; src < len(s); src += gap {
        ans = append(ans, s[src])
    }

    // for intermediate nodes, 
    // finding the nodes are at level and level + gap
    // but then there is another set of nodes 
    // you have to reach before you reach peer at distance gap
    for level := 1; level < numRows - 1; level++ {
        for node := level; node < len(s); node += gap {
            ans = append(ans, s[node])
            if peer := node + 2 * (numRows - level - 1); peer < len(s) {
                ans = append(ans, s[peer])
            }
        }
    }
    
    // destination nodes
    for dest := numRows - 1; dest < len(s); dest += gap {
        ans = append(ans, s[dest])
    }

    return string(ans);
  
  /* a few years ago, I also had a clever python submission.
  check my own leetcode submissions for answer */

}
