preco = float(input('Digite o preço do produto: '))
dep =  int(input("""Digite qual departamento você deseja acessar:
            
             [1] - Departamento A
             [2] - Departamento B
             [3] - Departamento C
             
             Digite o número do departamento desejado: """))

match dep:
    case 1:
        cor = int(input("""Digite qual departamento você deseja acessar:
            
             [1] - Azul
             [2] - Branca
             [3] - Verde
             [4] - Preta
             
             Digite o número do departamento desejado: """))
        
        match cor:
            case 1:
                desconto = preco - (preco * 0.06)
                print("O preço do produto com desconto é de ", desconto)
            case 2:
                desconto = preco - (preco * 0.07)
                print("O preço do produto com desconto é de ", desconto)
            case 3:
                desconto = preco - (preco * 0.08)
                print("O preço do produto com desconto é de ", desconto)
            case 4:
                desconto = preco - (preco * 0.09)
                print("O preço do produto com desconto é de ", desconto)
            case _:
                raise ValueError("Erro")

            
    
    case 2:
        cor = int(input("""Digite qual departamento você deseja acessar:
            
             [1] - Azul
             [2] - Branca
             [3] - Verde
             [4] - Preta
             
             Digite o número do departamento desejado: """))
        
        match cor:
            case 1:
                desconto = preco - (preco * 0.63)
                print("O preço do produto com desconto é de ", desconto)
            case 2:
                desconto = preco - (preco * 0.74)
                print("O preço do produto com desconto é de ", desconto)
            case 3:
                desconto = preco - (preco * 0.82)
                print("O preço do produto com desconto é de ", desconto)
            case 4:
                desconto = preco - (preco * 0.91)
                print("O preço do produto com desconto é de ", desconto)
            case _:
                raise ValueError("Erro")
        
    
    case 3: 
        cor = int(input("""Digite qual departamento você deseja acessar:
            
             [1] - Azul
             [2] - Branca
             [3] - Verde
             [4] - Preta
             
             Digite o número do departamento desejado: """))
        
        match cor:
            case 1:
                desconto = preco - (preco * 0.56)
                print("O preço do produto com desconto é de ", desconto)
            case 2:
                desconto = preco - (preco * 0.67)
                print("O preço do produto com desconto é de ", desconto)
            case 3:
                desconto = preco - (preco * 0.078)
                print("O preço do produto com desconto é de ", desconto)
            case 4:
                desconto = preco - (preco * 0.109)
                print("O preço do produto com desconto é de ", desconto)
            case _:
                raise ValueError("Erro")

    case _ :
        raise ValueError("Erro")
