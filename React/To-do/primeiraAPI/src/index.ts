import express, { Request, Response } from 'express';
import cors from 'cors'; // Importar cors antes de usá-lo

const app = express(); // Inicializar o app antes de usar
app.use(cors()); // Habilitar CORS
app.use(express.json()); // Permitir que o Express entenda JSON

const PORT = 3001; // Ajustar a porta se necessário

interface Todo {
    id: number;
    title: string;
    description: string;
    completed: boolean;
}

let todos: Todo[] = [
    { id: 1, title: 'Primeira API', description: 'Minha primeira vez fazendo uma api em typescript', completed: false },
    { id: 2, title: 'Primeira entrevista', description: 'Minha primeira vez participando de uma entrevista de emprego', completed: false }
];

/////////// Métodos //////////

/// Todas as tarefas
app.get('/todos', (req: Request, res: Response) => {
    res.json(todos);
});

/// Tarefa por ID
app.get('/todos/:id', (req: Request, res: Response) => {
    const id = parseInt(req.params.id);
    const todo = todos.find(t => t.id === id);
    todo ? res.json(todo) : res.status(404).json({ message: 'Tarefa não encontrada' });
});

/// Cria uma tarefa
app.post('/todos', (req: Request, res: Response) => {
    const { title, description, completed = false } = req.body;
    const newTodo: Todo = { id: todos.length + 1, title, description, completed };
    todos.push(newTodo);
    res.status(201).json(newTodo);
});

/// Atualiza uma tarefa
app.put('/todos/:id', (req: Request, res: Response) => {
    const id = parseInt(req.params.id);
    const { title, description, completed } = req.body;
    const todo = todos.find(t => t.id === id);

    if (todo) {
        todo.title = title ?? todo.title;
        todo.description = description ?? todo.description;
        todo.completed = completed ?? todo.completed;
        res.json(todo);
    } else {
        res.status(404).json({ message: 'A tarefa não foi encontrada' });
    }
});

/// Deleta uma tarefa
app.delete('/todos/:id', (req: Request, res: Response) => {
    const id = parseInt(req.params.id);
    todos = todos.filter(t => t.id !== id);
    res.status(204).send();
});

app.listen(PORT, () => {
    console.log(`Servidor iniciando com sucesso, rodando em http://localhost:${PORT}`);
});
