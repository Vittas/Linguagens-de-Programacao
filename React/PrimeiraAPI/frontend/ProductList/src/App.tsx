import { ProductApiProvider } from './services/ProductApiProvider'
import './App.css'
import { CreateProductButton } from './components/CreateProductButton'
import { MainPage } from './components/MainPage'
function App() {
  
  
  return (
    <>
      <ProductApiProvider>
        <MainPage/>
      </ProductApiProvider>
    </>
  )
}

export default App
