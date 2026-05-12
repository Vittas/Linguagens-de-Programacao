#include <iostream>
#include <iomanip>
#include <queue>

using namespace std;

queue<string> filaDeTarefasNormais, filaDeTarefasPrioritaria;



void addTarefaNormal(string tarefa){
    filaDeTarefasNormais.push(tarefa);
}

void addTarefaPrioritaria(string tarefa){
    filaDeTarefasPrioritaria.push(tarefa);
}

void statusDasFilas(queue<string> filatemporaria){
    int numeroDeTarefas = 1;
    if(!filatemporaria.empty()){
        while (!filatemporaria.empty())
        {
            string tarefa = filatemporaria.front();
            cout << numeroDeTarefas << ". "<< tarefa << "\n\n";
            // cout << tarefa;
            filatemporaria.pop();
            numeroDeTarefas++;
        }
    }
    else{
        cout << "(vazia)\n";
    }
}

int main(){
    int escolha;
    string  tarefaNormal, tarefaPrioritaria;

    while(true){
        cout << "1. Adicionar tarefa normal\n 2. Adicionar tarefa prioritária\n3. Processar próxima tarefa\n4. Ver status das filas\n5. Sair\nDigite uma opção: ";
        cin >> escolha;
        switch (escolha)
        {
            case 1:
                cout << "Digite a descrição da tarefa normal: ";
                cin >> tarefaNormal;
                addTarefaNormal(tarefaNormal);
                if(!filaDeTarefasNormais.empty()){
                    cout << "Tarefa normal "<< '"' <<tarefaNormal << '"' << " adicionada.\n";
                }
                else{
                    cout << "Falha ao adcionar a tarefa.\n Tente Novamente!\n";
                }
                break;

            case 2:
                cout << "Digite a descrição da tarefa prioritaria: ";
                cin >> tarefaNormal;
                addTarefaPrioritaria(tarefaNormal);
                if(!filaDeTarefasPrioritaria.empty()){
                    cout << "Tarefa prioritaria "<< tarefaPrioritaria << " adicionada.\n";
                }
                else{
                    cout << "Falha ao adcionar a tarefa.\n Tente Novamente!\n";
                }
                break;

            case 3:
                if(!filaDeTarefasPrioritaria.empty()){

                    cout << "Processando tarefa prioritária: " << tarefaPrioritaria.front() << "\n";
                    filaDeTarefasPrioritaria.pop();

                }
                else{
                    cout << "Processando tarefa normal: " << tarefaNormal.front() << "\n";
                    filaDeTarefasNormais.pop();

                }
                break;
            case 4:
                cout << "Status das filas: \n ";
                cout << "Tarefas prioritárias: \n";
                statusDasFilas(filaDeTarefasPrioritaria);
                cout << "Tarefas normais: \n";
                statusDasFilas(filaDeTarefasNormais);
                break;
            case 5:
                return 0;

            default:
                break;
        }
    }  
} 

