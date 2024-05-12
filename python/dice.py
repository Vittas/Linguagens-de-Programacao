import random


vlr_dado =  input("Digite o valor do dado: ")


dado =  vlr_dado.split("d") 

#print (dado)

x = int(dado[0])

dado = int(dado[1])

soma = 0 

for i in range(x):

    resultado = random.randint(1 , dado)
 
    print("O resultado do",vlr_dado,"foi:",resultado)

    soma = soma + resultado

print("\nO resultado foi:",soma)



    
    

    
 





