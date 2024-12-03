import re

arq = open("2024_03.txt")
input = arq.read()

# input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
regs = re.findall(r"mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)", input)
# print(regs)
valor = 0
for i in regs:
    if i[0] == "m":
        nums = re.findall(r"[0-9]{1,3}", i)
        valor += int(nums[0]) * int(nums[1])
        # print(nums)

print(valor)

valor2 = 0
ativo = True
for i in regs:
    if i[2] == "n":
        ativo = False
    if i[2] == "(":
        ativo = True
    if i[0] == "m" and ativo:
        nums = re.findall(r"[0-9]{1,3}", i)
        valor2 += int(nums[0]) * int(nums[1])

print(valor2)
