import express from 'express';
import cors from 'cors';
import itensRouter from './routers/itens.router';

const PORT = process.env.PORT || 4001;

const HOSTNAME = process.env.HOSTNAME || "http://localhost";

const app = express();

//adicionando o json como meio padrão de comunicação

app.use(express.json());
app.use(cors())
app.use(express.urlencoded({extended: true}))

app.get('/', (req, res) => {
    res.send('Hello World');
})

// define que a origem é na porta 3000 e possibilita que ela receba dados de outras portas
app.use(cors({
    origin: ['http://localhost:3000']
}))

app.use('/api', itensRouter);

app.use((req, res) => {
    res.status(404);
})

app.listen(PORT, () => {
    console.log(`Servidor inicializado com sucesso ${HOSTNAME}:${PORT}/api/itens`);
})