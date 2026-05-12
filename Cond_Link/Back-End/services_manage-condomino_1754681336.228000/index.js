const mysql = require('mysql2/promise');
const jwt = require('jsonwebtoken');

const dbConfig = {
  host: 'mysql5.condlink.com.br',
  user: 'joao',
  password: 'Ph=f-rF9A{8285&+',
  database: 'multicondlink',
  port: 3306,
  connectionLimit: 10,
  supportBigNumbers: true,
  bigNumberStrings: true,
};

const pool = mysql.createPool(dbConfig);
const JWT_SECRET = process.env.JWT_SECRET;

const authenticate = (req, res, next) => {
  const authHeader = req.headers.authorization;
  if (!authHeader || !authHeader.startsWith('Bearer ')) {
    return res.status(401).json({ message: 'Unauthorized: Missing or invalid token' });
  }
  const token = authHeader.split(' ')[1];
  try {
    const decoded = jwt.verify(token, JWT_SECRET);
    req.user = decoded;
    next();
  } catch (error) {
    return res.status(401).json({ message: 'Unauthorized: Invalid token' });
  }
};

const handleRequest = async (req, res) => {
  const { method } = req;
  const { id_condominio } = req.user;
  const condominoId = req.path.split('/').pop();

  try {
    const connection = await pool.getConnection();

    switch (method) {
      case 'POST':
        await connection.beginTransaction();
        try {
          // Get the next pessoa_id
          const [lastPessoa] = await connection.query('SELECT MAX(id) as maxId FROM `pessoa`');
          const nextPessoaId = lastPessoa[0].maxId ? (BigInt(lastPessoa[0].maxId) + 1n) : 1n;

          // Create new pessoa
          const newPessoa = {
            id: nextPessoaId,
            nome: req.body.nome,
            email: req.body.email,
            cpf: req.body.cpf_cnpj,
            rg: req.body.rg,
            date_created: new Date(Date.now() - 3 * 60 * 60 * 1000),
            last_updated: new Date(Date.now() - 3 * 60 * 60 * 1000),
          };
          await connection.query('INSERT INTO `pessoa` SET ?', newPessoa);

          // Get the next condomino_id
          const [lastCondomino] = await connection.query('SELECT MAX(id) as maxId FROM `condomino`');
          const nextCondominoId = lastCondomino[0].maxId ? (BigInt(lastCondomino[0].maxId) + 1n) : 1n;
          
          const newCondomino = { 
              ...req.body, 
              id: nextCondominoId,
              pessoa_id: nextPessoaId,
              id_condominio,
              data_cadastro: new Date(Date.now() - 3 * 60 * 60 * 1000),
              date_created: new Date(Date.now() - 3 * 60 * 60 * 1000),
          };

          await connection.query('INSERT INTO `condomino` SET ?', newCondomino);
          
          await connection.commit();
          
          const responseCondomino = {
            ...newCondomino,
            id: newCondomino.id.toString(),
            pessoa_id: newCondomino.pessoa_id.toString(),
          }

          delete responseCondomino.password;
          res.status(201).json(responseCondomino);

        } catch (err) {
          await connection.rollback();
          console.error('Transaction Error:', err);
          res.status(500).json({ message: 'Failed to create condomino.' });
        }
        break;

      case 'GET':
        if (condominoId && condominoId !== 'manage-condomino') {
          const [rows] = await connection.query('SELECT * FROM `condomino` WHERE `id` = ? AND `id_condominio` = ?', [condominoId, id_condominio]);
          if (rows.length === 0) {
            return res.status(404).json({ message: 'Condomino not found' });
          }
          const condomino = rows[0];
          delete condomino.password;

          const [parentescoRows] = await connection.query('SELECT * FROM `condomino` WHERE `id_superior` = ? AND `id_condominio` = ?', [condominoId, id_condominio]);
          const parentesco = parentescoRows.map(p => {
            delete p.password;
            return p;
          });

          condomino.parentesco = parentesco;

          const [animaisRows] = await connection.query('SELECT * FROM `animais_condominos` WHERE `id_condomino` = ?', [condominoId]);
          condomino.animais = animaisRows;

          const [veiculosRows] = await connection.query(
            'SELECT v.*, c.cor, m.nome as montadora FROM `veiculos_condominos` v LEFT JOIN `cor_veiculos_condominos` c ON v.id_cor = c.id LEFT JOIN `montadoras_veiculos_condominos` m ON v.id_montadora = m.id WHERE v.id_condomino = ?',
            [condominoId]
          );
          condomino.veiculos = veiculosRows;

          const [unidadesRows] = await connection.query(
            'SELECT u.* FROM `unidade` u JOIN `condomino_unidade` cu ON u.id = cu.id_unidade WHERE cu.id_condomino = ?',
            [condominoId]
          );
          condomino.unidades = unidadesRows;

          res.status(200).json(condomino);
        } else {
            const { page = 1, limit = 10, nome, cpf_cnpj, rg } = req.query;
            const offset = (page - 1) * limit;
            let whereClauses = ['`id_condominio` = ?'];
            let params = [id_condominio];

            if (nome) {
              whereClauses.push('`nome` LIKE ?');
              params.push(`%${nome}%`);
            }
            if (cpf_cnpj) {
              whereClauses.push('`cpf_cnpj` = ?');
              params.push(cpf_cnpj);
            }
            if (rg) {
              whereClauses.push('`rg` = ?');
              params.push(rg);
            }

            const whereSql = whereClauses.join(' AND ');
            const [rows] = await connection.query(
              `SELECT * FROM \`condomino\` WHERE ${whereSql} LIMIT ? OFFSET ?`,
              [...params, parseInt(limit), parseInt(offset)]
            );
            const [totalRows] = await connection.query(`SELECT COUNT(*) as count FROM \`condomino\` WHERE ${whereSql}`, params);
            
            const condominos = rows.map(condomino => {
                delete condomino.password;
                return condomino;
            });

            res.status(200).json({
              total: totalRows[0].count,
              page: parseInt(page),
              limit: parseInt(limit),
              data: condominos,
            });
        }
        break;

      case 'PUT':
        if (!condominoId) return res.status(400).json({ message: 'Condomino ID is required' });
        
        await connection.beginTransaction();
        try {
            const { body } = req;
            const [condominoRows] = await connection.query('SELECT pessoa_id FROM `condomino` WHERE `id` = ? AND `id_condominio` = ?', [condominoId, id_condominio]);

            if (condominoRows.length === 0) {
                await connection.rollback();
                return res.status(404).json({ message: 'Condomino not found' });
            }
            const { pessoa_id } = condominoRows[0];

            const updatedPessoa = {
                nome: body.nome,
                email: body.email,
                cpf: body.cpf_cnpj,
                rg: body.rg,
                last_updated: new Date(Date.now() - 3 * 60 * 60 * 1000),
            };
            await connection.query('UPDATE `pessoa` SET ? WHERE `id` = ?', [updatedPessoa, pessoa_id]);

            const updatedCondomino = { ...body };
            delete updatedCondomino.id;
            delete updatedCondomino.id_condominio;
            delete updatedCondomino.pessoa_id; 
            
            await connection.query('UPDATE `condomino` SET ? WHERE `id` = ? AND `id_condominio` = ?', [updatedCondomino, condominoId, id_condominio]);
            
            await connection.commit();
            
            res.status(200).json({ id: condominoId, ...body });

        } catch (err) {
            await connection.rollback();
            console.error('Transaction Error:', err);
            res.status(500).json({ message: 'Failed to update condomino.' });
        }
        break;

      case 'DELETE':
        if (!condominoId) return res.status(400).json({ message: 'Condomino ID is required' });
        await connection.query('DELETE FROM `condomino` WHERE `id` = ? AND `id_condominio` = ?', [condominoId, id_condominio]);
        res.status(204).send();
        break;

      default:
        res.setHeader('Allow', ['GET', 'POST', 'PUT', 'DELETE']);
        res.status(405).end(`Method ${method} Not Allowed`);
    }
    connection.release();
  } catch (error) {
    console.error('Database Error:', error);
    res.status(500).json({ message: 'Internal Server Error' });
  }
};

exports.manageCondomino = (req, res) => {
  res.set('Access-Control-Allow-Origin', '*');

  if (req.method === 'OPTIONS') {
    // Send response to preflight requests
    res.set('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE');
    res.set('Access-Control-Allow-Headers', 'Content-Type, Authorization');
    res.set('Access-Control-Max-Age', '3600');
    res.status(204).send('');
    return;
  }

  authenticate(req, res, () => handleRequest(req, res));
};
