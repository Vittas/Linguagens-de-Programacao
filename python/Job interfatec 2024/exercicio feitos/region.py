lines = int(input())
limites = []
sim = []
nao = []
for i in range(lines):
    enderecos = input()
    comeco,fim = enderecos.split(" ")
    fim = list(map(int,fim.split("-")))
    comeco = list(map(int,comeco.split("-")))
    limites.append({f"começo": comeco,f"fim":fim})

enderecos = int(input())
for endereco in range(enderecos):
    endereco = input()
    numeroMaior,numeroMenor = map(int,endereco.split("-"))
    for i in limites:
        if i["fim"][1] >= numeroMaior >= i["comeco"][0] and i["fim"][1] >= numeroMenor >= i["comeco"][1]:
            e = f"{numeroMaior}-{numeroMenor}"
            sim.append()
    if f"{numeroMaior}-{numeroMenor}" not in sim:
        nao.append(f"{numeroMaior}-{numeroMenor}")
        
print(sim)
print(nao)
    
    
    
    