import { useForm } from "react-hook-form";
import Button from "../../components/Button";
import Input from "../../components/Input";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";

import { Container, LoginContainer, Column, Spacing, Title } from "./styles";
import { defaultValues, IFormLogin } from "./types";

import { useState } from 'react';

const schema = yup
  .object({
    email: yup.string().email("E-mail inválido").required("Campo obrigatório"),
    password: yup
      .string()
      .min(6, "No minimo 6 caracteres")
      .required("Campo obrigatório"),
  })
  .required();

const Login = () => {
  const [logged, setLogged] = useState(false);


  const {
    control,
    formState: { errors, isValid },
  } = useForm<IFormLogin>({
    resolver: yupResolver(schema),
    mode: "onBlur",
    defaultValues,
    reValidateMode: "onChange",
  });


  const onClick = () => {
    if (!logged && isValid) {
      setLogged(true)
    } else {
      setLogged(false)
    }
  }

  return (
    <Container>
      <LoginContainer>
        {
          !logged ?
            <>
              <Column>
                <Title>Login</Title>
                <Spacing />
                <Input
                  name="email"
                  placeholder="Email"
                  control={control}
                  errorMessage={errors?.email?.message}
                />
                <Spacing />
                <Input
                  name="password"
                  type="password"
                  placeholder="Senha"
                  control={control}
                  errorMessage={errors?.password?.message}
                />
                <Spacing />
                <Button title="Entrar" onClick={onClick} />
              </Column>
            </>
            :
            <>
              <Column>
              <p>logged</p>
              <Spacing />
              <Button title="Sair" onClick={onClick} />
              </Column>
            </>
        }
      </LoginContainer>
    </Container>
  );
};

export default Login;
