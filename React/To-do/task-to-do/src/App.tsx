import React, { useEffect, useState } from 'react';
import './App.css';
import { Layout } from './components/Layout';
// import { Card } from './components/Card';
import { CreateTaskModal } from './components/modals/CreateTaskModal';

import { fetchTodos, Todo } from './services/ApiTaskManager';


const App: React.FC = () => {
  const[todos, setTodos] = useState<Todo[]>([])

  useEffect(() => {
    const loadTodos = async () => {
      const todosData = await fetchTodos();
      setTodos(todosData);
    };
    loadTodos();
  }, [])

  const [visibility, setVisibility] = useState(false)
  return (
    <>
      <Layout>
          <button 
          onClick={()=>{ let switchVisibility = (visibility) ? false : true; setVisibility(switchVisibility)}} 
          className='place-self-center text-white border-black rounded-md bg-blue-700 p-1'>
            Create
          </button>
       
        <div 
        id='menu-task' 
        className='flex justify-center'>
          <CreateTaskModal visibility={visibility}/>

          <ul>
            {todos.map((todo)=> (
              <div key={todo.id}>
                <li >{todo.title}</li>
                <li>{todo.description}</li>
              </div>

            ))}

            
          </ul>
        </div>


      </Layout>


    </>
  );
}

export default App;
