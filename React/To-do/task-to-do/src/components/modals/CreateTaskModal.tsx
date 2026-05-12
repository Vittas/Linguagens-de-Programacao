// import '../../App.css';
import { createTodo, Todo } from '../../services/ApiTaskManager';
import { useState } from 'react';


interface Visibility{
    visibility: boolean
}
export const CreateTaskModal = ({visibility}: Visibility) => {
    const [newTodoTitle, setNewTodoTitle] = useState('');
    const [newTodoDescription, setNewTodoDescription] = useState('');
    const[todos, setTodos] = useState<Todo[]>([])
    
    const handleAddTodo = async () =>{
        const newTodo = await createTodo(newTodoTitle, newTodoDescription);
        setTodos([...todos,newTodo])
        setNewTodoTitle('');
        setNewTodoDescription('');
    }
    
   if (!visibility) return null;
   
   return (
        <>
            <div 
            className='grid justify-items-center text-white  gap-2 shadow-lg p-[2em] rounded-[12px]  w-[500px] h-auto bg-slate-700'>
                <h1 
                className='text-white font-bold text-[24px]'>Create a task</h1>
                <div 
                className='flex flex-col gap-2 font-semibold '>
                    <label>Title</label>
                    <input 
                    onChange={(e)=>{setNewTodoTitle(e.target.value)}}
                    type="text" 
                    name="title" 
                    id="title"
                    className='rounded-md' />

                    <label>Description</label>
                    <input 
                    onChange={(e)=>{setNewTodoDescription(e.target.value)}}
                    type="text" 
                    name="description" 
                    id="decription"
                    className='rounded-md' />



                </div>
                <button
                onClick={() => {handleAddTodo()}}
                className='bg-blue-700 max-w-[80px] p-2 rounded-md ' 
                type="submit">
                    Create
                </button>

            </div>
        </> || undefined
    )
}