table: ticket
1	id	bigint(20) unsigned	NULL	NULL	NO	NULL			
2	tipo_remetente	varchar(50)	utf8	utf8_general_ci	NO	NULL			
3	id_remetente	bigint(20) unsigned	NULL	NULL	NO	NULL			
4	data_inicio	timestamp	NULL	NULL	NO	CURRENT_TIMESTAMP			
5	ultima_atualizacao	timestamp	NULL	NULL	NO	CURRENT_TIMESTAMP			
6	assunto	varchar(50)	utf8	utf8_general_ci	YES	NULL			
7	mensagem	mediumtext	utf8	utf8_general_ci	YES	NULL			
8	status	varchar(50)	utf8	utf8_general_ci	NO	NULL			
9	tipo_destinatario	varchar(50)	utf8	utf8_general_ci	NO	NULL			
10	id_destinatario	bigint(20) unsigned	NULL	NULL	NO	NULL			
11	tipo_ticket	varchar(60)	utf8	utf8_general_ci	YES	NULL			
12	tipo_servico	bigint(20) unsigned	NULL	NULL	NO	NULL			
13	campo_extra1	varchar(255)	utf8	utf8_general_ci	YES	NULL			
14	tipo_ajuda	varchar(50)	utf8	utf8_general_ci	YES	NULL			
15	remetente_nome	varchar(300)	utf8	utf8_general_ci	YES	NULL			
16	isLido	tinyint(1)	NULL	NULL	YES	0			
17	id_condominio	bigint(20) unsigned	NULL	NULL	NO	NULL			
18	id_responsavel	bigint(20) unsigned	NULL	NULL	YES	NULL			
19	arquivo_anexo	varchar(200)	utf8	utf8_general_ci	YES	NULL			
20	valor	double(10,2)	NULL	NULL	YES	NULL			
21	pago	tinyint(1)	NULL	NULL	YES	NULL			
22	data_vencimento_pagamento	timestamp	NULL	NULL	YES	NULL			
23	id_sinc	int(11)	NULL	NULL	YES	NULL			
24	dado_atualizado	tinyint(1)	NULL	NULL	YES	0			
25	data_atualizacao	timestamp	NULL	NULL	NO	CURRENT_TIMESTAMP			
26	tipo_servico_remetente	varchar(100)	utf8	utf8_general_ci	YES	NULL			
27	tipo_servico_destinatario	varchar(100)	utf8	utf8_general_ci	YES	NULL			



table: condomino
1	id	bigint(20) unsigned	NULL	NULL	NO	NULL			
2	id_antigo	bigint(20) unsigned	NULL	NULL	YES	NULL			
3	nome	varchar(255)	latin1	latin1_swedish_ci	NO				
4	rg	varchar(20)	latin1	latin1_swedish_ci	NO				
5	ie	varchar(15)	latin1	latin1_swedish_ci	YES	NULL			
6	data_nascimento	varchar(10)	latin1	latin1_swedish_ci	NO	NULL			
7	email	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
8	email_ativo	tinyint(4)	NULL	NULL	NO	1			
9	email_valido	tinyint(4)	NULL	NULL	NO	1			
10	unidade	varchar(2050)	latin1	latin1_swedish_ci	YES	NULL			
11	tipo_morador	int(1)	NULL	NULL	YES	NULL			
12	numero_lote	varchar(255)	latin1	latin1_swedish_ci	YES				
13	endereco	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
14	bairro	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
15	cidade	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
16	estado	varchar(2)	latin1	latin1_swedish_ci	YES	NULL			
17	placa_carro	varchar(10)	latin1	latin1_swedish_ci	YES				
18	username	varchar(255)	latin1	latin1_swedish_ci	YES				
19	password	varchar(255)	latin1	latin1_swedish_ci	YES				
20	id_superior	bigint(20) unsigned	NULL	NULL	YES	NULL			
21	status	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
22	tipoDependente	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
23	campo_extra3	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
24	ativo	int(1)	NULL	NULL	NO	1			
25	cep	varchar(10)	latin1	latin1_swedish_ci	YES	NULL			
26	cpf_cnpj	varchar(20)	latin1	latin1_swedish_ci	YES	NULL			
27	cnpj	varchar(18)	latin1	latin1_swedish_ci	YES	NULL			
28	provedorInternet	varchar(50)	latin1	latin1_swedish_ci	YES	NULL			
29	telefone	varchar(30)	latin1	latin1_swedish_ci	NO	NULL			
30	ddd_telefone	int(3)	NULL	NULL	NO	NULL			
31	operadora_telefone	varchar(50)	latin1	latin1_swedish_ci	YES	NULL			
32	banda_internet	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
33	inadimplente	int(1)	NULL	NULL	YES	0			
34	numeroVagaGaragem	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
35	telefoneCelular	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
36	ddd_telefoneCelular	int(3)	NULL	NULL	YES	0			
37	operadora_telefoneCelular	varchar(50)	latin1	latin1_swedish_ci	YES	NULL			
38	complemento	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
39	apartamento	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
40	numero	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
41	bloco	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
42	tipoPessoa	char(1)	latin1	latin1_swedish_ci	YES	NULL			
43	telefoneTrabalho	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
44	ddd_telefoneTrabalho	int(3)	NULL	NULL	YES	0			
45	operadora_telefoneTrabalho	varchar(50)	latin1	latin1_swedish_ci	YES	NULL			
46	principal	tinyint(1)	NULL	NULL	YES	NULL			
47	id_condominio	bigint(20) unsigned	NULL	NULL	NO	NULL			
48	foto	varchar(30)	latin1	latin1_swedish_ci	YES				
49	telefone2	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
50	ddd_telefone2	int(3)	NULL	NULL	YES	0			
51	operadora_telefone2	varchar(50)	latin1	latin1_swedish_ci	YES	NULL			
52	telefoneCelular2	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
53	ddd_telefoneCelular2	int(3)	NULL	NULL	YES	0			
54	operadora_telefoneCelular2	varchar(50)	latin1	latin1_swedish_ci	YES	NULL			
55	portadorDeficiencia	int(1)	NULL	NULL	YES	0			
56	sexo	varchar(1)	latin1	latin1_swedish_ci	YES	NULL			
57	numeroControleRemoto	varchar(50)	latin1	latin1_swedish_ci	YES	NULL			
58	id_sinc	int(11)	NULL	NULL	YES	NULL			
59	dado_atualizado	tinyint(1)	NULL	NULL	YES	0			
60	data_atualizacao	timestamp	NULL	NULL	NO	CURRENT_TIMESTAMP	on update CURRENT_TIMESTAMP		
61	endereco2	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
62	bairro2	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
63	cidade2	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
64	estado2	varchar(2)	latin1	latin1_swedish_ci	YES	NULL			
65	cep2	varchar(10)	latin1	latin1_swedish_ci	YES	NULL			
66	complemento2	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
67	apartamento2	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
68	numero2	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
69	bloco2	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
70	contatoemergencia	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
71	telefoneemergencia	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
72	ddd_telefoneEmergencia	int(3)	NULL	NULL	YES	0			
73	operadora_telefoneEmergencia	varchar(50)	latin1	latin1_swedish_ci	YES	NULL			
74	numeroVagaGaragem2	varchar(30)	latin1	latin1_swedish_ci	YES	NULL			
75	correspondencia_igual_unidade	bigint(20) unsigned	NULL	NULL	YES	NULL			
76	gtalk	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
77	msn	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
78	texto_notificacao	text	latin1	latin1_swedish_ci	YES	NULL			
79	representa_condominio	int(1)	NULL	NULL	YES	0			
80	pode_autorizar_na_portaria	tinyint(1)	NULL	NULL	YES	NULL			
81	status_email	varchar(2)	latin1	latin1_swedish_ci	YES	0			
82	historico_condomino	text	latin1	latin1_swedish_ci	YES	NULL			
83	data_cadastro	datetime	NULL	NULL	YES	NULL			
84	chave_pag_condominio	varchar(45)	latin1	latin1_swedish_ci	YES	NULL			
85	data_expiracao_pre_aprovado	timestamp	NULL	NULL	YES	NULL			
86	pode_gerenciar_visitantes	tinyint(1)	NULL	NULL	YES	0			
87	pessoa_id	bigint(20)	NULL	NULL	YES	NULL			
88	id_termo	bigint(20)	NULL	NULL	YES	NULL			
89	version	bigint(20)	NULL	NULL	YES	NULL			
90	status_termo	int(1)	NULL	NULL	YES	0			
91	ip_usuario	varchar(255)	latin1	latin1_swedish_ci	YES	NULL			
92	date_created	datetime	NULL	NULL	YES	NULL			
93	ip	varchar(255)	latin1	latin1_swedish_ci	YES	NULL


table: condomino_unidade
1	id	bigint(20)	NULL	NULL	NO	NULL			
2	id_condomino	bigint(20)	NULL	NULL	NO	NULL			
3	id_unidade	bigint(20)	NULL	NULL	NO	NULL			
4	tipo_morador	varchar(45)	latin1	latin1_swedish_ci	NO	NULL			
5	principal	tinyint(1)	NULL	NULL	NO	NULL			
6	data_atualizacao	timestamp	NULL	NULL	NO	CURRENT_TIMESTAMP			
7	dado_atualizado	int(11)	NULL	NULL	YES	0			
8	id_sinc	int(11)	NULL	NULL	YES	NULL			
9	data_entrada_inquilino	date	NULL	NULL	YES	NULL			
10	data_saida_inquilino	date	NULL	NULL	YES	NULL			


table: unidade
1	id	bigint(20)	NULL	NULL	NO	NULL			
2	numero	varchar(15)	latin1	latin1_swedish_ci	YES	NULL			
3	fracao_ideal	double	NULL	NULL	YES	NULL			
4	area_adquirida	double	NULL	NULL	YES	NULL			
5	tipo_unidade	varchar(1)	latin1	latin1_swedish_ci	YES	NULL			
6	rua	varchar(100)	latin1	latin1_swedish_ci	YES	NULL			
7	numero_residencia	int(11)	NULL	NULL	YES	NULL			
8	complemento	varchar(100)	latin1	latin1_swedish_ci	YES	NULL			
9	cep	varchar(45)	latin1	latin1_swedish_ci	YES	NULL			
10	ramal	varchar(10)	latin1	latin1_swedish_ci	YES	NULL			
11	id_condominio	bigint(20) unsigned	NULL	NULL	NO	NULL			
12	data_atualizacao	timestamp	NULL	NULL	NO	CURRENT_TIMESTAMP			
13	dado_atualizado	int(11)	NULL	NULL	YES	0			
14	id_sinc	int(11)	NULL	NULL	YES	NULL			
15	apartamento	int(11)	NULL	NULL	YES	NULL			
16	bloco	varchar(45)	latin1	latin1_swedish_ci	YES	NULL			
17	codigo_consumidor_agua	varchar(45)	latin1	latin1_swedish_ci	YES	NULL			
18	usuario_proprietario	varchar(45)	latin1	latin1_swedish_ci	YES	NULL			
19	senha_proprietario	varchar(45)	latin1	latin1_swedish_ci	YES	NULL			
20	usuario_inquilino	varchar(45)	latin1	latin1_swedish_ci	YES	NULL			
21	senha_inquilino	varchar(45)	latin1	latin1_swedish_ci	YES	NULL			
22	id_unidade	int(11)	NULL	NULL	NO	0			


note: id_destinatario é o id do condomino. na tabela unidade o 'numero' é o 'unidade' da tabela condomino.
