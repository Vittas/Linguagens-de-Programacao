c, n = list(map(int, input().split(" ")))
variavel = c/n
variavelVerdade = c//n
if variavel == variavelVerdade:
    print(0)
else:
    variavelUtil = n * variavelVerdade
    variavelUtil = c - variavelUtil
    print(variavelUtil)