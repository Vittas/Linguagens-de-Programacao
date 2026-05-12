import pandas as pd

file = 'analiseDados/dadosAnalise.csv'
df = pd.read_csv(file)
df.columns = df.columns.str.strip()

print("Colunas disponíveis:", df.columns.tolist())

# -----------------------------
# 1. O que dizer sobre o desempenho?
print("\n1. Análise descritiva do desempenho:\n")
desc = df.describe()
print(desc)

# -----------------------------
# 2. Comportamento dos dados (tendência, dispersão, outliers)

print("\n2. Comportamento dos dados:\n")
for col in ['Produtividade', 'Experiencia', 'Erros_Reportados', 'Tempo_Conclusao']:
    print(f"\nAnalisando: {col}")
    print(f"Média: {df[col].mean():.2f}")
    print(f"Mediana: {df[col].median():.2f}")
    print(f"Desvio Padrão: {df[col].std():.2f}")
    
    # Detectar outliers pelo método IQR
    Q1 = df[col].quantile(0.25)
    Q3 = df[col].quantile(0.75)
    IQR = Q3 - Q1
    outliers = df[(df[col] < (Q1 - 1.5 * IQR)) | (df[col] > (Q3 + 1.5 * IQR))]
    
    print(f"Outliers detectados: {len(outliers)}")

# -----------------------------
# 3. Correlações
print("\n3. Correlações entre variáveis:\n")
corr = df.corr()
print(corr)

plt.figure(figsize=(8, 6))
sns.heatmap(corr, annot=True, cmap='coolwarm', fmt=".2f")
plt.title('Mapa de Correlação')
plt.tight_layout()
plt.savefig('correlacao.png')
plt.close()

# -----------------------------
# 4. Melhor medida para cada variável
print("\n4. Melhor medida de tendência para cada variável:\n")

for col in ['Produtividade', 'Experiencia', 'Erros_Reportados', 'Tempo_Conclusao']:
    skew = df[col].skew()
    print(f"{col}: Assimetria = {skew:.2f}")
    if abs(skew) < 0.5:
        print(f"-> Distribuição aproximadamente simétrica. Usar média.\n")
    else:
        print(f"-> Distribuição assimétrica. Usar mediana.\n")

# -----------------------------
# 5. Variáveis discrepantes
print("\n5. Variáveis discrepantes:\n")
for col in ['Produtividade', 'Experiencia', 'Erros_Reportados', 'Tempo_Conclusao']:
    skew = df[col].skew()
    outliers = df[(df[col] < (df[col].quantile(0.25) - 1.5 * (df[col].quantile(0.75) - df[col].quantile(0.25)))) |
                  (df[col] > (df[col].quantile(0.75) + 1.5 * (df[col].quantile(0.75) - df[col].quantile(0.25))))]
    
    if abs(skew) > 1 or len(outliers) > 10:
        print(f"{col} possui comportamento discrepante. Assimetria: {skew:.2f}, Outliers: {len(outliers)}")

# -----------------------------
# 6. Histogramas para simetria/assimetria
print("\n6. Gerando histogramas...\n")
plt.figure(figsize=(12, 8))

for idx, col in enumerate(['Produtividade', 'Experiencia', 'Erros_Reportados', 'Tempo_Conclusao'], 1):
    plt.subplot(2, 2, idx)
    sns.histplot(df[col], bins=10, kde=True)
    plt.title(f"Histograma - {col}")

plt.tight_layout()
plt.savefig('histogramas.png')
plt.close()

# -----------------------------
# Extra: Scatterplots das principais correlações

plt.figure(figsize=(12, 4))

plt.subplot(1, 3, 1)
sns.scatterplot(x='Experiencia', y='Erros_Reportados', data=df)
plt.title('Experiência x Erros')

plt.subplot(1, 3, 2)
sns.scatterplot(x='Experiencia', y='Tempo_Conclusao', data=df)
plt.title('Experiência x Tempo')

plt.subplot(1, 3, 3)
sns.scatterplot(x='Produtividade', y='Erros_Reportados', data=df)
plt.title('Produtividade x Erros')

plt.tight_layout()
plt.savefig('scatterplots.png')
plt.close()

print("Análises concluídas! Gráficos salvos: histogramas.png, correlacao.png, scatterplots.png")
