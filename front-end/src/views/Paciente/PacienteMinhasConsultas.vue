<template>
  <div id="pacienteConsultas">
    <NavbarPaciente/>
    <div class="bloco-consultas">
      <table>
        <tr v-for="(cons,index) in divide(consultas)" :key="index">
          <td v-for="c in cons" :key="c.id">
            <div class="card">
              <img src="../../assets/images/logo.png" class="card-img-top" alt="Logo"> 
              <div class="card-body">
                <h5 class="card-title">Consulta</h5>
                <p class="card-text"><b>Médico:</b> {{c.medico.nome}}</p>
                <p class="card-text"><b>Data:</b> {{c.data}}</p>
                <p class="card-text"><b>Horário:</b> {{c.horario}}</p>
                <p class="card-text"><b>Preço:</b> {{c.preco}}</p>
              </div>
              <div class="cart-footer">
                <span class="icon btn-excluir" >
                  <font-awesome-icon :icon="['fas', 'trash-alt']" />
                </span>
                <router-link class="icon btn-editar" :to="{name: 'pacienteEditarConsulta', params: {id: c.id, consulta:c}}">
                  <font-awesome-icon :icon="['fas', 'edit']" />
                </router-link>
              </div>
            </div>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import NavbarPaciente from '../../components/NavbarPaciente/NavbarPaciente.vue'


export default {
  data(){
    return {
      consultas: [
        {id: 0, preco: 100, data: '03-02-2022', medico: {nome: 'Anderson'}, horario: '9:00'},
        {id: 1, preco: 100, data: '03-02-2022', medico: {nome: 'Anderson'}, horario: '9:00'},
        {id: 2, preco: 100, data: '03-02-2022', medico: {nome: 'Anderson'}, horario: '9:00'},
        {id: 3, preco: 100, data: '03-02-2022', medico: {nome: 'Anderson'}, horario: '9:00'},
        {id: 4, preco: 100, data: '03-02-2022', medico: {nome: 'Anderson'}, horario: '9:00'},
        {id: 5, preco: 100, data: '03-02-2022', medico: {nome: 'Anderson'}, horario: '9:00'},
        {id: 6, preco: 100, data: '03-02-2022', medico: {nome: 'Anderson'}, horario: '9:00'},
        {id: 7, preco: 100, data: '03-02-2022', medico: {nome: 'Anderson'}, horario: '9:00'},
        {id: 8, preco: 100, data: '03-02-2022', medico: {nome: 'Anderson'}, horario: '9:00'},
      ],
    }
  },
  methods: {
    divide(cons){
      let res = []
      let aux = []
      for(let i=0; i<cons.length; i++){
        aux.push(cons[i])
        if(((i+1)%4 === 0 && i !== 0) || i === cons.length-1){
          res.push(aux)
          aux = []
        }
      }
      console.log(res)
      return res
    }
  },
  components:{
    NavbarPaciente
  }
}
</script>

<style>
.bloco-consultas{
  display: flex;
}

.bloco-consultas table{
  margin-left: 10px;
  margin-right: 10px;
}

.bloco-consultas td{
  min-width: 18vw;
  padding: 10px;
}

.card{
  display: flex;
  width: 100%;
  height: 100%;
}

.card img{
  width: 100%;
  height: 100%;
  margin-bottom: 0vw;
}

.card-body{
  margin-top: 0;
  width: 100%;
  height: 100%;
}

.card-title{
  font-size: 2vw;
  margin-bottom: 0.1vw;
  margin-top: 0;
}
.card-text{
  font-size: 1.5vw;
  margin-bottom: 0.05vw;
  margin-top: 0;
}

.card-footer{
  width: 100%;
  height: 3vw;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon {
  margin-left: 10px;
  margin-right: 10px;
}

.icon:hover{
  cursor: pointer;
}

.btn-editar{
  color: var(--primary);
}

.btn-excluir{
  color: red;
}

</style>