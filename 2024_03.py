import re

arq = open("2024_03.txt")
input = arq.read()

# input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
resultado = re.findall(r"mul\([0-9]{1,3},[0-9]{1,3}\)", input)
print(resultado)
valor = 0
for i in resultado:
    nums = re.findall(r"[0-9]{1,3}", i)
    valor += int(nums[0]) * int(nums[1])
    # print(nums)

print(valor)
