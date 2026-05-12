import { Box, Button, Center, Flex, Heading, Input } from '@chakra-ui/react';
import './App.css';
import { Provider } from './components/ui/provider';
import { login } from './services/login';

function App() {
  return (
    <Provider>
      <Flex direction={'column'} bg={'white'} h={'100vh'} w={'100%'} padding={'2em'}>
        <Center>
          <Flex direction={'column'} spaceY={'1em'} alignItems={'center'} bg={'green.500'} w={'25em'} padding={'2em'} rounded={'md'}>
            
            <Heading textStyle={'4xl'} fontWeight={'bold'}>Login</Heading>
            
            <Input padding={'1em'} placeholder='Email' bg={'white'} variant={'subtle'} />
            <Input padding={'1em'} placeholder='Senha' bg={'white'} variant={'subtle'} />

            <Button fontWeight={'semibold'} onClick={login} bg={'green.700'} padding={'1em'}>Entrar</Button>
            
          </Flex>
        </Center>
      </Flex>
    </Provider>
  );
}

export default App;
