
#include <iostream>
#include <iomanip>
// #include <Windows.h>
#include <vector>
#include <algorithm>

using namespace std;

struct Aluno{
    string nome;
    int matricula;
    vector<float> notas[3];
};

void inputText(string &varText){
    cin.ignore();
    getline(cin, varText);
}

int main(){
    // SetConsoleOutPutCP(CP_UTF8);
    
    vector<Aluno> vetorAlunos;
    vector<float> mediaAlunos;
    float media;

    for(int i = 0; i < 5; i++){
        Aluno aluno;
        cout << "nome do aluno:\n";
        inputText(aluno.nome);
        cout << "matricula do aluno:\n";
        cin >> aluno.matricula;
        for(int i = 0; i < 3; i++){
            float nota;
            cout<< "nota" << i+1 <<":\n";
            cin >> nota;
            media += nota;
            aluno.notas->push_back(nota);
        }
        media = media/3;
        mediaAlunos.push_back(media);
        media = 0;
        vetorAlunos.push_back(aluno);
    };

    for(int i = 0; i < 5; i++){
        cout << "Nome: " << vetorAlunos[i].nome << "\n";
        cout << "Matricula: " << vetorAlunos[i].matricula << "\n";
        cout << "media: " << mediaAlunos[i] << "\n";
    };

    auto maxMedia = max_element(mediaAlunos.begin(), mediaAlunos.end());

    int index = distance(mediaAlunos.begin(), maxMedia);
    cout << "O aluno com maior média é: " << vetorAlunos[index].nome << ", com a média de: " << *maxMedia;
    return 0;
};
