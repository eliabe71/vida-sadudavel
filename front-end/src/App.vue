<template>
  <div id="app">
    <div id="form">
      <div class="text-center">
        <h1 class="title green-text">Login</h1>
      </div>
      <div id="logo" class="logo-img">
        <img src="./assets/images/logo.png"/>
      </div>
      <form id="frm">
        <div id="input-group-1" class="form-group">
          <label class="mt-3 green-text" for= "input-1">Email ou nome de usuário</label>
          <input id="input-1" class="form-control" v-model="user.login" type="text"
            placeholder="Digite o email ou nome de usuário"
            required
          />
        </div>

        <div id="input-group-2" class="form-group mt-2">
          <label class="mt-2 green-text" for= "input-1">Senha</label>
          <input id="input-2" class="form-control" v-model="user.password" type="password"
            placeholder="Digite sua senha"
            required
          />
        </div>

        <div id="select-group-1" class="form-group mt-2 green-text">
          <label class="mt-2 green-text" for="select-1">Selecione o tipo de usuário</label>
          <select value="Paciente" id="select-1" class="form-select" v-model="userType" required>
            <option v-for="user in userTypeOptions" :key="user.value">{{user.text}}</option>
          </select>
        </div>
        <b-button id="btn-entrar" type="submit" variant="success" class="green-text">Entrar</b-button>
      </form>
      <div class="text-center">
        <a id="link-cadastro" href="#">Não tem cadastro? Cadastre-se</a>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Login",
    data() {
      return {
        user: {
          login: "", 
          password: ""
        },
        users: [
          {id: 0, login:"anderson", password:"123"},
          {id: 2, login:"andzika", password:"123"}
        ],
        userType: "Paciente",
        userTypeOptions: [
          {value: "Paciente", text: "Paciente"},
          {value: "Médico", text: "Médico"},
          {value: "Recepcionista", text: "Recepcionista"}
        ]
      }
    }
  }
  
</script>

<style>

*{
  font-family: Arial, Sans-serif, times;
}

#app {
  width: 100vw;
  height: 100vh;
  background: linear-gradient(to bottom right,#fff, var(--primary));
  display: flex;
  align-items: center;
  justify-content: center;
}

#form-title{
  text-align: center;
}

.title{
  font-size: 15pt;
  margin-bottom: 20px;
}

.green-text{
  color: var(--primary);
  font-weight: bold;
}

.logo-img{
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 30px;
}

.logo-img img{
  border: 0;
  width: 120px;
  height: 80px;
}

#form{
  width: 33%;
  background-color: var(--background-default);
  border: 1px solid #aaa;
  border-radius: 8px;
  padding: 40px 15px 40px 15px;
  box-shadow: -18px -18px var(--primary);
}

#select-1:hover{
  cursor: pointer;
}

.text-center{
  width: 100%;
  text-align: center;
}

#btn-entrar{
  width: 100%;
  font-weight: bold;
  margin-top: 50px;
  margin-bottom: 10px;
}

#btn-entrar:hover{
  background-color: var(--secondary);
}

#link-cadastro{
  text-decoration: none;
}
</style>
