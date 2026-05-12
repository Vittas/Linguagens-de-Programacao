#include <iostream>
 
using namespace std;
 
struct node {
    int data;
    node* next;
};
 
void inputFront(node* &head, int value) {
    node* newNode = new node();
    newNode->data = value;
    newNode->next = head;
    head = newNode;
}
 

void printLista(node* head){
    node* atual = head;
    while (atual != nullptr){
        cout << atual->data;
        if(atual->next != nullptr){
            cout << ", ";
        }    
        atual = atual->next;
    }
    cout << "\n";
}
 
void inverteLista(node* &head){
    node* anterior = nullptr;
    node* atual = head;
    node* next = nullptr;
    while(atual != nullptr){
        next = atual->next;
        atual->next = anterior;
        anterior = atual;
        atual = next;
    }
    head = anterior;
}
 
void popLista(node* &head, int value) {
    if (head == nullptr) {
        cout << "Lista vazia." << endl;
        return;
    }
    node* atual = head;
    node* anterior = nullptr;
    while (atual != nullptr && atual->data != value) {
        anterior = atual;
        atual = atual->next;
    }
    if (atual == nullptr) {
        cout << "Valor não encontrado." << endl;
        return;
    }
    if (anterior == nullptr) {
        head = atual->next;
    }
    else {
        anterior->next = atual->next;
    }
    delete atual;
    cout << "Removido com sucesso." << endl;
}
 
int main(){
    // SetConsoleOutPutCP(CP_UTF8);

    node* head = nullptr;
 
    inputFront(head, 10);
    inputFront(head, 20);
    inputFront(head, 30);
    inputFront(head, 40);
    inputFront(head, 50);
 
    printLista(head);
    popLista(head, 30);
    cout << "Lista encadeada: ";
    printLista(head);
    inverteLista(head);
    cout << "Lista encadeada invertida: ";
    printLista(head);
 
    return 0;
}
 