import random


vlr_dado =  input("Digite o valor do dado: ")


dado =  vlr_dado.split("d") 

#print (dado)

dado = int(dado[1])

soma = 0 


resultado = random.randint(1 , dado)
 
print("O resultado do ",vlr_dado,"foi:",resultado)





    
    

    
 





