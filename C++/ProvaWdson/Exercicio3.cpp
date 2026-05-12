#include <iostream>
#include <iomanip>
// #include <Windows.h>
#include <stack>
#include <queue>
 
using namespace std;
 
queue<string> filaPacientes;
int tamanhoFila;
stack<string> historicoPacientes;
int tamanhoPilha;
 
void inputText(string &varText){
    cin.ignore();
    getline(cin, varText);
}
 
void adicionarPaciente(string nome){
    filaPacientes.push(nome);
}
 
void chamarProximoPaciente(){
   	tamanhoFila = filaPacientes.size();
    if(tamanhoFila > 0){
    	cout<< "Paciente chamado: " << filaPacientes.front() << "\n";
        historicoPacientes.push(filaPacientes.front());
        filaPacientes.pop();

    }
    else{
        cout << "Não há nenhum paciente na fila\n";
    }
}


void mostrarProximoPaciente(){
    cout << filaPacientes.front();
}

void mostrarHistoricoPacientes(){
    if (!historicoPacientes.empty()){
        cout << "Histórico de pacientes chamados:" << endl;
        while (!historicoPacientes.empty()) {
            cout << historicoPacientes.top() << endl;
            historicoPacientes.pop();
        }
    }
    else {
        cout << "Nenhum paciente foi chamado ainda." << endl;
    }
}
 
int main(){
    // SetConsoleOutputCP(CP_UTF8);
    string nomePaciente;
 
    queue<string> filaTemp;
    int escolha;
 
    while (true)
    {
        cout << "\n1- Adicionar paciente\n" << "2- chamar próximo paciente\n" << "3- Mostrar o próximo paciente\n" << "4- Número de pacientes na fila\n" << "5- Sair\n";
        cin >> escolha;
        switch (escolha)
        {
        case 1:
            cout << "Insira o nome do paciente:\n";
            inputText(nomePaciente);
            adicionarPaciente(nomePaciente);
            break;
        case 2:
 
            chamarProximoPaciente();
            filaTemp = filaPacientes;
            tamanhoFila = filaPacientes.size();
        cout << "Pacientes na fila: \n";
        	for(int i = 0; i < tamanhoFila; i++){
                cout << i+1 <<"- "<< filaTemp.front() << "\n";
                filaTemp.pop();
            }        
            break;
        case 3:
            mostrarProximoPaciente();
            break;

        case 4:
            tamanhoFila = filaPacientes.size();
            cout << "O número de pacientes na fila é: "<< tamanhoFila << "\n";
            break;

        case 5:
            mostrarHistoricoPacientes();
            return 0;
        }
    }
}