
import { useState } from 'react';
import gitLogo from '../assets/github.png'
import Input from '../components/Input';
import Button from '../components/Button';
import ItemRepo from '../components/ItemRepo';
import { api } from '../services/api';

import { Container, Logo } from './styles';

function App() {

  const [currentRepo, setCurrentRepo] = useState('');
  const [repos, setRepos] = useState([]);

  const handleSearchRepo = async () => {

    await api.get(`users/${currentRepo}/repos`).catch(q => []).then(q => { 
      setRepos(q.data)
    })

  }

  const handleRemoveRepo = (item) => {
    console.log('Removendo registro', item.id);
    let list = repos.filter(q => (q.id !== item.id))
    setRepos(list)
  }

  return (
    <>
      <Logo>
        <img src={gitLogo} width={72} height={72} alt="github logo" />
        <p>Digite o nome do usuário e liste seus repos </p>
      </Logo>
      <Input value={currentRepo} onChange={(e) => setCurrentRepo(e.target.value)} />
      <Button onClick={handleSearchRepo} />
      <Container >
        {
          (!repos || repos.length === 0) 
            ? <p>Item nao localizado</p>
            : repos.map((item, idx) => (
              <div key={idx}>
                <p>{item.id} - {item.full_name}</p>
                <ItemRepo handleRemoveRepo={() => { handleRemoveRepo(item) }} repo={item.full_name} />
              </div>
            ))
          }

      </Container>
    </>
  );
}

export default App;
