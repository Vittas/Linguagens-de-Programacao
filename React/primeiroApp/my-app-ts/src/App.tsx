import React from 'react';
import logo from './logo.svg';
import './App.css';
import {Card} from './Components/Card'
import { Layout } from './Components/Layout';




function App() {
  return (
    <>
      <Layout>
        <div>
          <h1>Pagina com cards</h1>
          <Card 
            id={1} 
            paragraph='History' 
            details='Brazil'
          />
          <Card
          id={2} 
          paragraph='Math' 
          details='Calculus'
          />
          <Card
          id={3} 
          paragraph='Programming' 
          details='Typescript'
          />


        </div>
      </Layout>
    </>
  );
}

export default App;
