conteudo = open("2024_06.txt").readlines()

conteudo = [
    "....#.....",
    ".........#",
    "..........",
    "..#.......",
    ".......#..",
    "..........",
    ".#..^.....",
    "........#.",
    "#.........",
    "......#...",
]

CIMA = (-1, 0)
DIRE = (0, 1)
BAIX = (1, 0)
ESQU = (0, -1)

def somDir(a, b):
    return (a[0] + b[0], a[1] + b[1])

x, y = 0, 0
direcao = BAIX
for linha in conteudo:
    if "^" in linha:
        for c in linha:
            if c == "^":
                break
            else:
                y += 1
        break
    else:
        x += 1

print(f"x: {x}, y: {y}")
frente = somDir((x, y), direcao)
print(frente)

# while x + direcao[0] >= 0 and x + direcao[1] < len(conteudo):

