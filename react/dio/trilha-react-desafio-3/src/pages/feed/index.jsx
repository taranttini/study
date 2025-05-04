import React from 'react'
import { Card } from '../../components/Card';
import { UserInfo } from '../../components/UserInfo';

import { Header } from '../../components/Header';

import { Container, Column, Title, TitleHighlight } from './styles';

const Feed = () => {
  return (
    <>
        <Header autenticado={true}/>
        <Container>
            <Column flex={3}>
                <Title>Feed</Title>
                <Card imageIcon={"1"} title={"Codigo 01"} message={"xpto lorem sprum"} tags={"#c #microsoft"} userName={"User 1"} time={"8 min"} />
                <Card imageIcon={"2"} title={"Codigo 02"} message={"xpto lorem sprum"} tags={"#js #node"} userName={"User 2"} time={"16 min"} />
                <Card imageIcon={"3"} title={"Codigo 03"} message={"xpto lorem sprum"} tags={"#go #google"} userName={"User 3"} time={"24 min"} />
                <Card imageIcon={"4"} title={"Codigo 04"} message={"xpto lorem sprum"} tags={"#java #oracle"} userName={"User 4"} time={"32 min"} />
                <Card imageIcon={"5"} title={"Codigo 05"} message={"xpto lorem sprum"} tags={"#sql #ibm"} userName={"User 5"} time={"40 min"} />
                <Card imageIcon={"6"} title={"Codigo 06"} message={"xpto lorem sprum"} tags={"#nosql #mongodb"} userName={"User 6"} time={"48 min"} />
                <Card imageIcon={"7"} title={"Codigo 07"} message={"xpto lorem sprum"} tags={"#vb #microsoft"} userName={"User 7"} time={"56 min"} />
                <Card imageIcon={"8"} title={"Codigo 08"} message={"xpto lorem sprum"} tags={"#rust #mozilla"} userName={"User 8"} time={"64 min"} />
                <Card imageIcon={"9"} title={"Codigo 09"} message={"xpto lorem sprum"} tags={"#delphi #borland"} userName={"User 9"} time={"72 min"} />
                <Card imageIcon={"10"} title={"Codigo 10"} message={"xpto lorem sprum"} tags={"#carbon #google"} userName={"User 10"} time={"80 min"} />
                
            </Column>
            <Column flex={1}>
              <TitleHighlight> # RANKING 5 TOP DA SEMANA </TitleHighlight>
                <UserInfo nome="User One" image="https://avatars.githubusercontent.com/u/1?v=" percentual={20}/>
                <UserInfo nome="User Two" image="https://avatars.githubusercontent.com/u/2?v=" percentual={40}/>
                <UserInfo nome="User Three" image="https://avatars.githubusercontent.com/u/3?v=" percentual={60}/>
                <UserInfo nome="User Four" image="https://avatars.githubusercontent.com/u/4?v=" percentual={80}/>
                <UserInfo nome="User Five" image="https://avatars.githubusercontent.com/u/5?v=" percentual={99}/>
            </Column>
        </Container>
    </>
  )
}

export { Feed }; 
