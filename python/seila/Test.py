import ast
import datetime
banco_emprestimo = []
banco_livros = []


def carregar_dados():
    try:

        arquivo_livros = open("dadoLivros.txt", "r")
        arquivo_emprestimo = open("dadoEmprestimo.txt", "r")

        for line in arquivo_livros:
            if len(line) <= 1:
                continue
            banco_livros.append(ast.literal_eval(line))

        
        for line in arquivo_emprestimo:
            if len(line) <= 1:
                continue
            banco_emprestimo.append(ast.literal_eval(line))

        arquivo_livros.close()
        arquivo_emprestimo.close()

        print (f"Banco Livros: {banco_livros}")
        print (f"Banco emprestimo: {banco_emprestimo}")


    except Exception as error:
        print (error)



carregar_dados()

def salvar_dados_emprestimo(emprestimos):
    arquivo_emprestimo = open("dadoEmprestimo.txt", "a")
    arquivo_emprestimo.write(f"{emprestimos}"+ '\n')
    arquivo_emprestimo.close()

def salvar_dados_livros():
    arquivo_livros = open("dadoLivros.txt", "a")
    arquivo_livros.write(f"{livros}" + '\n')
    arquivo_livros.close()

def cadastrar(livros):
    banco_livros.append(livros)
    salvar_dados_livros()
   
 
def consultar():
    n = 0
    for x in banco_livros:
        result = ""

        for key in banco_livros[n].keys():
            # print(f"""{banco_livros[n][key]} id = {n +1}""")
            result += f"[{banco_livros[n][key]}] "
        n += 1
        print(f"{result} | id [{n}]")
 
def pegar_emprestimo(nome, livro, data_emprestimo):
    livro_escolhido = banco_livros[livro-1]
    emprestimos = {"titulo":livro_escolhido["titulo"], "nome": nome , "codigo": livro_escolhido["codigo"], "data": data_emprestimo.strftime("%d/%m/%y")}

    banco_emprestimo.append(emprestimos)
    salvar_dados_emprestimo(emprestimos)
    # relatorio_emprestimo()

def relatorio_emprestimo():
    n = 0
    
    print("EMPRESTIMOS REALIZADOS")
    tempo_agora = datetime.datetime.now()
    tempo_agora = tempo_agora.strftime("%d/%m/%y")
    # data_emprestimo =  banco_emprestimo.keys()
    for x in banco_emprestimo:
        
        result = ""

        for key in banco_emprestimo[n].keys():
            # print(f"""{banco_livros[n][key]} id = {n +1}""")
            result += f"[{banco_emprestimo[n][key]}] "
            data = banco_emprestimo[n][key]
        n += 1
        
        print(f"{result} | id [{n}] | Emprestimo feito a {data}")


nome = input("Cadastre seu nome de usuário: ")
while True:

    opcoes = int(input("""
                       |1| Cadastrar livro  |2| Consultar livros 
                       |3| Realizar emprestimo  |4| Relatório emprestimo
                                |5| Sair 

                       Digite:
                       """))
    match opcoes:
        case 1:
            while True:
                titulo = str(input("Digite o titulo do livro: " ))
                autor = str(input("Digite o autor do livro: " ))
                ano = str(input("Digite o ano do livro: " ))
                cod = str(input("Digite o codigo do livro: " ))
                livros = {"titulo": titulo, "autor": autor, "ano": ano, "codigo": cod}

                cadastrar(livros)
                s = input("continuar a cadastrar? sim | não: ")
                if s != "sim":
                    break
        case 2:
    
                print(consultar())
        case 3:

            # emprestar = input("deseja realizar um emprestimo? sim | não: ")
            while True:
                livro = int(input("Qual livro deseja pegar? | Digite o id do livro: "))
                data_emprestimo = datetime.datetime.now()
                pegar_emprestimo(nome, livro, data_emprestimo)
                emprestar = input("deseja realizar outro emprestimo? sim | não: ")
                if emprestar != "sim":
                    break

        case 4:
            relatorio_emprestimo()
        
        case 5:
            sair = input("Deseja sair? sim | nao: ")
            if sair == "sim":
                break
 