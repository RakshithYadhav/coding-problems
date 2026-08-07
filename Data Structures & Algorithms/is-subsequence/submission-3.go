func isSubsequence(s string, t string) bool {
//1. Understand    → Restate in one sentence. Make your own example.
    // check if s is a subsequence of t
    // s is a new string formed from deleting some characters without disturbing the relative position.

//2. Constraints   → What does n tell you about target complexity?
    // 
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.


// Algorithm.
// 1. Convert T string to map of string and index.
// 2. Iterate through s.
//      If a char is not in tmap:
//             return False
//     If a char is present in tmap:
//            curPosition = tmap[char]
//             if curPosition < prevPosition:
//                 return False
//             else:
//                 prevPosition = curPosition

// tmap := make(map[byte]int)

// for i := 0; i < len(t); i++ {
//     tmap[t[i]] = i
// }

// prev := -1
// for j := 0; j < len(s); j++ {
//     cur, exists := tmap[s[j]]

//     if !exists {
//         return false
//     }

//     if cur < prev {
//         return false
//     } else {
//         prev = cur
//     }
// }

// return true

i := 0

for j := 0; j < len(t) && i < len(s); j++ {
    if s[i] == t[j] {
        i++
    }
}

return i == len(s)

//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//8. Code          → Only now. Trace each block as you write it.
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
