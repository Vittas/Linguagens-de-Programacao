import api from './Api';

export interface Todo {
    id: number,
    title: string,
    description: string,
    completed: boolean
}

export const fetchTodos = async () => {
    const response = await api.get<Todo[]>('/todos');
    return response.data;

};
/// FAZER UM FETCH POR ID E UM DELETE   
export const fetchTodoById = async (id: number) => {
    const response = await api.get<Todo>(`/todos/${id}`);
    return response.data;
};
 
export const deleteTodo = async (id: number) => {
    await api.delete(`/todos/${id}`);
    return id; // Retorna o ID da tarefa deletada, se necessário
};
export const createTodo = async (title:string, description:string) => {
    const response = await api.post<Todo>('/todos', {title, description,complete: false});
    return response.data;
};