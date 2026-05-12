#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

vector<int> nodeVector;

struct node
{
    int data;
    node *left;
    node *right;
};

node *createNode(int value)
{
    node *newNode = new node();
    newNode->data = value;
    newNode->left = nullptr;
    newNode->right = nullptr;
    return newNode;
}

int getSize(node *root)
{
    if (root == nullptr)
    {
        return 0;
    }
    return getSize(root->left) + getSize(root->right) + 1;
}

void printNodeData(node *root)
{
    if (root == nullptr)
    {
        return;
    }
    nodeVector.push_back(root->data);
    cout << root->data << " ";
    printNodeData(root->left);
    printNodeData(root->right);
}

node *searchNodeData(node *root, int nodeData)
{
    if (root == nullptr || root->data == nodeData)
    {
        return root;
    }
    return (root->data > nodeData)
               ? searchNodeData(root->left, nodeData)
               : searchNodeData(root->right, nodeData);
}

void searchAndPrintNode(node *root, int nodeData)
{
    node *result = searchNodeData(root, nodeData);
    if (result != nullptr)
    {
        cout << "Valor encontrado: " << result->data << endl;
    }
    else
    {
        cout << "Valor NÃO encontrado!" << endl;
    }
}

void sortingNodeVector()
{
    sort(nodeVector.begin(), nodeVector.end());
}

void deleteTree(node *root)
{
    if (root == nullptr)
    {
        return;
    }
    deleteTree(root->left);  // Deleta a subárvore esquerda
    deleteTree(root->right); // Deleta a subárvore direita
    delete root;             // Deleta o nó atual
    root = nullptr;          // Boa prática: evitar ponteiro dangling
}

int main()
{
    node *root = createNode(50);
    root->left = createNode(20);
    root->right = createNode(70);
    root->left->left = createNode(10);
    root->left->right = createNode(30);
    root->right->left = createNode(60);
    root->right->right = createNode(80);

    cout << "Valores dos nós: ";
    printNodeData(root);
    cout << endl;

    sortingNodeVector();
    cout << "Valores em ordem crescente: ";
    for (int num : nodeVector)
    {
        cout << num << " ";
    }
    cout << endl;

    int numeroBusca;
    cout << "Digite um número para buscar na árvore: ";
    cin >> numeroBusca;
    cout << "Buscando o numero " << numeroBusca << " na árvore:\n";
    searchAndPrintNode(root, numeroBusca);

    cout << "Tamanho da árvore: " << getSize(root) << "\n";

    cout << "Deletando a árvore...\n";
    deleteTree(root);
    return 0;
}