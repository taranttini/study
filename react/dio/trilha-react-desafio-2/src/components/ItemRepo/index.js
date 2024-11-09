import React from 'react'
import { ItemContainer } from './styles';

function ItemRepo({repo, handleRemoveRepo}) {

  const handleRemove = () => {
    handleRemoveRepo(repo.id)
  }

  const handleLink = (url) => {
    window.location.href = `https://github.com/${url}`;
  }

  return (
    <ItemContainer >
        <h3>{repo.name}</h3>
        <p>{repo.full_name}</p>
        <a href="#" rel="noreferrer" onClick={()=>handleLink(repo)} >Ver repositório</a>&nbsp;&nbsp;&nbsp; 
        <a href="#" rel="noreferrer" className="remover" onClick={handleRemove}>Remover</a>
        <hr />
    </ItemContainer>
  )
}

export default ItemRepo;
