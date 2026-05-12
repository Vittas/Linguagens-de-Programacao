import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [TextoReverse , SetTextoReverse] = useState("")

  const reverteTexto = () =>{
    let texto: string = (document.getElementById("texto") as HTMLInputElement).value
    let listaTexto: string[] = texto.split("")
    texto = ""
    for(let i in listaTexto.reverse()){
      texto += listaTexto[i]
    }
    SetTextoReverse(texto)
  }

  const traduzTexto = () => {
    let texto = (document.getElementById("texto") as HTMLInputElement).value

    let dict = {
      "eu": "I",
      "amo": "love",
      "doce": "candy"
    }

    console.log()

  }
  return (
    <>
      <div>
        <input type="text" name="texto" onSelect={reverteTexto} id="texto" />

        <div>{TextoReverse}</div>

      </div>
      
      <div>
        
        <input type="text" name="texto" onSelect={traduzTexto} id="texto" />

        <div>{Texto}</div>

      </div>



    </>
  )
}

export default App
