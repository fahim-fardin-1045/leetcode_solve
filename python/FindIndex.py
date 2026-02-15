class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        n, m = len(haystack), len(needle)

        for i in range(n - m + 1):
            if haystack[i:i+m] == needle:
                return i
        return -1


def main():
    haystack = input("Enter haystack: ")
    needle = input("Enter needle: ")

    sol = Solution()               # Create object
    result = sol.strStr(haystack, needle)  # Call method properly
    print("Answer:", result)


if __name__ == "__main__":
    main()
