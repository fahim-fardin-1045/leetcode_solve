class Solution:
    def divide(self, a: int, b: int) -> int:
        if b == 1:
            return a
        if a == -(2**31) and b == -1:
            return 2**31 - 1
        
        sign = (a>0 and b>0) or (a<0 and b<0)

        a = -a if a>0 else a
        b = -b if b>0 else b
        ans = 0

        while a <= b :
            temp = b 
            c = 1 
            
            while temp >= (-(2**31)) and a<=(temp << 1):
                temp <<= 1
                c <<= 1
            a -= temp
            ans += c
        return ans if sign else -ans 


def main():
    a = int (input("enter value of a : "))
    b = int (input("enter value of b : "))

    sol = Solution()
    result = sol.divide(a,b)
    print("Result",result)


if __name__ == "__main__":
    main()
