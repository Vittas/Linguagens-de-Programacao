import express from 'express';
import Item from '../models/item';
import { error } from 'console';

const itensRouter = express.Router();

const itens: Item[] = []

//cria um item
itensRouter.post('/itens', (req, res) => {
    const id: number = itens.length+1
    const nome: string = req.body.nome
    const descricao: string = req.body.descricao

    if(nome && nome != "" && descricao){
        const item: Item = {
            id: id,
            nome: nome,
            descricao: descricao
        }
        itens.push(item)
        res.json(item);

    }
    

})

itensRouter.get('/itens', (req, res) => {
    res.json(itens)
})

//Procura por um produto especifico
itensRouter.get('/itens/:id', (req, res) => {
    const id: number = +req.params.id; // Converte o parâmetro para número
    
    const item: Item = itens.find(item => item.id === id)!; // Busca pelo ID no array
    if (item) {
        res.json(item);
    } else {
        res.status(404).json({ error: 'Item não encontrado' });
    }

});


//atualiza um item especifico
itensRouter.put('/itens/:id', (req, res) => {
    const id : number = +req.params.id;
    const item: Item | undefined = itens.find(item => item.id === id)
    const nome: string = req.body.nome;

    if(item){
        if(nome && nome != ""){
            item!.nome = nome;
            res.json(item);
        }
        else{
            res.status(400).json({error: 'O nome inserido não é válido!'})
        }
    }
    else{
        res.status(404).json({error: 'Item não encontado'})

    }


})

// deleta todos os itens
itensRouter.delete('/itens', (req, res) => {
    let numberItens: number = itens.length
    for(let i: number = 0; i < numberItens; i++){
        itens.pop()
    }
    res.json(itens)
    
})

//deleta um item especifico
itensRouter.delete('/itens/:id', (req, res) => {
    const id : number = +req.params.id;
    const item: Item | undefined = itens.find(item => item.id === id)
    let index = itens.indexOf(item!)
    if(item){
        itens.splice(index, 1)
        res.status(204).json(itens)
    }
    else{
        res.status(404).json({error: 'Item não encontrado!'})
    }
})

export default itensRouter;