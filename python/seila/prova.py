import ast #usa um eval especial que ira tornar as informações do arquivo txt em dicionario novamente
import datetime #permite usar medidas de tempo
banco_emprestimo = [] #cria a lista que vai armazenar tudo sobre os emprestimos
banco_livros = [] # cria a lista que vai armazenar tudo sobre os livros 


def carregar_dados():
    try:

        arquivo_livros = open("dadoLivros.txt", "r")  #Tenta abrir um arquivo
        arquivo_emprestimo = open("dadoEmprestimo.txt", "r")

        for line in arquivo_livros: #Um for que vai andar pelas linhas do arquivo, se elas forem igual ou menor que 1 ira continuar
            if len(line) <= 1:
                continue
            banco_livros.append(ast.literal_eval(line)) #a lista banco_livros recebe todas as informações das linhas do arquivo txt

        
        for line in arquivo_emprestimo:
            if len(line) <= 1:
                continue
            banco_emprestimo.append(ast.literal_eval(line))#a lista banco_emprestimos recebe todas as informações das linhas do arquivo txt

        #Fecha os arquivos

        arquivo_livros.close() 
        arquivo_emprestimo.close()

        #Teste para ver se estava funcionando
        # print (f"Banco Livros: {banco_livros}")
        # print (f"Banco emprestimo: {banco_emprestimo}")


    except Exception as error: #caso algo de errado, irá exibir o erro
        print (error)



carregar_dados() #executa a função antes de tudo, para carregar os arquivos

def salvar_dados_emprestimo(emprestimos): #função para salvar o emprestimo em um txt, usando de parametro a variavel emprestimos (armazena a informação do usuario, livro etc)
    arquivo_emprestimo = open("dadoEmprestimo.txt", "a") #cria ou abre um arquivo txt, e irá adicionar as informações passadas
    arquivo_emprestimo.write(f"{emprestimos}"+ '\n')    #adiciona a variavel emprestimos na qual armazeno tudo sobre quem fez, qual livro etc
    arquivo_emprestimo.close() #fecha o arquivo

def salvar_dados_livros(): #função para salvar o os livros em um txt, usando de parametro a variavel emprestimos (armazena a informação do titulo, autor etc)
    arquivo_livros = open("dadoLivros.txt", "a") #cria ou abre um arquivo txt, e irá adicionar as informações passadas
    arquivo_livros.write(f"{livros}" + '\n') #adiciona a variavel livro na qual armazeno tudo sobre os livros(titulo, autor, etc)
    arquivo_livros.close() #fecha o arquivo

def cadastrar(livros):  #função para cadastrar os livros
    banco_livros.append(livros) #adiciona a variavel livros, na lista banco_livros
    salvar_dados_livros() #chamaa função salvar_dados_livros()
   
 
def consultar(): #função para realizar a consulta de todos os livros e suas informações
    n = 0 
    for x in banco_livros: #for que vai andar pelo indice da lista de livros
        result = "" #variavel que irá armazenar as informações de um livro

        for key in banco_livros[n].keys(): # for que vai andar pelos itens do dicionário especifico 
            # print(f"""{banco_livros[n][key]} id = {n +1}""")
            result += f"[{banco_livros[n][key]}] " #a variavel vai somar todas as informações de um dicionario especifico
        n += 1 # variavel que armazena um id de um livro
        print(f"{result} | id [{n}]")
 
def pegar_emprestimo(nome, livro, data_emprestimo): #função que realiza um emprestimo usando os parametros, nome do usuario, livro(id do livro escolhido), e data do emprestimo
    livro_escolhido = banco_livros[livro-1] # aqui a variavel ira armazenar todas as informações de um livro escolhido
    emprestimos = {"titulo":livro_escolhido["titulo"], "nome": nome , "codigo": livro_escolhido["codigo"], "data": data_emprestimo.strftime("%d/%m/%y")} # a variavel emprestimos ira armazenar tudo sobre a pessoa, o codigo do livro e a data do emprestimo

    banco_emprestimo.append(emprestimos) #adiciona todas as informações de um determinado emprestimo para a lista banco_emprestimo
    salvar_dados_emprestimo(emprestimos) #chama a função salvar_dados_emprestimo()
    # relatorio_emprestimo()

def relatorio_emprestimo(): #função que irá mostrar todos o emprestimos feitos até o momento
    n = 0 #variavel que ira servir como id
    
    print("EMPRESTIMOS REALIZADOS") #formatação
    tempo_agora = datetime.datetime.now() #pega o tempo atual
    tempo_agora = tempo_agora.strftime("%d/%m/%y") # mexe no tempo atual
    # data_emprestimo =  banco_emprestimo.keys()
    for x in banco_emprestimo: #for que vai andar pelo indice da lista de emprestimo
        
        result = "" #variavel que ira armazenar e mostrar os emprestimos já feitos

        for key in banco_emprestimo[n].keys(): #for que vai andar pelos itens especificos do dicionario de um determinado emprestimo
            # print(f"""{banco_livros[n][key]} id = {n +1}""")
            result += f"[{banco_emprestimo[n][key]}] " #a variavel vai somar e formatar as informações de um determinado dicionario de emprestimo
            data = banco_emprestimo[n][key] #essa variavel ira mostrar a data feita quando houve o emprestimo
        n += 1
        
        print(f"{result} | id [{n}] | Emprestimo feito a {data}") #vai mostrar todas as informações sobre o emprestimo e a data em que eles foram realizados


nome = input("Cadastre seu nome de usuário: ") #variavel que vai armazenar o nome do usuario
while True: #while que ira executar tudo sobre o codigo
#logo a baixo crio uma variavel que servirá para decidir as ações do usuário
    opcoes = int(input("""
                       |1| Cadastrar livro  |2| Consultar livros 
                       |3| Realizar emprestimo  |4| Relatório emprestimo
                                |5| Sair 

                       Digite:
                       """))
    match opcoes:
        case 1: #caso a variavel opções seja 1 irá ocorrer outro while
            while True: #nesse while coleto as principais informações dos livros (titulo, autor, ano e codigo)
                titulo = str(input("Digite o titulo do livro: " )) 
                autor = str(input("Digite o autor do livro: " ))
                ano = str(input("Digite o ano do livro: " ))
                cod = str(input("Digite o codigo do livro: " ))
                livros = {"titulo": titulo, "autor": autor, "ano": ano, "codigo": cod} #crio um dicionario que armazena todas as informações para o registro do livro

                cadastrar(livros) #chama a função cadastrar e envia o parametro sendo ele a variavel livros
                s = input("continuar a cadastrar? sim | não: ") #caso deseje continuar a cadastrar ele ira continuar o while
                if s != "sim": #se não ira voltar para a tela de decisão do usuario
                    break
        case 2: #caso digite 2 irá ser mostrado todas  informações sobre os livros registrados
    
                print(consultar())
        case 3: #caso digite 3 ira executar um while

            while True: #esse while ira coletar o id do livro desejado para emprestimo
                livro = int(input("Qual livro deseja pegar? | Digite o id do livro: "))
                data_emprestimo = datetime.datetime.now() # aqui a variavel armazena a data do emprestimo
                pegar_emprestimo(nome, livro, data_emprestimo) # chama a função pegar_emprestrimo() usand os parametros nome do usuario, id do livro desejado e data do emprestimo realizado
                emprestar = input("deseja realizar outro emprestimo? sim | não: ") #caso deseje continuar, o while irá continuar
                if emprestar != "sim": # caso não, ele ira parar
                    break

        case 4: # chama a função relatorio_emprestimo, mostrando todos os emprestimos já feitos
            relatorio_emprestimo() 
        
        case 5: # caso digite 5 ele ira perguntar se quer encerrar o programa
            sair = input("Deseja sair? sim | nao: ") # caso não deseje, ele ira continuar o while principal
            if sair == "sim": 
                break
 