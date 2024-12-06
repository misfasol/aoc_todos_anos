arq = open("2024_04.txt")
linhas = [l.strip() for l in arq.readlines()]

# linhas = [
#     "MMMSXXMASM",
#     "MSAMXMSMSA",
#     "AMXSXMAAMM",
#     "MSAMASMSMX",
#     "XMASAMXAMM",
#     "XXAMMXXAMA",
#     "SMSMSASXSS",
#     "SAXAMASAAA",
#     "MAMMMXMMMM",
#     "MXMXAXMASX",

# PARTE 1
contador = 0
for linha in range(len(linhas)):
    for col in range(len(linhas[0])):
        for i in [-1, 0, 1]:
            for j in [-1, 0, 1]:
                string = ""
                for c in [0, 1, 2, 3]:
                    if linha + i * c < 0 or col + j * c < 0:
                        break
                    try:
                        string += linhas[linha + i * c][col + j * c]
                    except:
                        break
                if string == "XMAS":
                    contador += 1

print(contador)

# PARTE 2

contador = 0
comb = ["SM", "MS"]
for linha in range(len(linhas)):
    for col in range(len(linhas[0])):
        if linhas[linha][col] == "A":
            try:
                if (linhas[linha-1][col-1] + linhas[linha+1][col+1]) in comb and (linhas[linha-1][col+1] + linhas[linha+1][col-1]) in comb:
                    contador += 1
            except:
                pass

print(contador)
