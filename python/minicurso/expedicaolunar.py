def expedicao(W, equipamentos):
    dp = [[0] * (W + 1) for _ in range(len(equipamentos) + 1)]
    for i in range(1, len(equipamentos) + 1):
        peso, dados = equipamentos[i-1]
        for w in range(1, W + 1):
            if peso <= w:
                dp[i][w] = max(dp[i-1][w], dp[i-1][w-peso] + dados)
            else:
                dp[i][w] = dp[i-1][w]
    return dp[-1][-1]
# Exemplo
W = 50
equipamentos = [(10, 60), (20, 100), (30, 120), (25, 80), (15, 50)]
print(expedicao(W, equipamentos)) 