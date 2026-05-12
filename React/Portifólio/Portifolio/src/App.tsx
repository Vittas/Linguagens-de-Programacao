import { useState } from 'react'
import ft_perfil from '/home/vittas/Documentos/Pasta_para_Todo_Governar/Linguagens_de_Programacao/React Js/Portifólio/Portifolio/src/assets/file.jpeg'
import './App.css'

function App() {
  const [color,setColor] = useState("#000000");
  const [color2,setColor2] = useState("#FFFFFF");

  function changeColor1(){
    const hex = "ABCDEF0123456789";
    const charLength = hex.length;
    let result = '';
    for ( let i = 0; i < 6; i ++) {
      result += hex.charAt(Math.floor(Math.random() * charLength));
    }
    // const randomColor =  Math.floor(Math.random()*16777215).toString(16);
    setColor("#" + result);
    console.log(result);

  }

  function changeColor2(){
    const hex = "ABCDEF0123456789";
    const charLength = hex.length;
    let result = '';
    for ( let i = 0; i < 6; i ++) {
      result += hex.charAt(Math.floor(Math.random() * charLength));
    }
    // const randomColor =  Math.floor(Math.random()*16777215).toString(16);
    setColor2("#" + result);
    console.log(result);

  }

  return (
    <>
      <h1 style={{color: color2}}>Portifólio</h1>
      <div className="card">
        <div id='perfil-bg' style={{backgroundColor: color}}>
          {/* <img src={ft_perfil} id='ft-perfil'/> */}
        </div>
        <h2>Vitor Gabriel</h2>
        <button  onClick={changeColor1}>Change bg color </button>
        <br />
        <br />
        <button  onClick={changeColor2}>Change text color </button>

      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App