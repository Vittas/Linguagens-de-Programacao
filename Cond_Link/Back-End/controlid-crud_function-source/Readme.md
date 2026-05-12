# Login GCloud
gcloud auth login

# Configure o projeto atual do gcloud para o projeto que contém sua função. Use o comando:
gcloud config set project earnest-cosmos-175020

# Listar projetos
gcloud projects list

# Deploy
gcloud functions deploy controlid-crud \
--entry-point BootUp \
--runtime go121 \
--trigger-http \
--allow-unauthenticated \
--gen2 \
--region=us-central1 \
--vpc-connector projects/earnest-cosmos-175020/locations/us-central1/connectors/vpc-function

# Run mode dev
- go build
- ./test-controlid-resynchronize

# Retorno de imagem ruim
Payload sended: {'statusCode': 200, 'message': '{"results":[{"user_id":43548,"success":false,"errors":[{"code":2,"message":"Face not detected"}]}]}', 'payload': {'results': [{'user_id': 43548, 'success': False, 'errors': [{'code': 2, 'message': 'Face not detected'}]}]}, '_idFront': 'd587f9e1-00ab-40be-a728-cd829814a294', 'idTask': 'task_orangepi01_101', 'date': '2024-07-04 00:45:09'}