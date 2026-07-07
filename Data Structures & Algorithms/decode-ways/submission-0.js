class Solution {
    /**
     * @param {string} s
     * @return {number}
     */
    numDecodings(s) {
        // Edge Case: If it starts with 0, no ways possible.
    if (s[0] === '0') return 0;

    const n = s.length;
    // dp[i] = ways to decode string of length i
    const dp = new Array(n + 1).fill(0);

    dp[0] = 1; // Empty string base case
    dp[1] = 1; // We already checked s[0] !== '0', so length 1 is valid

    for (let i = 2; i <= n; i++) {
        // 1. Single Digit Check
        // We look at the character at index i-1 (since dp is 1-indexed)
        const oneDigit = parseInt(s.substring(i - 1, i));
        if (oneDigit >= 1 && oneDigit <= 9) {
            dp[i] += dp[i - 1];
        }

        // 2. Double Digit Check
        // We look at characters at i-2 and i-1
        const twoDigits = parseInt(s.substring(i - 2, i));
        if (twoDigits >= 10 && twoDigits <= 26) {
            dp[i] += dp[i - 2];
        }
    }

    return dp[n];
    }
}
