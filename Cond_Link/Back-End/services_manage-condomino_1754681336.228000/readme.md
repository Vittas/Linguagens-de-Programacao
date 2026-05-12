# Manage Condomino Function

This document provides instructions on how to deploy and test the `manage-condomino` Google Cloud Function.

## 1. Deploy the Function

To deploy the function, use the following `gcloud` command. Make sure you are in the `functions/manage-condomino` directory before running the command.

### Deploy Command

```bash
gcloud functions deploy manage-condomino \
  --entry-point manageCondomino \
  --runtime nodejs20 \
  --trigger-http \
  --allow-unauthenticated \
  --gen2 \
  --region=us-central1 \
  --source . \
  --vpc-connector projects/earnest-cosmos-175020/locations/us-central1/connectors/vpc-function \
  --set-env-vars JWT_SECRET=2j43h234k2@#$kdnjas.asd@#dasd
```

**Note:** Replace `YOUR_JWT_SECRET` with the same secret key used for the `login-admin` function.

## 2. Test the Function with cURL

After deploying, you can test the function using `cURL`. First, get a valid JWT token by logging in with the `login-admin` function.

### cURL Commands

Replace `[YOUR_TOKEN]` with the JWT token obtained from the login.
Replace `[CONDOMINO_ID]` with the ID of the condomino you want to manage.

**Create a new Condomino:**
```bash
curl -X POST "https://us-central1-earnest-cosmos-175020.cloudfunctions.net/manage-condomino" \
-H "Authorization: Bearer [YOUR_TOKEN]" \
-H "Content-Type: application/json" \
-d '{
  "id_antigo": null,
  "nome": "João da Silva",
  "rg": "12.345.678-9",
  "ie": null,
  "data_nascimento": "1990-01-15",
  "email": "joao.silva@example.com",
  "email_ativo": 1,
  "email_valido": 1,
  "unidade": "Apto 101",
  "tipo_morador": 1,
  "numero_lote": "Lote 25",
  "endereco": "Rua das Flores, 123",
  "bairro": "Centro",
  "cidade": "São Paulo",
  "estado": "SP",
  "placa_carro": "ABC-1234",
  "username": "joao.silva",
  "password": "securepassword123",
  "id_superior": null,
  "status": "ativo",
  "tipoDependente": null,
  "campo_extra3": null,
  "ativo": 1,
  "cep": "01000-000",
  "cpf_cnpj": "123.456.789-00",
  "cnpj": null,
  "provedorInternet": "Vivo Fibra",
  "telefone": "1122223333",
  "ddd_telefone": 11,
  "operadora_telefone": "Vivo",
  "banda_internet": "100 Mega",
  "inadimplente": 0,
  "numeroVagaGaragem": "G1-10",
  "telefoneCelular": "11999998888",
  "ddd_telefoneCelular": 11,
  "operadora_telefoneCelular": "Claro",
  "complemento": "Bloco A",
  "apartamento": "101",
  "numero": "123",
  "bloco": "A",
  "tipoPessoa": "F",
  "telefoneTrabalho": "1144445555",
  "ddd_telefoneTrabalho": 11,
  "operadora_telefoneTrabalho": "Tim",
  "principal": 1,
  "foto": "foto_joao.jpg",
  "telefone2": null,
  "ddd_telefone2": null,
  "operadora_telefone2": null,
  "telefoneCelular2": null,
  "ddd_telefoneCelular2": null,
  "operadora_telefoneCelular2": null,
  "portadorDeficiencia": 0,
  "sexo": "M",
  "numeroControleRemoto": "C-55",
  "id_sinc": null,
  "dado_atualizado": 0,
  "data_atualizacao": "2025-07-02T12:00:00Z",
  "endereco2": null,
  "bairro2": null,
  "cidade2": null,
  "estado2": null,
  "cep2": null,
  "complemento2": null,
  "apartamento2": null,
  "numero2": null,
  "bloco2": null,
  "contatoemergencia": "Maria Silva",
  "telefoneemergencia": "11988887777",
  "ddd_telefoneEmergencia": 11,
  "operadora_telefoneEmergencia": "Oi",
  "numeroVagaGaragem2": null,
  "correspondencia_igual_unidade": null,
  "gtalk": null,
  "msn": null,
  "texto_notificacao": null,
  "representa_condominio": 0,
  "pode_autorizar_na_portaria": 1,
  "status_email": "1",
  "historico_condomino": null,
  "chave_pag_condominio": null,
  "data_expiracao_pre_aprovado": null,
  "pode_gerenciar_visitantes": 1,
  "id_termo": null,
  "version": 1,
  "status_termo": 1,
  "ip_usuario": "192.168.1.100",
  "ip": "192.168.1.100"
}'
```

**List Condominos (with pagination and search):**
```bash
# List first page, 5 items per page
curl -X GET "https://us-central1-earnest-cosmos-175020.cloudfunctions.net/manage-condomino?page=1&limit=5" \
-H "Authorization: Bearer [YOUR_TOKEN]"

# Search by name on the second page
curl -X GET "https://us-central1-earnest-cosmos-175020.cloudfunctions.net/manage-condomino?nome=Silva&page=2&limit=10" \
-H "Authorization: Bearer [YOUR_TOKEN]"

# Search by CPF/CNPJ
curl -X GET "https://us-central1-earnest-cosmos-175020.cloudfunctions.net/manage-condomino?cpf_cnpj=123.456.789-00" \
-H "Authorization: Bearer [YOUR_TOKEN]"

# Search by RG
curl -X GET "https://us-central1-earnest-cosmos-175020.cloudfunctions.net/manage-condomino?rg=12.345.678-9" \
-H "Authorization: Bearer [YOUR_TOKEN]"
```

**Get a specific Condomino by ID:**
```bash
curl -X GET "https://us-central1-earnest-cosmos-175020.cloudfunctions.net/manage-condomino/[CONDOMINO_ID]" \
-H "Authorization: Bearer [YOUR_TOKEN]"
```

**Update a Condomino:**
```bash
curl -X PUT "https://us-central1-earnest-cosmos-175020.cloudfunctions.net/manage-condomino/[CONDOMINO_ID]" \
-H "Authorization: Bearer [YOUR_TOKEN]" \
-H "Content-Type: application/json" \
-d '{
  "nome": "Condomino Atualizado"
}'
```

**Delete a Condomino:**
```bash
curl -X DELETE "https://us-central1-earnest-cosmos-175020.cloudfunctions.net/manage-condomino/[CONDOMINO_ID]" \
-H "Authorization: Bearer [YOUR_TOKEN]"
