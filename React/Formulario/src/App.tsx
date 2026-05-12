import { useState } from 'react'
import './App.css'
import * as React from 'react';




function App() {

  const [nomes,setNome]= React.useState([] as React.ReactElement[]);
  function listarPessoas(){
    let lista_pessoas : React.ReactElement[] = []
    let nome = (document.getElementById("nome") as HTMLInputElement).value ?? '';
    if (nomes.length !== 0){
      for(let i : number = 0; i < nomes.length ; i ++){
        lista_pessoas.push(<li>{nomes[i]}</li>)
  
      }
      lista_pessoas.push(<li>{nome}</li>)
      // console.log(nomes)
      setNome(lista_pessoas)
    }
    else{
      lista_pessoas.push(<li>{nome}</li>)
      // console.log(nomes)
      setNome(lista_pessoas)
    }


  }

  return (
    <>
      <div>
        <form>
          <label>topico</label>
          <input type="text" name="topico" id="" />
          <label>nomes</label>
          <input type="text" name="nome" id="nome" />
          <br />
          <button type="button" onClick={() => listarPessoas()}>clickar</button>
        </form>
      </div>
      <div>
        <ul>
          
          {nomes}
        </ul>
      </div>
    </>
  )
}

export default App
