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

regras: list[str] = []
contador = 0
for linha in conteudo:
    contador += 1
    if linha == "":
        break
    regras.append(linha)

entradas: list[str] = []
while contador < len(conteudo):
    entradas.append(conteudo[contador])
    contador += 1

final = 0
final2 = 0

def emOrdem(paginas: list[str]) -> bool:
    for pag in range(len(paginas)):
        for regra in regras:
            if regra.startswith(paginas[pag]):
                fim = regra[3:5]
                for i in range(len(paginas)):
                    if paginas[i] == fim:
                        if i < pag:
                            return False
    return True

for entrada in entradas:
    paginas = entrada.split(",")
    valido = emOrdem(paginas)
    if valido:
        final += int(paginas[len(paginas)//2])
    else:
        while not valido:
            for pag in range(len(paginas)):
                for regra in regras:
                    if regra.startswith(paginas[pag]):
                        fim = regra[3:5]
                        for i in range(len(paginas)):
                            if paginas[i] == fim:
                                if i < pag:
                                    temp = paginas[i]
                                    teim = paginas[pag]
                                    paginas[i] = teim
                                    paginas[pag] = temp
                                    valido = emOrdem(paginas)
        final2 += int(paginas[len(paginas)//2])
            


print(f"final : {final}")  # 143
print(f"final2: {final2}") # 123
