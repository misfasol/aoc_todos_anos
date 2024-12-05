arq = open("2024_05.txt")
conteudo = [l.strip() for l in arq.readlines()]

# conteudo = [
#     "47|53",
#     "97|13",
#     "97|61",
#     "97|47",
#     "75|29",
#     "61|13",
#     "75|53",
#     "29|13",
#     "97|29",
#     "53|29",
#     "61|53",
#     "97|53",
#     "61|29",
#     "47|13",
#     "75|47",
#     "97|75",
#     "47|61",
#     "75|61",
#     "47|29",
#     "75|13",
#     "53|13",
#     "",
#     "75,47,61,53,29",
#     "97,61,53,29,13",
#     "75,29,13",
#     "75,97,47,61,53",
#     "61,13,29",
#     "97,13,75,29,47"
# ]

# print(conteudo[1169:1181])

regras = []
contador = 0
for linha in conteudo:
    contador += 1
    if linha == "":
        break
    regras.append(linha)

entradas = []
while contador < len(conteudo):
    entradas.append(conteudo[contador])
    contador += 1

# print(regras)
# print(entradas)

final = 0

for entrada in entradas:
    paginas = entrada.split(",")
    print(entrada, paginas)
    for i in paginas:
        print(i)

    break

print(f"final: {final}")
