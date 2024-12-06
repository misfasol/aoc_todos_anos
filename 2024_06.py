conteudo = [list(l.strip()) for l in open("2024_06.txt").readlines()]

# conteudo = [list(l) for l in [
#     "....#.....",
#     ".........#",
#     "..........",
#     "..#.......",
#     ".......#..",
#     "....#.....",
#     ".#..^#....",
#     "........#.",
#     "#.........",
#     "......#...",
# ]]

CIMA = (-1, 0)
DIRE = (0, 1)
BAIX = (1, 0)
ESQU = (0, -1)

def somDir(a, b):
    return (a[0] + b[0], a[1] + b[1])

def rota(a):
    if a == CIMA:
        return DIRE
    elif a == DIRE:
        return BAIX
    elif a == BAIX:
        return ESQU
    elif a == ESQU:
        return CIMA

def direChar(a):
    if a == CIMA:
        return "^"
    elif a == DIRE:
        return ">"
    elif a == BAIX:
        return "v"
    else:
        return "<"

def resolver(conteudo: list[str]):
    xa, ya = 0, 0
    direAtu = CIMA
    for linha in range(len(conteudo)):
        for col in range(len(conteudo[0])):
            if "^" == conteudo[linha][col]:
                xa = linha
                ya = col

    se = set()

    while True:
        
        se.add((xa, ya))
        conteudo[xa][ya] = direChar(direAtu)
        xf, yf = somDir((xa, ya), direAtu)
        if xf < 0 or xf >= len(conteudo) or yf < 0 or yf >= len(conteudo[0]):
            return (len(se), None)
        contador = 0
        while conteudo[xf][yf] == "#":
            direAtu = rota(direAtu)
            xf, yf = somDir((xa, ya), direAtu)
            contador += 1
            if contador == 4:
                return (None, True)
        else:
            xa, ya = xf, yf
        if conteudo[xa][ya] == direChar(direAtu):
            return (None, True)

#     se.add((xa, ya))
#     xf, yf = somDir((xa, ya), direAtu)
#     if xf < 0 or xf >= len(conteudo) or yf < 0 or yf >= len(conteudo[0]):
#         break
#     oqtafrente = conteudo[xf][yf]
#     if oqtafrente == "#":
#         direAtu = rota(direAtu)
#     conteudo[xa][ya] = direChar(direAtu)
#     if oqtafrente == direChar(direAtu):
#         return (None, True)
#     xa, ya = somDir((xa, ya), direAtu)
# return (len(se), False)

final1, _ = resolver([l.copy() for l in conteudo])

print(f"final1: {final1}")

final2 = 0

for l in range(len(conteudo)):
    for c in range(len(conteudo[0])):
        nova = [l.copy() for l in conteudo]
        if nova[l][c] == ".":
            nova[l][c] = "#"
            _, eloop = resolver(nova)
            if eloop:
                final2 += 1
    print(f"linha: {l}")

print(f"final2: {final2}")
